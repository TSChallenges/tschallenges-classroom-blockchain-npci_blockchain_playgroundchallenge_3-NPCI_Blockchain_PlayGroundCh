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
	LoanID        string    `json:"loanID"`
	ApplicantName string    `json:"applicantName"`
	LoanAmount    float64   `json:"loanAmount"`
	TermMonths    int       `json:"termMonths"`
	InterestRate  float64   `json:"interestRate"`
	Outstanding   float64   `json:"outstanding"`
	Status        string    `json:"status"`
	Repayments    []float64 `json:"repayments"`
}

// TODO: Implement ApplyForLoan
func (c *LoanContract) ApplyForLoan(ctx contractapi.TransactionContextInterface, loanID, applicantName string, loanAmount float64, termMonths int, interestRate float64) error {
	existingLoan, err := ctx.GetStub().GetState(loanID)
	if err != nil {
		return fmt.Errorf("failed to check existing loan: %v", err)
	}
	if existingLoan != nil {
		return fmt.Errorf("loan with ID %s already exists", loanID)
	}
	loadReq := Loan{
		LoanID:        loanID,
		ApplicantName: applicantName,
		LoanAmount:    loanAmount,
		TermMonths:    termMonths,
		InterestRate:  interestRate,
	}
	loanReqPayload, err := json.Marshal(loadReq)
	if err != nil {
		return fmt.Errorf("error marshalling load request")
	}

	return ctx.GetStub().PutState(loanID, loanReqPayload)
}

// TODO: Implement ApproveLoan
func (c *LoanContract) ApproveLoan(ctx contractapi.TransactionContextInterface, loanID string, status string) error {
	loanReqPayload, err := ctx.GetStub().GetState(loanID)
	if err != nil {
		return fmt.Errorf("failed while fetching loan")
	}

	var loan Loan
	err = json.Unmarshal(loanReqPayload, &loan)
	if err != nil {
		return err
	}

	if status != "Approved" && status != "Rejected" {
		return fmt.Errorf("invalid status: must be 'Approved' or 'Rejected'")
	}
	loan.Status = status

	updatedLoan, err := json.Marshal(loan)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(loanID, updatedLoan)
}

// TODO: Implement MakeRepayment
func (c *LoanContract) MakeRepayment(ctx contractapi.TransactionContextInterface, loanID string, repaymentAmount float64) error {
	loanReq, err := ctx.GetStub().GetState(loanID)
	if err != nil {
		return fmt.Errorf("could not get loan")
	}

	var loan Loan
	err = json.Unmarshal(loanReq, &loan)
	if err != nil {
		return err
	}

	loan.Outstanding -= repaymentAmount
	loan.Repayments = append(loan.Repayments, repaymentAmount)

	loanPyaload, err := json.Marshal(loan)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(loanID, loanPyaload)
}

// TODO: Implement CheckLoanBalance
func (c *LoanContract) CheckLoanBalance(ctx contractapi.TransactionContextInterface, loanID string) (*Loan, error) {
	loanReq, err := ctx.GetStub().GetState(loanID)
	if err != nil {
		return nil, fmt.Errorf("could not fetch loan")
	}

	var loan Loan
	err = json.Unmarshal(loanReq, &loan)
	if err != nil {
		return nil, err
	}

	return &loan, nil
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
