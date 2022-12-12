package models

import (
	"time"

	"gorm.io/gorm"
)

type Pokemon struct {
	gorm.Model
	Name           string  `json:"name"`
	Gender         string  `json:"gender"`
	Category       string  `json:"category"`
	Description    string  `json:"description"`
	Height         float64 `json:"height"`
	Weight         float64 `json:"weight"`
	Level          string  `json:"level"`
	Favorite       bool    `json:"favorite"`
	HealthPoints   int     `json:"health_points"`
	Attack         int     `json:"attack"`
	SpecialAttack  int     `json:"special_attack"`
	Defense        int     `json:"defense"`
	SpecialDefense int     `json:"special_defense"`
	Velocity       int     `json:"velocity"`

	Types []Type `gorm:"many2many:pokemons_types"`
}

type Type struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`

	Pokemons []Pokemon `gorm:"many2many:pokemons_types"`
}

type PokemonsTypes struct {
	PokemonID int `gorm:"primaryKey" json:"pokemon_id"`
	TypeID    int `gorm:"primaryKey" json:"type_id"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Image struct {
	gorm.Model
	PokemonID   int     `gorm:"primaryKey" json:"pokemon_id"`
	Pokemon     Pokemon `json:"pokemon"`
	Image       string  `json:"image"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
}

type Video struct {
	gorm.Model
	PokemonID   int     `gorm:"primaryKey" json:"pokemon_id"`
	Pokemon     Pokemon `json:"pokemon"`
	Video       string  `json:"video"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
}
