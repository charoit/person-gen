package files

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/charoit/person-gen/pkg/faker"
)

type Storage struct {
	sys embed.FS
}

func NewStorage(sys embed.FS) *Storage {
	return &Storage{
		sys: sys,
	}
}

func (s *Storage) Load() (faker.Fake, error) {
	var fake = faker.Fake{
		Assets: map[string]faker.Asset{},
	}
	err := fs.WalkDir(s.sys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		dir := strings.Split(path, string(os.PathSeparator))
		//fmt.Println(path, len(dir))

		switch {
		// hosts
		case len(dir) == 2 && !d.IsDir():
			if fake.Hosts, err = s.readFile(path); err != nil {
				return err
			}
		// asset
		case len(dir) == 2 && d.IsDir():
			fake.Assets[dir[1]] = map[faker.Sex]faker.Gender{}

		// gender
		case len(dir) == 3 && d.IsDir():
			fake.Assets[dir[1]][faker.ParseSex(dir[2])] = faker.Gender{}

		// data file
		case len(dir) == 4 && !d.IsDir():
			var data []string
			if data, err = s.readFile(path); err != nil {
				return err
			}

			sex := faker.ParseSex(dir[2])
			gender := fake.Assets[dir[1]][sex]
			switch {
			case strings.HasPrefix(dir[3], "name"):
				gender.Name = append(gender.Name, data...)
			case strings.HasPrefix(dir[3], "patronymics"):
				gender.Patronymic = data
			case strings.HasPrefix(dir[3], "surname"):
				gender.Surname = data
			}
			fake.Assets[dir[1]][sex] = gender
		}
		return nil
	})

	return fake, err
}

func (s *Storage) readFile(path string) ([]string, error) {
	content, err := s.sys.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	return strings.Split(string(content), "\n"), nil
}
