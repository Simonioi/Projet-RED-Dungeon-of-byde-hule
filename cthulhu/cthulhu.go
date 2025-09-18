package cthulhu

type Attack struct {
	Name      string
	Damage    int
	HitChance float64
}

type Monster struct {
	Name      string
	MaxHP     int
	CurrentHP int
	Attack1   []Attack
	Attack2   []Attack
	Attack3   []Attack
	Attack4   []Attack
	Attack5   []Attack
}
// création de Cthulhu, le Boss Final
func Cthulhu() Monster {
	return Monster{
		Name:      "Cthulhu",
		MaxHP:     500,
		CurrentHP: 500,
		Attack1: []Attack{
			{Name: "trempette", Damage: 1, HitChance: 1.0},
		},
		Attack2: []Attack{
			{Name: "tentacule", Damage: 2, HitChance: 1.0},
		},
		Attack3: []Attack{
			{Name: "lancé de canard en plastique", Damage: 1, HitChance: 1.0},
		},
		Attack4: []Attack{
			{Name: "image subliminale", Damage: 4, HitChance: 1.0},
		},
		Attack5: []Attack{
			{Name: "effondrement mental", Damage: 50, HitChance: 1.0},
		},
	}
}
