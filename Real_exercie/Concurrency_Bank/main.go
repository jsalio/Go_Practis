// Package main implements a concurrent banking system simulation
// that demonstrates safe concurrent operations on a bank account
// using mutex locks and channels for transaction processing.
package main

import (
	"fmt"
	"sync"
	"time"
)

// Account represents a bank account with a number and balance.
// It uses a mutex to ensure thread-safe operations on the balance.
type Account struct {
	number  int
	balance float64
	mutex   sync.Mutex
}

// TransactionType defines the possible types of transactions
// that can be performed on an account.
type TransactionType string

const (
	Deposit       TransactionType = "Deposit"       // Represents a deposit transaction
	Withdraw      TransactionType = "Withdraw"      // Represents a successful withdrawal
	FailedWitdraw TransactionType = "FailedWitdraw" // Represents a failed withdrawal attempt
)

// Transaction represents a banking transaction with its type,
// amount, and the final balance after the transaction.
type Transaction struct {
	Type        TransactionType // Type of transaction (Deposit/Withdraw/FailedWithdraw)
	Amount      float64         // Amount involved in the transaction
	FinalAmount float64         // Account balance after the transaction
}

// Deposit adds the specified amount to the account balance.
// It uses a mutex to ensure thread-safe operation and sends
// a transaction record through the provided channel.
func (acc *Account) Deposit(amount float64, transactionChannel chan<- Transaction) {
	acc.mutex.Lock()
	defer acc.mutex.Unlock()
	acc.balance += amount

	fmt.Printf("Depósito de %.2f realizado. Saldo actual: %.2f\n", amount, acc.balance)

	transactionChannel <- Transaction{
		Type:        Deposit,
		Amount:      amount,
		FinalAmount: acc.balance,
	}
}

// Withdraw attempts to subtract the specified amount from the account balance.
// If there are insufficient funds, it records a failed withdrawal.
// Uses mutex for thread safety and sends transaction details through the channel.
func (acc *Account) Withdraw(amount float64, transactionChannel chan<- Transaction) {
	acc.mutex.Lock()
	defer acc.mutex.Unlock()

	if acc.balance < amount {
		fmt.Printf("Retiro de %.2f fallido: saldo insuficiente (%.2f)\n", amount, acc.balance)
		transactionChannel <- Transaction{
			Type:        FailedWitdraw,
			Amount:      amount,
			FinalAmount: acc.balance,
		}
		return // Add return statement to prevent further execution
	}

	acc.balance -= amount
	fmt.Printf("Retiro de %.2f realizado. Saldo actual: %.2f\n", amount, acc.balance)

	transactionChannel <- Transaction{
		Type:        Withdraw,
		Amount:      amount,
		FinalAmount: acc.balance,
	}
}

// processTransaction handles the processing of transactions from the channel.
// It prints notifications for each transaction and continues until the channel is closed.
func processTransaction(acc *Account, transactionChannel chan Transaction, wg *sync.WaitGroup) {
	defer wg.Done()
	for transaction := range transactionChannel {
		if transaction.Type == FailedWitdraw {
			fmt.Printf("Notificación: Retiro fallido de %.2f, Saldo insuficiente: %.2f\n",
				transaction.Amount, transaction.FinalAmount)
		} else {
			fmt.Printf("Notificación: %s de %.2f, Saldo final: %.2f\n",
				transaction.Type, transaction.Amount, transaction.FinalAmount)
		}
	}
}

// main initializes and runs the concurrent banking simulation.
// It creates an account and spawns multiple goroutines to perform
// concurrent deposits and withdrawals, demonstrating thread-safe
// operations using mutex locks and channel-based communication.
func main() {
	fmt.Println("Concurrency Bank simulation")

	account := &Account{
		number:  123456,
		balance: 0,
	}

	transactionChannel := make(chan Transaction)
	var wg sync.WaitGroup
	var transactionWG sync.WaitGroup // WaitGroup for transaction goroutines

	wg.Add(1)
	go processTransaction(account, transactionChannel, &wg)

	// Launch 15 goroutines for concurrent transactions
	for i := 0; i <= 15; i++ {
		transactionWG.Add(1)
		go func(id int) {
			defer transactionWG.Done()
			account.Deposit(float64(id*10+50), transactionChannel)
			time.Sleep(time.Millisecond * 100)
			account.Withdraw(float64(id*10+30), transactionChannel)
		}(i)
	}

	// Wait for all transactions to complete
	transactionWG.Wait()

	// Close the channel after all transactions are complete
	close(transactionChannel)

	// Wait for the notification goroutine to finish
	wg.Wait()

	// Display final account balance
	fmt.Printf("Saldo final de la cuenta: %.2f\n", account.balance)
}
