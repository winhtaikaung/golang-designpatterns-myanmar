package factory

type hmk struct {
	gun
}

func newHmk() iGun {
	return &hmk{
		gun: gun{
			name:  "HMK gun",
			power: 1,
		},
	}
}
