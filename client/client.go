package main

import (
	"fmt"
	"log"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

func main() {
	walletPath := "./wallet"
	ccpPath := "./connection-org1.yaml"

	wallet, err := gateway.NewFileSystemWallet(walletPath)
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	gw, err := gateway.Connect(
		gateway.WithConfig(ccpPath),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}

	contract := network.GetContract("loan")

	// TODO: Call ApplyForLoan
	_, err = contract.SubmitTransaction("ApplyForLoan", "loan1", "John Doe", "5000", "12", "5.5")
	if err != nil {
		log.Fatalf("Failed to apply for loan: %v", err)
	}
	fmt.Println("Loan successfully applied.")

	// TODO: Call CheckLoanBalance
	result, err := contract.EvaluateTransaction("CheckLoanBalance", "loan1")
	if err != nil {
		log.Fatalf("Failed to get loan balance: %v", err)
	}
	fmt.Printf("Loan details: %s
", string(result))
}
