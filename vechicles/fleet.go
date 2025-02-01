package vechicles

import (
	"errors"
	"strings"
)

type Fleet struct {
	vechicles map[string]Vechicle
}

func (f *Fleet) contains(id string) bool {
	_, found := f.vechicles[id]
	return found
}

func (f *Fleet) AddVechicle(v Vechicle) (error, bool) {
	if f.contains(v.GetId()) {
		return errors.New("Vechicle is already in fleet"), false
	}
	f.vechicles[v.GetId()] = v
	return nil, true
}

func (f *Fleet) RemoveVechicle(id string) (error, bool) {
	if !f.contains(id) {
		return errors.New("Vechicle isnt present in fleet"), false
	}
	delete(f.vechicles, id)
	return nil, true
}

func (f *Fleet) GetFleet() map[string]Vechicle {
	return f.vechicles
}

func (f *Fleet) GetFilteredFleet(marka string) map[string]Vechicle {
	filtered := make(map[string]Vechicle)
	for id, vehicle := range f.vechicles {
		if strings.EqualFold(strings.ToLower(vehicle.GetMarka()), strings.ToLower(marka)) {
			filtered[id] = vehicle
		}
	}
	return filtered
}

func NewFleet() *Fleet {
	return &Fleet{
		vechicles: make(map[string]Vechicle),
	}
}
