package members

import "gorm.io/gorm"

type Member struct {
	ID   int
	Name string
	Age  int
	gorm.Model
}
