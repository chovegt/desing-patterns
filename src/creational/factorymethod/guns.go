package factorymethod

type Guns interface {
	Name() string
	Recharge() string
	Shoot() string
}
