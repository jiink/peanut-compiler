// fsm.go
package fsm
const (
	StateStart = iota
	StateIdentifier
	StateInteger
	StateReal
)
type FSM struct {
	currentState int
}
// create a new one and sets it
func NewFSM() *FSM {
	return &FSM{currentState: StateStart}
}
// resets fsm
func (f *FSM) Reset() {
	f.currentState = StateStart
}
// transition updates based on input
func (f *FSM) Transition(inputChar rune) {
		switch f.currentState {
	case StateStart:
		if isLetter(inputChar) {
			f.currentState = StateIdentifier
		} else if isDigit(inputChar) {
			f.currentState = StateInteger
		} else if inputChar == '.' {
			f.currentState = StateReal
		}
	case StateIdentifier:
		if !isLetter(inputChar) && !isDigit(inputChar) {
			f.currentState = StateStart
		}
	case StateInteger:
		if !isDigit(inputChar) {
			if inputChar == '.' {
				f.currentState = StateReal
			} else {
				f.currentState = StateStart
			}
		}
	case StateReal:
		if !isDigit(inputChar) {
			f.currentState = StateStart
		}
	}
}
// checks if it's in an accepting state
func (f *FSM) IsAcceptingState() bool {
	return f.currentState == StateIdentifier || f.currentState == StateInteger || f.currentState == StateReal
}
// checks if it's a letter
func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}
// checks if it's a digit
func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
