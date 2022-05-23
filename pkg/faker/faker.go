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
		year, age := f.randYearAndAge()

		person.Surname = f.randSurname(asset, sex)
		person.Name = f.randName(asset, sex)
		person.Patronymic = f.randPatronymic(asset, sex)
		person.Age = age
		person.Email = f.randEmail(person.Surname, person.Name, year)
		person.Phone = f.randPhone()
	}

	return person
}

// randomRange random int in range.
func (f *Faker) randomRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// randSurname random surname.
func (f *Faker) randSurname(asset string, sex Sex) string {
	idx := rand.Intn(len(f.data.Assets[asset][sex].Surname))

	return f.data.Assets[asset][sex].Surname[idx]
}

// randName random name.
func (f *Faker) randName(asset string, sex Sex) string {
	idx := rand.Intn(len(f.data.Assets[asset][sex].Name))

	return f.data.Assets[asset][sex].Name[idx]
}

// randPatronymic random patronymic.
func (f *Faker) randPatronymic(asset string, sex Sex) string {
	idx := rand.Intn(len(f.data.Assets[asset][sex].Patronymic))

	return f.data.Assets[asset][sex].Patronymic[idx]
}

// randYearAndAge random birth year and age.
func (f *Faker) randYearAndAge() (int, int) {
	year := f.randomRange(1950, 2005)
	date := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	age := int(time.Now().Sub(date) / (time.Hour * 24 * 365))

	return year, age
}

// randEmail random email.
func (f *Faker) randEmail(surname, name string, year int) string {
	prefix := "email"
	first := strings.ToLower(translit.EncodeToICAO(surname))
	middle := strings.ToLower(translit.EncodeToICAO(name))
	host := f.data.Hosts[rand.Intn(len(f.data.Hosts))]

	return fmt.Sprintf("%s.%s.%s%d@%s", prefix, first, middle, year, host)
}

// randPhone random phone.
func (f *Faker) randPhone() string {
	return fmt.Sprintf("+7999%d", f.randomRange(9000000, 9999999))
}
