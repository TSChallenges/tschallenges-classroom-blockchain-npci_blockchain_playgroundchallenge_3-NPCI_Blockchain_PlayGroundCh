# Bank Loan Management System on Hyperledger Fabric

## Description
This project demonstrates a blockchain-based bank loan management system using Hyperledger Fabric. The system supports loan applications, approvals, repayments, and balance inquiries.

## Getting Started

### Install Dependencies
1. Install **Docker** and **Docker Compose**.
2. Install the latest version of **Go**.
3. Install **Hyperledger Fabric CLI tools**.

### Start the Network
1. Navigate to your Fabric test network folder:
   ```bash
   cd fabric-samples/test-network
   ```
2. Start the network and create a channel:
   ```bash
   ./network.sh up createChannel -c mychannel -ca
   ```

### Deploy Chaincode
1. Deploy the loan management chaincode:
   ```bash
   ./network.sh deployCC -ccn loan -ccp ./chaincode -ccl go
   ```
2. Common issues and resolutions:
   - **Incorrect Chaincode Path**: Ensure the path in the `-ccp` parameter points to your chaincode folder.
   - **Chaincode Not Starting**: Check for syntax errors or missing imports in your chaincode.

---

## Key Concepts and Examples

### 1. Using `GetStub().PutState` and `GetStub().GetState`
- **Storing Data**: Use `PutState` to store key-value pairs in the ledger:
  ```go
  err := ctx.GetStub().PutState("loan1", []byte("Loan Data"))
  if err != nil {
      return fmt.Errorf("Failed to store data: %v", err)
  }
  ```
- **Retrieving Data**: Use `GetState` to fetch data by key:
  ```go
  data, err := ctx.GetStub().GetState("loan1")
  if err != nil {
      return fmt.Errorf("Failed to retrieve data: %v", err)
  }
  if data == nil {
      return fmt.Errorf("Loan not found")
  }
  ```

### 2. Submit vs. Evaluate Transactions
- **SubmitTransaction**: For writing data to the ledger (e.g., applying for loans, making repayments).
- **EvaluateTransaction**: For reading data from the ledger (e.g., checking loan balances).
  ```go
  // Submit
  contract.SubmitTransaction("ApplyForLoan", "loan1", "John Doe", "5000", "12", "5.5")

  // Evaluate
  result, _ := contract.EvaluateTransaction("CheckLoanBalance", "loan1")
  fmt.Printf("Loan details: %s
", string(result))
  ```

### 3. Test Scenarios
- Validate inputs:
  - Reject loans with missing or invalid fields (e.g., negative loan amount).
  - Reject repayments exceeding outstanding balance.
- Handle edge cases:
  - Attempt to retrieve details of a non-existent loan.
  - Approve loans only in the "Pending" status.

---

## Client Configuration
1. **Sample `connection-org1.yaml` File**:
   Ensure the client connects to the network:
   ```yaml
   name: mychannel
   version: 1.0.0
   client:
     organization: Org1
   ```
2. **Wallet Setup**:
   Use a pre-created wallet identity (e.g., `appUser`):
   ```go
   gw, err := gateway.Connect(
       gateway.WithConfig(ccpPath),
       gateway.WithIdentity(wallet, "appUser"),
   )
   ```

---

## Testing Guide

### Example Transactions
1. **Apply for a Loan**:
   ```bash
   SubmitTransaction("ApplyForLoan", "loan1", "John Doe", "5000", "12", "5.5")
   ```
   Expected Output:
   ```
   Loan successfully applied
   ```

2. **Approve a Loan**:
   ```bash
   SubmitTransaction("ApproveLoan", "loan1", "Approved")
   ```
   Expected Output:
   ```
   Loan status updated to Approved
   ```

3. **Repay a Loan**:
   ```bash
   SubmitTransaction("MakeRepayment", "loan1", "1000")
   ```
   Expected Output:
   ```
   Repayment recorded. Outstanding balance updated.
   ```

4. **Check Loan Balance**:
   ```bash
   EvaluateTransaction("CheckLoanBalance", "loan1")
   ```
   Expected Output:
   ```
   Outstanding Balance: 4000
   ```

---

## Submission
Commit your completed code to the GitHub Classroom repository provided in your assignment link.
