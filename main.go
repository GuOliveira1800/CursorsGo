package main

import (
	"fmt"
	"orientecao_obj_go/contas"
	"orientecao_obj_go/titular"
)

func PagarBoletos(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contadoGustavo := contas.ContaPoupanca{
		Titular: titular.Titular{
			"Gustavo",
			"08991198910",
			"Developer"},
		NumeroAgencia: 1,
		NumeroConta:   1,
	}

	contadoGustavo.Depositar(200)

	PagarBoletos(&contadoGustavo, 100)

	fmt.Println(contadoGustavo.ObterSaldo())
}
