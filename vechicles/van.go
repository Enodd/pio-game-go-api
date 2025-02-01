package vechicles

import "fmt"

type Van struct {
	Car
	dlugosc int
}

type VanOption func(*Van)

func (v *Van) GetInfo() string {
	return fmt.Sprintf("%s\nmarka: %s\ntyp: %s\nliczba miejsc: %d\ndlugosc: %d", v.id, v.marka, "Dostawczak", v.liczbaMiejsc, v.dlugosc)
}

func (v *Van) SetDlugosc(dlugosc int) {
	v.dlugosc = dlugosc
	_ = v
}

func (v *Van) WillLoadFit(loadLength int) bool {
	return loadLength <= v.dlugosc
}

func NewVan(marka string) *Van {
	return &Van{
		Car: Car{
			marka: marka,
		},
		dlugosc: 250,
	}
}

func NewVanDlugosc(marka string, dlugosc int) *Van {
	return &Van{
		Car: Car{
			marka: marka,
		},
		dlugosc: dlugosc,
	}
}
