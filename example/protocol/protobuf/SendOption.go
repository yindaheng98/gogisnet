package protocol

type SendOption struct {
	Option string
}

func (o SendOption) String() string {
	return o.Option
}
