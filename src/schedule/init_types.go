package schedule

type InitVariable struct {
	TriggerType TriggerTypes
}
type TriggerTypes int

const (
	_           TriggerTypes = iota
	EtcdTrigger
	RestTrigger
	RpcTrigger
)
