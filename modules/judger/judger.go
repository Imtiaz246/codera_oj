package judger

type Interface interface {
	MustRegister(key any)
	StartActivity(key any)
	AcknowledgeActivity(key any)
}

func NewJudger() *Type {
	return &Type{}
}

type Type struct {
	registeredJudger judgerConn
}

type t interface{}
type connectChan chan chanType
type chanType struct {
	// todo: fill up data
}
type judgerConn map[t]connectChan

func (j *Type) MustRegister(key any) {

}

func (j *Type) StartActivity(key any) {

}

func (j *Type) AcknowledgeActivity(key any) {

}
