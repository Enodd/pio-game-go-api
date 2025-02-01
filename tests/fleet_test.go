package vechicles_test

import (
	"testing"

	. "pio/vechicles"

	"github.com/stretchr/testify/assert"
)

var fleet *Fleet = NewFleet()

func TestMain(m *testing.M) {
	mazda := NewCar("Mazda")
	mazda.SetId("SG99999")
	skoda := NewVan("Skoda")
	skoda.SetId("SG99998")
	m.Run()
}

func TestAddVechicle(t *testing.T) {
	newVechicle := NewVan("Skoda")
	newVechicle.SetId("SG99999")
	_, success := fleet.AddVechicle(newVechicle)
	assert.Equal(t, success, true)
}

func TestAddVechicleError(t *testing.T) {
	newVechicle := NewVan("Skoda")
	newVechicle.SetId("SG99999")
	err, _ := fleet.AddVechicle(newVechicle)
	assert.EqualError(t, err, "Vechicle is already in fleet")
}
