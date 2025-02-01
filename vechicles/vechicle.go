package vechicles

type Vechicle interface {
	SetId(newId string)
	SetLiczbaMiejsc(liczbaMiejsc int)
	GetInfo() string
	GetId() string
	GetLiczbaMiejsc() int
	GetMarka() string
}
