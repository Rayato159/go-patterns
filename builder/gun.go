package builder

import (
	"encoding/json"
	"fmt"
)

// Product part
type IGunBuilder interface {
	SetName(name string)
	setMod(m *mod)
	setPrice(price float64)
	getName() string
	getMod() *mod
	getPrice() float64
	getGun() *gun
	Display()
}

func (g *gun) SetName(name string)    { g.Name = name }
func (g *gun) setMod(m *mod)          { g.Mod = m }
func (g *gun) setPrice(price float64) { g.Price = price }
func (g *gun) getName() string        { return g.Name }
func (g *gun) getMod() *mod           { return g.Mod }
func (g *gun) getPrice() float64      { return g.Price }
func (g *gun) getGun() *gun           { return g }
func (g *gun) Display() {
	obj, _ := json.MarshalIndent(&g, "", "\t")
	fmt.Println(string(obj))
}

type gun struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Mod   *mod    `json:"mod"`
}

type scope string

const (
	x1 scope = "x1"
	x2 scope = "x2"
)

type mod struct {
	Laser      bool  `json:"laser"`
	Flashlight bool  `json:"flashlight"`
	Silencer   bool  `json:"silencer"`
	Scope      scope `json:"scope"`
}

func GunInit() IGunBuilder {
	return &gun{
		Mod: &mod{},
	}
}

// Engineer part
type GunEngineer struct {
	builder IGunBuilder
}

func SpawnGunEngineer(b IGunBuilder) *GunEngineer {
	return &GunEngineer{
		builder: b,
	}
}

func (en *GunEngineer) SetBuilder(b IGunBuilder) { en.builder = b }

func (en *GunEngineer) BuildGun() (*gun, error) {
	switch en.builder.getName() {
	case "ak47":
		en.builder.setMod(&mod{
			Laser:      false,
			Flashlight: false,
			Silencer:   false,
			Scope:      x1,
		})
		en.builder.setPrice(2700)
		return en.builder.getGun(), nil
	case "m4a4":
		en.builder.setMod(&mod{
			Laser:      true,
			Flashlight: true,
			Silencer:   false,
			Scope:      x2,
		})
		en.builder.setPrice(3100)
		return en.builder.getGun(), nil
	case "m4a1":
		en.builder.setMod(&mod{
			Laser:      true,
			Flashlight: false,
			Silencer:   true,
			Scope:      x1,
		})
		en.builder.setPrice(2900)
		return en.builder.getGun(), nil
	default:
		return nil, fmt.Errorf("engineer doesn't have a knowledge to build that gun")
	}
}

func Output() {
	list := []string{"m4a1", "m4a4", "ak47"}

	gunTemplate := GunInit()
	for _, v := range list {
		gunTemplate.SetName(v)
		gunEngineer := SpawnGunEngineer(gunTemplate)

		// Building
		gun, err := gunEngineer.BuildGun()
		if err != nil {
			fmt.Println(err)
			continue
		}
		gun.Display()
	}
}
