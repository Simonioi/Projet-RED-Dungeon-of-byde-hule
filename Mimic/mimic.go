package mimic

type Attack struct {
	Name      string
	Damage    int
	HitChance float64
}

type Monster struct {
	Name      string
	MaxHP     int
	CurrentHP int
	Attacks   []Attack
}

func Mimic() Monster {
	return Monster{
		Name:      "Mimic",
		MaxHP:     10,
		CurrentHP: 10,
		Attacks: []Attack{
			{Name: "Morsure", Damage: 2, HitChance: 0.85},
		},
	}
}
