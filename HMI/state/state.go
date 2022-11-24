package state

type CurrentState string

const (
	NORMAL      CurrentState = "normal"
	Recording   CurrentState = "recording"
	Calibrating CurrentState = "calibrating"
)

type SystemState struct {
	State CurrentState
}

func NewState() *SystemState {
	return &SystemState{
		State: NORMAL,
	}
}

func (st *SystemState) Transition(newState CurrentState) {
	st.State = newState
}
