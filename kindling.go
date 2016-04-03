/*

kindling.go

Kindling Model - Kindling consists of a set of Twigs (URLs)
Implements gorm ORM: https://github.com/jinzhu/gorm

*/

package main

import ()

// Kindling model - simple storage of Kindling Ids
type Kindling struct {
	Id    int    `json:"kindling_id"`
	Twigs []Twig `json:"twigs"`
}

// Kindlings - Slice of Kindling
type Kindlings []Kindling

func CreateTableKindling() {
	db := OpenDbConn()

	// Create Table
	db.CreateTable(&Kindling{})
}

func AddKindling() int {
	var kindling Kindling

	// Add Kindling to DB
	db := OpenDbConn()
	db.Create(&kindling)

	return kindling.Id
}

func GetKindling(kindlingId int) Kindling {
	var kindling Kindling

	db := OpenDbConn()
	db.Preload("Twigs").First(&kindling, kindlingId)

	return kindling
}
