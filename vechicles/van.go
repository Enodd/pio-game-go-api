package vechicles

import "fmt"

type van struct {
	car
	dlugosc int
}

func NewVan(marka string) *van {
	return &van{
		car: car{
			marka: marka,
		},
		dlugosc: 250,
	}
}

func (v *van) SetDlugosc(dlugosc int) *van {
	v.dlugosc = dlugosc
	return v
}

func (v *van) GetInfo() string {
	return fmt.Sprintf("%s\nmarka: %s\ntyp: %s\nliczba miejsc: %d\ndlugosc: %d", v.id, v.marka, "dostawczy", v.liczbaMiejsc, v.dlugosc)
}
