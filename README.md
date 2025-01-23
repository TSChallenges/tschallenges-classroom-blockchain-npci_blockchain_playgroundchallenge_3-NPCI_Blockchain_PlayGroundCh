# Playground-Challenge-3 : Bank Loan Management System on Hyperledger Fabric

## Description
The objective of this project is to build a blockchain-based bank loan management system using Hyperledger Fabric. The system will support loan applications, approvals, repayments, and balance inquiries.

## Getting Started
1. **Install Dependencies**:
   - Docker and Docker Compose
   - Go (latest version)
   - Hyperledger Fabric CLI tools
2. **Start the Network**:
   ```bash
   ./network.sh up createChannel -c mychannel -ca
   ```
3. **Deploy Chaincode**:
   ```bash
   ./network.sh deployCC -ccn loan -ccp ./chaincode -ccl go
   ```
4. **Run the Client Application**:
   ```bash
   go run client/client.go
   ```

## Functionality
- **Apply for a Loan**: Submit loan details to the blockchain.
- **Approve a Loan**: Update loan status (admin-only).
- **Repay a Loan**: Record repayments and update outstanding balance.
- **Check Loan Balance**: Retrieve loan details from the ledger.

## TODO
Refer to `TODO.md` for remaining implementation tasks.

## Submission
Commit your code to the Github code repository
