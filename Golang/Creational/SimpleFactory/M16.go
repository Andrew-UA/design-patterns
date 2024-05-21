package main

type M16 struct {
	Gun
}

func newM16() IGun {
	return &M16{
		Gun{
			name:  "M16",
			power: 4,
		},
	}
}
