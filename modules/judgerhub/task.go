package judgerhub

const (
	// subType indicates a submission type
	subType = iota

	// solType indicates a solution type
	solType
)

type task struct {
	subType int

	// id indicates the database record id of the task
	id uint

	// ttl indicates the time limit of the problem for which the task is created
	ttl float64

	// eventChan used to communicate with the task_watcher go routine
	eventChan chan taskEvent
}

type taskEvent struct {
}

func (t *task) taskWatcher() {

}
