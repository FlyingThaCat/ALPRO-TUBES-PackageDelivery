package utils

import (
	"fmt"
	"strings"

	"PackageDelivery/types"
)

func ParseCity(name string) (types.Cities, error) {
	switch strings.ToLower(name) {
	case "jakarta":
		return types.Jakarta, nil
	case "bogor":
		return types.Bogor, nil
	case "depok":
		return types.Depok, nil
	case "tangerang":
		return types.Tangerang, nil
	case "bekasi":
		return types.Bekasi, nil
	default:
		return "", fmt.Errorf("kota tidak dikenali: %s", name)
	}
}
