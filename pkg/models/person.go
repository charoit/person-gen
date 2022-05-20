package models

type Person struct {
	Surname    string `csv:"surname" json:"surname,omitempty"`
	Name       string `csv:"name" json:"name,omitempty"`
	Patronymic string `csv:"patronymic" json:"patronymic,omitempty"`
	Age        int    `csv:"age" json:"age"`
	Email      string `csv:"email" json:"email,omitempty"`
	Phone      string `csv:"phone" json:"phone,omitempty"`
}
