package models

import "gorm.io/gorm"

type Driver struct {
  gorm.Model
  Name  string `json:"name"`
  CNH   string `json:"cnh" gorm:"unique"`
  BirthDate string `json:"birth_date"`
}