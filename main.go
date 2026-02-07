package main

import (
	"fmt"
)

type Transaction struct {
	Category string
	Amount   int
}

func main() {
	transactions := []Transaction{
		{Category: "salary", Amount: 5000},
		{Category: "food", Amount: 500},
		{Category: "tech", Amount: 2000},
	}
	balance := 0

	for _, t := range transactions {

		if t.Category == "salary" {
			balance += t.Amount
			fmt.Printf("Поздравляем вам пришла %s ваш баланс теперь равен %d \n", t.Category, balance)

		} else {
			balance -= t.Amount
			fmt.Printf("cписание %d на категорию %s ваш баланс теперь равен %d\n", t.Amount, t.Category, balance)
		}
	}
}