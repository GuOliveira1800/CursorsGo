package contas

import (
	"fmt"
	t "orientecao_obj_go/titular"
)

type ContaCorrente struct {
	Titular       t.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {

	if c.saldo >= valorSaque {
		c.saldo -= valorSaque
		return "Saque feito com sucesso!"
	}

	return "Houve um problema no saque!"
}

func (c *ContaCorrente) Depositar(valorDeposito float64) string {

	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito feito com sucesso!"
	}

	return "Houve um problema no Deposito!"
}

func (c *ContaCorrente) Transferencia(valor float64, contaDestino *ContaCorrente) string {

	if c.saldo >= valor {
		c.Sacar(valor)
		contaDestino.Depositar(valor)
		return "Transferencia feita com sucesso!"
	}

	return "Houve um problema na Transferencia!"
}

func (c *ContaCorrente) ObterSaldo() string {
	return fmt.Sprint("%s%f", "Saldo: ", c.saldo)
}
