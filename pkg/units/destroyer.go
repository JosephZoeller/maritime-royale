package units

type destroyer struct {
	symbol string
}

func NewDestroyer() destroyer {
	return destroyer{symbol: "D"}
}
func (s destroyer) Draw() {

}
