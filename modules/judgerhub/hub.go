package judgerhub

type Interface interface {
	MustRegister(key any)
	StartActivity(key any)
	AcknowledgeActivity(key any)
	GetSignalChan() chan hubSignal
}

func NewJudgerHub() Interface {
	return &HubContext{}
}

type t interface{}
type HubContext struct {
	signalChan chan hubSignal
	hub        map[t]judgerContext
}

type judgerContext struct {
	cpus       int
	processing map[t]task
}

type hubSignal struct {
}

func (h *HubContext) MustRegister(key any) {

}

func (h *HubContext) StartActivity(key any) {

}

func (h *HubContext) AcknowledgeActivity(key any) {

}

func (h *HubContext) GetSignalChan() chan hubSignal {
	return h.signalChan
}
