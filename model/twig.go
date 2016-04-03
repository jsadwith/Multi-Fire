/*

twig.go

Twig Model - Twig consists of a single URL tied to a KindlingId
Implements gorm ORM: https://github.com/jinzhu/gorm

*/

package model

import (
	"github.com/jsadwith/Multi-Fire/database"
)

// Twig model - stores a Twig (URL) to a KindlingId
type Twig struct {
	KindlingId int    `gorm:"ForeignKey:Id" json:"-"`
	Url        string `json:"url"`
}

// Twigs - Slice of Twig
type Twigs []Twig

func CreateTableTwig() {
	db := database.OpenDbConn()

	// Create Table
	db.CreateTable(&Twig{})
}

func AddTwig(twig Twig) bool {
	db := database.OpenDbConn()

	// Add Twig to DB
	db.Create(twig)

	return true
}
