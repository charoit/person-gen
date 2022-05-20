package storage

import "github.com/charoit/person-gen/pkg/faker"

type Keeper interface {
	Load() (faker.Fake, error)
}
