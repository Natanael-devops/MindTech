package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Contato struct {
	gorm.Model
	Nome     string `form:"nome" validate:"nonzero"`
	Telefone string `form:"telefone" validate:"nonzero"`
	Email    string `form:"email" validate:"nonzero"`
}

func ValidarNome(nome *Contato) error {
	if err := validator.Validate(nome); err != nil {
		return err
	}
	return nil
}

func ValidarTelefone(telefone *Contato) error {
	if err := validator.Validate(telefone); err != nil {
		return err
	}
	return nil
}

func ValidarEmail(email *Contato) error {
	if err := validator.Validate(email); err != nil {
		return err
	}
	return nil
}
