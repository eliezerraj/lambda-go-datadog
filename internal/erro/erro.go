package erro

import (
	"errors"

)

var (
	ErrUnmarshal 		= errors.New("Erro na conversão do JSON")
	ErrUnauthorized 	= errors.New("Erro de autorização")
	ErrMethodNotAllowed  = errors.New("Método não disponivel")
)