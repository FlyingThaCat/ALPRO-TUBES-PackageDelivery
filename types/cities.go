package types

type Tempat struct {
	City      string
	Longitude float64
	Latitude  float64
}

type Cities string

type ListKota []string

// Data list kota
var Kota = ListKota{
	"Jakarta",
	"Bogor",
	"Depok",
	"Tangerang",
	"Bekasi",
}

const (
	Jakarta   Cities = "Jakarta"
	Bogor     Cities = "Bogor"
	Depok     Cities = "Depok"
	Tangerang Cities = "Tangerang"
	Bekasi    Cities = "Bekasi"
)

var DaftarTempat = []Tempat{
	{City: "Jakarta", Longitude: 106.8456, Latitude: -6.2088},
	{City: "Bogor", Longitude: 106.7930, Latitude: -6.5917},
	{City: "Depok", Longitude: 106.8188, Latitude: -6.4025},
	{City: "Tangerang", Longitude: 106.6406, Latitude: -6.1702},
	{City: "Bekasi", Longitude: 106.9896, Latitude: -6.2416},
}

// AllCities mengembalikan list nama kota sebagai []string
func AllCities() []string {
	return Kota
}

// ParseCity mengkonversi string nama kota ke tipe Cities
func ParseCity(cityName string) Cities {
	switch cityName {
	case "Jakarta":
		return Jakarta
	case "Bogor":
		return Bogor
	case "Depok":
		return Depok
	case "Tangerang":
		return Tangerang
	case "Bekasi":
		return Bekasi
	default:
		return Cities(cityName)
	}
}
