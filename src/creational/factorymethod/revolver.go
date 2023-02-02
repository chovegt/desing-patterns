package factorymethod

type revolver struct {
	bullets int
}

func NewRevolver() Guns {
	return &revolver{
		bullets: 1,
	}
}

func (r *revolver) Name() string {
	return "revolver"
}

func (a *revolver) Recharge() string {
	a.bullets = 1
	return "recharge complete"
}

func (a *revolver) Shoot() string {

	if a.bullets > 0 {
		a.bullets--
		return "shoot revolver"
	}

	return "recharge revolver"
}
