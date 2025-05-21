package types

type PackageTypeList []string

var PackageTypes = PackageTypeList{"Reguler", "Express", "SameDay"}

func IsValidPackageType(packageType string) bool {
	for _, validType := range PackageTypes {
		if validType == packageType {
			return true
		}
	}
	return false
}