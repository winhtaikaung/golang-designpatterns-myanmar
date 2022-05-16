package builder

type chineseChef struct {
	spices     string
	meat       string
	vegetables []string
}

func newChineseChef() *chineseChef {
	return new(chineseChef)
}

func (b *chineseChef) setSpices() {
	b.spices = "Mala Spices"
}

func (b *chineseChef) setMeat() {
	b.meat = "Chicken"
}

func (b *chineseChef) setVegetables() {
	vegetables := []string{"corn", "lili", "bokchoy"}
	b.vegetables = vegetables
}

func (b *chineseChef) getCurry() curry {
	return curry{
		Spices:     b.spices,
		Meat:       b.meat,
		Vegetables: b.vegetables,
		Name:       "mala curry",
	}
}
