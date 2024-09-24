package contas

import (
	"fmt"
	t "orientecao_obj_go/titular"
)

type ContaPoupanca struct {
	Titular                              t.Titular
	NumeroAgencia, Operacao, NumeroConta int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) string {

	if c.saldo >= valorSaque {
		c.saldo -= valorSaque
		return "Saque feito com sucesso!"
	}

	return "Houve um problema no saque!"
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) string {

	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito feito com sucesso!"
	}

	return "Houve um problema no Deposito!"
}

func (c *ContaPoupanca) ObterSaldo() string {
	return fmt.Sprint("Saldo: ", c.saldo)
}
