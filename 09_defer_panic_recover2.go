package main

import (
	"errors"
	"fmt"
	"log"
)

// Account represents a bank account
type Account struct {
	Balance float64
}

// Transfer transfers money from one account to another
func Transfer(source, destination *Account, amount float64) error {
	if source.Balance < amount {
		return errors.New("insufficient balance in the source account")
	}

	// Simulate an unexpected failure during the transfer
	if amount > 1000 {
		panic("unexpected error: transfer amount too high")
	}

	source.Balance -= amount
	destination.Balance += amount

	return nil
}

func main() {
	// Create two accounts
	sourceAccount := &Account{Balance: 1500}
	destinationAccount := &Account{Balance: 500}

	// Function to process the transaction safely
	processTransaction := func() {
		// Use recover to catch any panic
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Error during transaction: %v\n", r)
			}
		}()

		// Simulate releasing resources after the transaction
		defer fmt.Println("Resources released (e.g., closing database connections)")

		// Attempt to perform the transfer
		err := Transfer(sourceAccount, destinationAccount, 1200)
		if err != nil {
			log.Println("Transfer error:", err)
			return
		}

		fmt.Println("Transfer completed successfully!")
	}

	// Process the transaction
	processTransaction()

	// Display final balances
	fmt.Printf("Source account balance: $%.2f\n", sourceAccount.Balance)
	fmt.Printf("Destination account balance: $%.2f\n", destinationAccount.Balance)
}
