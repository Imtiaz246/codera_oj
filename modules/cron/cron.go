package cron

import (
	"fmt"
	"reflect"
	"time"
)

type Cron struct {
	// interval indicates the time duration of running the taskFunc
	interval time.Duration

	// startsAt indicates the time from when the Cron will be started
	startsAt time.Time

	// taskFunc indicates the task function what will be run
	// in every interval period from startsAt time. The taskFunc
	// has tobe a func type.
	taskFunc any

	// totExecNo indicates how many times the Cron will run.
	// By default, it's infinite if not set.
	totExecNo int

	// signalChan is a `receive only` channel where the cronSignal object
	// is sent before and after of every execution of the taskFunc
	signalChan <-chan Signal
}

type Status string

const (
	CronStatusTaskProvisioned Status = "task has been provisioned"
	CronStatusTaskSuccessfull Status = "task has run successfully"
)

// Signal represents the event sent from the cron. Before and after the execution
// of the taskFunc a signal will be generated and sent.
type Signal struct {
	// cronTime is the time when the signal is generated
	cronTime time.Time

	// serialNo indicates the number of time the task has executed
	serialNo int

	// execRemaining indicates remaining number of times
	// the task will be executed in the future
	execRemaining int

	// status indicates the status of the task.
	status Status

	// returnValue holds the return value from the taskFunc
	// after the successful execution
	returnValue []reflect.Value
}

func NewCron() *Cron {
	return &Cron{
		totExecNo: -1,
	}
}

func (c *Cron) Every(dur time.Duration) *Cron {
	c.interval = dur
	return c
}

func (c *Cron) StartsAt(t time.Time) *Cron {
	c.startsAt = t
	return c
}

func (c *Cron) Do(taskFunc interface{}) (<-chan Signal, error) {
	if reflect.ValueOf(taskFunc).Kind() != reflect.Func {
		return nil, fmt.Errorf("task type has to be a function")
	}
	// todo: handle logic and run the task function

	return c.signalChan, nil
}
