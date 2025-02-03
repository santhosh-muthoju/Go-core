package functions

import "fmt"

func CreateAccount(amount float64) func(float64) float64 {
	balance := amount

	return func(amount float64) float64 {
		balance += amount
		return balance
	}
}

func SamplePrintClosure(account func(float64) float64) {
	fmt.Println("After depositing 50:", account(50.0))
	fmt.Println("After withdrawing 30:", account(-30.0))
	fmt.Println("After depositing 20:", account(20.0))
}

func NewAccount(amt float64) func(float64) float64 {
	accBal := amt

	return func(amt float64) float64 {
		accBal = accBal + amt
		return accBal
	}

}
