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

func Skelly() Monster {
	return Monster{
		Name:      "Skelly",
		MaxHP:     20,
		CurrentHP: 20,
		Attacks: []Attack{
			{Name: "Épée", Damage: 5, HitChance: 0.90},
		},
		
	}
}
// Skelly est un squelette guerrier il est résistant, et frappe fort.