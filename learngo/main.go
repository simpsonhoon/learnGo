package main

import (
	"fmt"
	"learngo/modules/accounts"
)

func main() {
	account := accounts.NewAccount("simpson")
	account.Deposit(100)
	fmt.Println(account.Balance())
	err := account.Withdraw(200)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
