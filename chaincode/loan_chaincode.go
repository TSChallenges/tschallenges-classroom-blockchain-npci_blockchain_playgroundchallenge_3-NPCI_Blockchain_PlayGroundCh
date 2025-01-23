package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type LoanContract struct {
	contractapi.Contract
}

type Loan struct {
	LoanID        string   `json:"loanID"`
	ApplicantName string   `json:"applicantName"`
	LoanAmount    float64  `json:"loanAmount"`
	TermMonths    int      `json:"termMonths"`
	InterestRate  float64  `json:"interestRate"`
	Outstanding   float64  `json:"outstanding"`
	Status        string   `json:"status"`
	Repayments    []float64 `json:"repayments"`
}

// TODO: Implement ApplyForLoan
func (c *LoanContract) ApplyForLoan(ctx contractapi.TransactionContextInterface, loanID, applicantName string, loanAmount float64, termMonths int, interestRate float64) error {
	return nil
}

// TODO: Implement ApproveLoan
func (c *LoanContract) ApproveLoan(ctx contractapi.TransactionContextInterface, loanID string, status string) error {
	return nil
}

// TODO: Implement MakeRepayment
func (c *LoanContract) MakeRepayment(ctx contractapi.TransactionContextInterface, loanID string, repaymentAmount float64) error {
	return nil
}

// TODO: Implement CheckLoanBalance
func (c *LoanContract) CheckLoanBalance(ctx contractapi.TransactionContextInterface, loanID string) (*Loan, error) {
	return nil, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(LoanContract))
	if err != nil {
		fmt.Printf("Error creating loan chaincode: %s", err)
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting loan chaincode: %s", err)
	}
}
