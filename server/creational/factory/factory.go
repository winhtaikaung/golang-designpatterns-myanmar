package factory

import "fmt"

func GetGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "hmk" {
		return newHmk(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
