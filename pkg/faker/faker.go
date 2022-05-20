package faker

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/charoit/person-gen/pkg/models"
	"github.com/essentialkaos/translit/v2"
)

type Faker struct {
	data *Fake
}

func New(data *Fake) *Faker {
	rand.Seed(time.Now().UnixNano())
	return &Faker{
		data: data,
	}
}

func (f *Faker) MakePerson(asset string, sex Sex) models.Person {
	person := models.Person{}

	if _, ok := f.data.Assets[asset]; ok {
		if sex == All {
			sex = Sex(f.randomRange(int(Female), int(Male)))
		}
		person.Surname = f.randSurname(asset, sex)
		person.Name = f.randName(asset, sex)
		person.Patronymic = f.randPatronymic(asset, sex)
		person.Email = f.randEmail(person.Surname)
		person.Phone = f.randPhone()
	}

	return person
}

func (f *Faker) randomRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func (f *Faker) randSurname(asset string, sex Sex) string {
	idx := rand.Intn(len(f.data.Assets[asset][sex].Surname))
	return f.data.Assets[asset][sex].Surname[idx]
}

func (f *Faker) randName(asset string, sex Sex) string {
	idx := rand.Intn(len(f.data.Assets[asset][sex].Name))
	return f.data.Assets[asset][sex].Name[idx]
}

func (f *Faker) randPatronymic(asset string, sex Sex) string {
	idx := rand.Intn(len(f.data.Assets[asset][sex].Patronymic))
	return f.data.Assets[asset][sex].Patronymic[idx]
}

func (f *Faker) randEmail(surname string) string {
	suffix := f.randomRange(1950, 2005)
	prefix := strings.ToLower(translit.EncodeToICAO(surname))
	host := f.data.Hosts[rand.Intn(len(f.data.Hosts))]

	return fmt.Sprintf("%s%d@%s", prefix, suffix, host)
}

func (f *Faker) randPhone() string {
	return fmt.Sprintf("+7%d", f.randomRange(9000000000, 9999999999))
}
