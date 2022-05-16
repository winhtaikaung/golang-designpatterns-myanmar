/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"app/server/creational/factory"
	"fmt"
)

func main() {
	hmkFactory, _ := factory.GetGun("hmk")

	fmt.Println(hmkFactory.GetName())

}
