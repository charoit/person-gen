package faker

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type (
	// FakeData данные для генерации.
	FakeData struct {
		Gender []Gender // Может быть более 2-х полов, для тех кто не определился
		Hosts  []string // Хосты для почты
	}

	// Gender содержит данные для одного пола.
	Gender struct {
		Surname    []string
		Name       []string
		Patronymic []string
	}

	// Person сгенерированная персона.
	Person struct {
		Surname    string
		Name       string
		Patronymic string
		Email      string
		Phone      string
		Region     string
	}
)

type Generator struct {
	data *FakeData
}

func NewGenerator(data *FakeData) *Generator {
	rand.Seed(time.Now().UnixNano())
	return &Generator{
		data: data,
	}
}

func (g *Generator) MakePerson() Person {
	gender := rand.Intn(len(g.data.Gender))
	surname := g.rndSurname(gender)
	person := Person{
		Surname:    surname,
		Name:       g.rndName(gender),
		Patronymic: g.rndPatronymic(gender),
		Email:      g.rndEmail(surname),
		Phone:      g.rndPhone(),
		Region:     g.rndRegion(),
	}

	return person
}

func (g *Generator) rndSurname(gender int) string {
	return g.data.Gender[gender].Surname[rand.Intn(len(g.data.Gender[gender].Surname))]
}

func (g *Generator) rndName(gender int) string {
	return g.data.Gender[gender].Name[rand.Intn(len(g.data.Gender[gender].Name))]
}

func (g *Generator) rndPatronymic(gender int) string {
	return g.data.Gender[gender].Patronymic[rand.Intn(len(g.data.Gender[gender].Patronymic))]
}

func (g *Generator) rndEmail(surname string) string {
	min := 1950
	max := 2005
	suffix := rand.Intn(max-min+1) + min // в диапазоне 1950-2005гг
	prefix := strings.ToLower(translit.EncodeToICAO(surname))
	host := g.data.Hosts[rand.Intn(len(g.data.Hosts))]

	return fmt.Sprintf("%s%d@%s", prefix, suffix, host)
}

func (g *Generator) rndPhone() string {
	min := 9000000000
	max := 9999999999
	return fmt.Sprintf("+7%d", rand.Intn(max-min+1)+min)
}

func (g *Generator) rndRegion() string {
	min := 1
	max := 95
	return fmt.Sprintf("%d", rand.Intn(max-min+1)+min)
}
