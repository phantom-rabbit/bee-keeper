package module

import "gorm.io/gorm"

type Node struct {
	gorm.Model
	Ip       string
	Port     int
	User     string
	Password string
}

func GetNodeById(id interface{}) (Node, error) {
	var n Node
	err := conn.Where("id = ?", id).Take(&n).Error
	return n, err
}