/*

twig.go

Twig Model - Twig consists of a single URL tied to a KindlingId
Implements gorm ORM: https://github.com/jinzhu/gorm

*/

package main

import (
)

// Twig model - stores a Twig (URL) to a KindlingId
type Twig struct {
	KindlingId int    `gorm:"ForeignKey:Id" json:"-"`
	Url        string `json:"url"`
}

// Twigs - Slice of Twig
type Twigs []Twig

func CreateTableTwig() {
	db := OpenDbConn()

	// Create Table
	db.CreateTable(&Twig{})
}

func AddTwig(twig Twig) bool {
	db := OpenDbConn()

	// Add Twig to DB
	db.Create(twig)

	return true
}
