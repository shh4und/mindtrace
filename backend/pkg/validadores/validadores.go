package validadores

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type CPF string

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())

	validate.RegisterCustomTypeFunc(ValidarCPF, CPF(""))
}

// funcao para validar struct com tags "validate" utilizando o go-validator
func ValidarStruct(structPValidar any) error {
	err := validate.Struct(structPValidar)
	if err != nil {
		var erroValidacaoInvalido *validator.InvalidValidationError
		if errors.As(err, &erroValidacaoInvalido) {
			return fmt.Errorf("codigo produziu um erro invalido para validacao: %v", err)
		}

		var errosValidacao validator.ValidationErrors
		if errors.As(err, &errosValidacao) {
			return fmt.Errorf("erro na validacao de dados: %v", errosValidacao.Error())
		}
	}

	return nil
}

func ValidarCPF(field reflect.Value) any {
	if cpf, ok := field.Interface().(CPF); ok {
		return string(cpf)
	}

	return nil
}
