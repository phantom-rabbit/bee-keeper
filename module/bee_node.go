package module

import "gorm.io/gorm"

type BeeNode struct {
	gorm.Model
	Ip   string
	Port int

	Owner    string
	Contract string
}

