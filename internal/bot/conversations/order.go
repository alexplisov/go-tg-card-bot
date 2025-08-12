package conversations

type OrderStep = int

const (
	OrderStepName         OrderStep = 0
	OrderStepRequirements OrderStep = 1
)

type OrderConversationState struct {
	OrderStep    OrderStep
	UserName     string
	Requirements string
}

func NewOrderConversationState() *OrderConversationState {
	return &OrderConversationState{
		OrderStep:    OrderStepName,
		UserName:     "",
		Requirements: "",
	}
}
