package openkcb

type Error struct {
	Code string
	Msg  string
}

func (e Error) Error() string {
	return "kcb: code(" + e.Code + "), msg(" + e.Msg + ")"
}
