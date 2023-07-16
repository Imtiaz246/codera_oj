package cronera

import (
	"context"
	"fmt"
	"time"
)

type Cronera struct {
	// interval indicates the time duration of running the taskFunc.
	interval time.Duration

	// gap indicates the gap in days between two running jobs.
	// By default, it's set to 1.
	gap int64

	// startDate indicates the date from when the Cronera will be started.
	// By default, it'll be set to start of the day if not set.
	startDate string

	// at indicates the times in a day for the job to be run.
	// For ex. 10:00, 12:00, ......., 23:00 etc
	at []string

	// totExecNo indicates how many times the Cronera will run.
	// By default, it's infinite if not set.
	totExecNo int64
}

func New() *Cronera {
	return &Cronera{
		totExecNo: -1,
		at:        make([]string, 0),
	}
}

func (c *Cronera) Every(gap int64) *Cronera {
	c.gap = gap
	return c
}

func (c *Cronera) Day() *Cronera {
	if c.gap == 0 {
		c.gap = 1
	}
	c.interval = time.Hour * 24 * time.Duration(c.gap)
	return c
}

func (c *Cronera) StartDate(d string) *Cronera {
	c.startDate = d
	return c
}

func (c *Cronera) At(st ...string) *Cronera {
	c.at = st
	return c
}

func (c *Cronera) MustExec(toteExecNo int64) *Cronera {
	c.totExecNo = toteExecNo
	return c
}

func (c *Cronera) Do(ctx context.Context, taskFunc any, args ...any) (<-chan *Signal, error) {
	if err := checkTaskFuncValidity(taskFunc, args...); err != nil {
		return nil, fmt.Errorf("task function is not valid: %v", err)
	}
	var (
		serialNo        = int64(0)
		signalStream    = make(chan *Signal, 10)
		nextRunningDate time.Time
	)

	if c.startDate != "" {
		date, err := parseDate(c.startDate)
		if err != nil {
			return nil, err
		}
		if date.Before(truncateToStartOfDay(time.Now())) {
			return nil, fmt.Errorf("startDate is before current date")
		}
		nextRunningDate = date
	} else {
		nextRunningDate = truncateToStartOfDay(time.Now())
	}

	go func() {
		signalStream <- createJobProvisionedSignal(c)
		defer close(signalStream)
		for {
			todayDate := <-waitUntil(nextRunningDate)
			todayTimes := parseTodayTimes(c.at, todayDate)
			for _, t := range todayTimes {
				if time.Now().Sub(t) > time.Second*1 {
					continue
				}
				select {
				case <-waitUntil(t):
					serialNo++
					go func() {
						out := execTaskFunc(taskFunc, args...)
						dropSignalIfFull(signalStream)
						signalStream <- newSignalWithOptions(t, serialNo, min(c.totExecNo, c.totExecNo-serialNo), StatusTaskSuccessful, out)
					}()
					if serialNo == c.totExecNo {
						return
					}
				case <-ctx.Done():
					return
				}
			}
			nextRunningDate = truncateToStartOfDay(nextRunningDate.Add(time.Duration(c.gap) * time.Hour * 24))
		}
	}()

	return signalStream, nil
}
