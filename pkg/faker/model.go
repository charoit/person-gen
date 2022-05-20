package faker

import "strings"

type Sex int

const (
	All Sex = iota
	Female
	Male
)

func ParseSex(sex string) Sex {
	switch strings.ToLower(sex) {
	case "female":
		return Female
	case "male":
		return Male
	default:
		return All
	}
}

func (s Sex) String() string {
	return []string{"all", "female", "male"}[s]
}

type Fake struct {
	Assets map[string]Asset
	Hosts  []string // Хосты для почты
}

type Asset map[Sex]Gender

// FakeData данные для генерации.
type FakeData struct {
	Name   string
	Gender []Gender // Может быть более 2-х полов, для тех кто не определился
	Hosts  []string // Хосты для почты
}

// Gender содержит данные для одного пола.
type Gender struct {
	Surname    []string
	Name       []string
	Patronymic []string
}
