package vechicles

import "fmt"

type car struct {
	marka        string
	id           string
	liczbaMiejsc int
}

func NewCar(marka string) *car {
	return &car{
		marka: marka,
	}
}

func (c *car) GetMarka() string {
	return c.marka
}

func (c *car) SetId(id string) *car {
	c.id = id
	return c
}

func (c *car) GetId() string {
	return c.id
}

func (c *car) SetLiczbaMiejsc(liczbaMiejsc int) *car {
	c.liczbaMiejsc = liczbaMiejsc
	return c
}

func (c *car) GetLiczbaMiejsc() int {
	return c.liczbaMiejsc
}

func (c *car) GetInfo() string {
	return fmt.Sprintf("%s\nmarka: %s\ntyp: %s\nliczba miejsc: %d", c.id, c.marka, "osobowy", c.liczbaMiejsc)
}
