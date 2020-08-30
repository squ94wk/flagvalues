package values

type Value struct {
	setter   func(string) error
	stringer func() string
	typ      func() string
	boolFlag bool
}

type Option func(*Value)

func NewValue(typ func() string, setter func(string) error, stringer func() string, opts ...Option) Value {
	v := Value{
		setter:   setter,
		stringer: stringer,
		typ:      typ,
	}
	for _, opt := range opts {
		opt(&v)
	}
	return v
}

func (g Value) String() string {
	return g.stringer()
}

func (g Value) Set(val string) error {
	return g.setter(val)
}

func (g Value) Type() string {
	return g.typ()
}

func (g Value) IsBoolFlag() bool {
	return g.boolFlag
}

func WithBoolFlag() Option {
	return func(v *Value) {
		v.boolFlag = true
	}
}
