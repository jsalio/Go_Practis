# Concurrent Banking System in Go

This exercise demonstrates the implementation of a thread-safe banking system using Go's concurrency features. The program simulates multiple concurrent transactions on a bank account while ensuring data consistency and proper synchronization.

## Overview

The program implements a simple banking system that can handle multiple concurrent deposits and withdrawals while maintaining data integrity. It showcases several important Go concurrency patterns and best practices.

## Key Features

- Thread-safe account operations using mutex locks
- Concurrent transaction processing using goroutines
- Channel-based communication for transaction notifications
- Proper synchronization using WaitGroups
- Error handling for failed transactions

## Implementation Details

### Core Components

1. **Account Structure**
   ```go
   type Account struct {
       number  int
       balance float64
       mutex   sync.Mutex
   }
   ```
   - Represents a bank account with thread-safe balance operations
   - Uses mutex to prevent race conditions

2. **Transaction Types**
   ```go
   type TransactionType string

   const (
       Deposit       TransactionType = "Deposit"
       Withdraw      TransactionType = "Withdraw"
       FailedWitdraw TransactionType = "FailedWitdraw"
   )
   ```
   - Defines possible transaction types
   - Includes handling for failed withdrawals

3. **Transaction Structure**
   ```go
   type Transaction struct {
       Type        TransactionType
       Amount      float64
       FinalAmount float64
   }
   ```
   - Records transaction details
   - Tracks final balance after each operation

### Main Operations

1. **Deposit Operation**
   - Thread-safe method to add funds
   - Uses mutex lock for synchronization
   - Sends transaction record through channel

2. **Withdraw Operation**
   - Thread-safe method to remove funds
   - Includes balance verification
   - Handles failed withdrawal attempts
   - Sends transaction status through channel

3. **Transaction Processing**
   - Handles transaction notifications
   - Processes transactions from channel
   - Provides real-time transaction updates

## Concurrency Patterns Used

1. **Mutex Locks**
   - Protects shared resources (account balance)
   - Ensures thread-safe operations
   - Prevents race conditions

2. **Goroutines**
   - Enables concurrent transaction processing
   - Simulates multiple users accessing the account
   - Demonstrates parallel execution

3. **Channels**
   - Facilitates communication between goroutines
   - Handles transaction notifications
   - Ensures ordered processing

4. **WaitGroups**
   - Synchronizes goroutine execution
   - Ensures all transactions complete
   - Prevents premature program termination

## Running the Program

The program creates an account and spawns multiple goroutines to perform concurrent transactions. Each goroutine:
1. Makes a deposit
2. Waits briefly
3. Attempts a withdrawal

The program demonstrates:
- Safe concurrent operations
- Proper error handling
- Real-time transaction notifications
- Final balance verification

## Learning Outcomes

This exercise helps understand:
- Go's concurrency model
- Thread-safe programming
- Channel-based communication
- Synchronization patterns
- Error handling in concurrent systems

## Best Practices Demonstrated

1. **Resource Protection**
   - Using mutex for shared resources
   - Proper lock/unlock patterns
   - Deferred mutex unlocking

2. **Error Handling**
   - Graceful handling of failed withdrawals
   - Clear error notifications
   - Transaction status tracking

3. **Clean Synchronization**
   - Proper use of WaitGroups
   - Channel closure after completion
   - Orderly program termination

## Conclusion

This exercise provides a practical example of implementing a concurrent system in Go, demonstrating important concepts like:
- Thread safety
- Concurrent programming
- Channel-based communication
- Proper synchronization
- Error handling

It serves as a good foundation for understanding more complex concurrent systems and real-world banking applications. 