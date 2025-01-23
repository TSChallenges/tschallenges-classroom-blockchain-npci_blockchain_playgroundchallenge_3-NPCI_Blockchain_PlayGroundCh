# TODOs for Bank Loan Management System

## Chaincode (Go)
1. Implement `ApplyForLoan` to add loan details to the ledger.
2. Implement `ApproveLoan` to update the loan status (e.g., "Approved" or "Rejected").
3. Implement `MakeRepayment` to deduct the repayment amount from the outstanding balance.
4. Implement `CheckLoanBalance` to retrieve loan details from the ledger.

## Client (Go)
1. Call `ApplyForLoan` in `client.go` to submit a loan application.
2. Call `CheckLoanBalance` to retrieve and print loan details.
3. Add functions for `ApproveLoan` and `MakeRepayment` in the client application.
4. Test interactions with the chaincode by applying for a loan, approving it, making a repayment, and checking the balance.

## Testing
1. Test multiple loan applications and ensure the ledger reflects accurate details.
2. Test edge cases, such as:
   - Making a repayment greater than the outstanding balance.
   - Checking the balance of a non-existent loan.
