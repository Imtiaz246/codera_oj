package cronera

import (
	"fmt"
	"reflect"
	"time"
)

type Status string

var (
	StatusJobProvisioned Status = "Job has been provisioned"
	StatusTaskSuccessful Status = "Task has run successfully. Task serial no: %v"
)

// Signal represents the event sent from the cronera. Before and after the execution
// of the taskFunc a signal will be generated and sent through a channel.
type Signal struct {
	// cronTime is the time when the signal is generated
	time time.Time

	// serialNo indicates the number of time the task has executed
	serialNo int64

	// execRemaining indicates remaining number of times
	// the task will be executed in the future
	execRemaining int64

	// status indicates the status of the task.
	status string

	// returnValue holds the return value from the taskFunc
	// after the successful execution
	returnValues []reflect.Value
}

func newSignalWithOptions(t time.Time, sn, er int64, st Status, rv []reflect.Value) *Signal {
	s := &Signal{
		time:          t,
		serialNo:      sn,
		execRemaining: er,
		returnValues:  rv,
	}
	if st == StatusJobProvisioned {
		s.status = string(StatusJobProvisioned)
	} else {
		s.status = fmt.Sprintf(string(st), sn)
	}

	return s
}

func createJobProvisionedSignal(c *Cronera) *Signal {
	return &Signal{
		time:          time.Now(),
		serialNo:      0,
		execRemaining: c.totExecNo,
		status:        string(StatusJobProvisioned),
	}
}

func dropSignalIfFull(signalStream chan *Signal) {
	if len(signalStream) == cap(signalStream) {
		<-signalStream
	}
}
