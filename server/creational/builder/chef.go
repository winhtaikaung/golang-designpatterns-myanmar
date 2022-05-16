package builder

type chef struct {
	builder iChefCooker
}

func NewMainChef(b iChefCooker) *chef {
	mainChef := new(chef)
	mainChef.builder = b
	return mainChef
}

func (d *chef) SetBuilder(b iChefCooker) {
	d.builder = b
}

func (d *chef) MakeCurry() curry {
	d.builder.setMeat()
	d.builder.setSpices()
	d.builder.setVegetables()
	return d.builder.getCurry()
}
