package builder

type indianChef struct {
	spices     string
	meat       string
	vegetables []string
}

func newIndianChef() *indianChef {
	return new(indianChef)
}

func (b *indianChef) setSpices() {
	b.spices = "Tendori Spices"
}

func (b *indianChef) setMeat() {
	b.meat = "Chicken"
}

func (b *indianChef) setVegetables() {
	vegetables := []string{"onion", "garlic", "Tomato"}
	b.vegetables = vegetables
}

func (b *indianChef) getCurry() curry {
	return curry{
		Spices:     b.spices,
		Meat:       b.meat,
		Vegetables: b.vegetables,
		Name:       "Tendori Chicken",
	}
}
