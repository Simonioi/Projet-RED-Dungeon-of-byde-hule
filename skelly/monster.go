package skelly

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

func NewSkelly() Monster {
	return Monster{
		Name:      "Skelly",
		MaxHP:     20,
		CurrentHP: 20,
		Attacks: []Attack{
			{Name: "Épée", Damage: 3, HitChance: 0.90},
			{Name: "Lancé d'os", Damage: 3, HitChance: 0.75},
		},
	}
}
