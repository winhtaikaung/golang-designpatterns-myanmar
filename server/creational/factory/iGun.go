package factory

type iGun interface {
	setName(name string)
	setPower(power int)
	GetName() string
	GetPower() int
}
