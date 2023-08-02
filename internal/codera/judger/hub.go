package judger

import (
	"github.com/imtiaz246/codera_oj/modules/workqueue"
	"log"
	"time"
)

var (
	hub = NewJudgerHub()
)

func init() {

}

func NewJudgerHub() *HubContext {
	return &HubContext{
		submissionQueue: workqueue.NewWorkQueue[queueType](),
		hub:             make(map[any]*judgerContext),
	}
}

type HubContext struct {
	submissionQueue workqueue.Interface[queueType]
	hub             map[any]*judgerContext
}

type queueType struct {
	tableType int
	id        uint
	code      string
	language  string
	ttl       float64
}

type judgerContext struct {
	cpus       int
	notBefore  time.Time
	notAfter   time.Time
	processing map[any]*task
}

func (h *HubContext) MustRegister(key any, cpus int, notBefore, notAfter time.Time) {
	h.hub[key] = &judgerContext{
		cpus:       cpus,
		notBefore:  notBefore,
		notAfter:   notAfter,
		processing: make(map[any]*task),
	}

	// todo: spawn judger context watcher, which will watch for expiry time of that judger

	log.Printf("Judger with key: %v and cpus: %v registered successfully in judger hub  for time [%v-%v]", key, cpus, notBefore, notAfter)
}

func (h *HubContext) StartActivity(key any, tableType int, id uint, ttl float64, ttc int) {
	task := &task{
		tableType: tableType,
		id:        id,
		ttl:       ttl,
		ttc:       ttc,
		tcRuntime: make([]float64, 0),
		eventChan: make(chan TaskEvent),
	}
	taskKey := task.genTaskKey()
	h.hub[key].processing[taskKey] = task

	// FIXME: move to a dedicated function
	go task.taskWatcher()

	log.Printf("Started activity for <tableType: %v>, <id: %v>, <ttl: %v>, <ttc: %v>", tableType, id, ttc, ttc)
}

func (h *HubContext) AcknowledgeTaskActivity(key any) {

}
