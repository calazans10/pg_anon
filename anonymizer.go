package main

import "github.com/brianvoe/gofakeit/v6"

type Anonymizer struct {
	Type string
}

func (a Anonymizer) Fake(value string) string {
	switch a.Type {
	case "name":
		return gofakeit.Name()
	case "first_name":
		return gofakeit.Person().FirstName
	case "last_name":
		return gofakeit.Person().FirstName
	case "email":
		return gofakeit.Email()
	case "phone":
		return gofakeit.Phone()
	case "empty":
		return ""
	default:
		return value
	}
}
