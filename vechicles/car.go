package vechicles

import (
	"fmt"
	"strings"
)

type Car struct {
	marka        string
	id           string
	liczbaMiejsc int
}

func (c *Car) GetInfo() string {
	return fmt.Sprintf("%s\nmarka: %s\ntyp: %s\nliczba miejsc: %d", c.id, c.marka, "osobowy", c.liczbaMiejsc)
}

func (c *Car) GetMarka() string {
	return c.marka
}

func (c *Car) SetId(id string) {
	c.id = id
	_ = c
}

func (c *Car) GetId() string {
	return c.id
}

func (c *Car) SetLiczbaMiejsc(liczbaMiejsc int) {
	c.liczbaMiejsc = liczbaMiejsc
	_ = c
}

func (c *Car) GetLiczbaMiejsc() int {
	return c.liczbaMiejsc
}

func (c *Car) CheckId(id string) bool {
	return strings.EqualFold(strings.ToLower(c.id), strings.ToLower(id))
}

func NewCar(marka string) *Car {
	return &Car{
		marka: marka,
	}
}
