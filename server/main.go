/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"app/server/creational/builder"
	"app/server/creational/factory"
	"fmt"
)

func main() {
	// Factory Pattern
	hmkFactory, _ := factory.GetGun("hmk")

	fmt.Println(hmkFactory.GetName())

	// Builder pattern
	chineseChef := builder.GetChefBuilder("chinese")

	mainChef := builder.NewMainChef(chineseChef)
	curry := mainChef.MakeCurry()
	fmt.Println(curry.Name)

	// Set Builder to indian
	indianChef := builder.GetChefBuilder("indian")
	mainChef.SetBuilder(indianChef)
	curry = mainChef.MakeCurry()
	fmt.Println(curry.Name)

}
