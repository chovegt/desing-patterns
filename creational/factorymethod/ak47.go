package factorymethod

type Ak47 struct {
	// code here attributes...
	bullets int
}

func NewAk47() Guns {
	return &Ak47{
		bullets: 6,
	}
}

func (Ak47) Name() string {
	return "Ak47"
}

func (a *Ak47) Recharge() string {
	a.bullets = 6
	return "recharge complete"
}

func (a *Ak47) Shoot() string {

	if a.bullets > 0 {
		a.bullets--
		return "shoot ak47"
	}

	return "recharge ak47"
}
