package domain

import (
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
    Field string `json:"field"`
    Message   string `json:"message"`
}

func GetErrorMsg(fe validator.FieldError) string {
    switch fe.Tag() {
        case "required":
            return "Deve ser preenchido"
        case "email":
            return "E-mail inválido"
	}
return "Erro desconhecido"
}