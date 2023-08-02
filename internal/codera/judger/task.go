package judger

import (
	"fmt"
)

const (
	// subType indicates a submission type
	subType = iota

	// solType indicates a solution type
	solType
)

type task struct {
	// tableType indicates the table type for which the task is created for.
	// There are 2 kinds of table for this context (problem_submission, problem_solution)
	tableType int

	// id indicates the database record id of the task
	id uint

	// ttl indicates the time limit of the problem for which the task is created
	ttl float64

	// ttc indicates the total number of testcase of the problem for which the task is created
	ttc int

	// tcRuntime indicates the test case runtimes
	tcRuntime []float64

	// eventChan used to communicate with the task_watcher go routine
	eventChan chan TaskEvent
}

func (t *task) genTaskKey() string {
	return fmt.Sprintf("%v.%v", t.tableType, t.id)
}

// TaskEvent describes the events which will be sent for the task watcher
type TaskEvent struct {
}

// TaskResult describes the result which will be sent to the HubSignal
type TaskResult struct {
}

func (t *task) taskWatcher() {

}
