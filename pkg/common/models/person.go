package models

import _"gorm.io/gorm"

type Person struct {
    ID          uint64  `gorm:"primary_key" json:"id"`
    Name    	string  `json:"name"`
    Profissao 	string  `json:"profissao"`
    TeamID     uint    `json:"-"`
    Team		Team    `gorm:"foreignkey:TeamID:constraint:onUpadate:CASCADE,onDelete:CASCADE" json:"team"`
    TaskID     uint    `json:"-"`
    Task        Task    `gorm:"foreignkey:TaskID:constraint:onUpadate:CASCADE,onDelete:CASCADE" json:"task"`
}