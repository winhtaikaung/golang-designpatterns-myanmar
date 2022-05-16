package builder

type iChefCooker interface {
	setSpices()
	setMeat()
	setVegetables()
	getCurry() curry
}

func GetChefBuilder(builderType string) iChefCooker {
	if builderType == "chinese" {
		return &chineseChef{}
	}
	if builderType == "indian" {
		return &indianChef{}
	}
	return nil
}
