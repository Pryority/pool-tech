// cmd/server/handlers/debug.go
package handlers

import (
	"bytes"

	"github.com/ethereum/go-ethereum/accounts/abi"
	// ... Other necessary imports ...
)

// Assume you have a variable that holds the contract ABI
var contractABI abi.ABI

// Example contract ABI JSON
var contractJSON = []byte(`[
    {
        "constant": false,
        "inputs": [
            {
                "name": "param1",
                "type": "uint256"
            }
        ],
        "name": "exampleFunction",
        "outputs": [],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    }
]`)

// ContractFunction holds information about a contract function and its inputs
type ContractFunction struct {
	Name   string
	Inputs abi.Arguments
}

func init() {
	var err error
	// Load the contract ABI
	contractABI, err = abi.JSON(bytes.NewReader(contractJSON))
	if err != nil {
		// Handle the error
	}
}

// func LoadDebug(c *fiber.Ctx) error {
// 	// ... Load the contract ABI ...

// 	// Generate forms for each function in the ABI
// 	contractFunctions := generateContractFunctions(contractABI.Methods)

// 	return c.Render("debug", fiber.Map{
// 		"RouteName":         "debug",
// 		"ContractFunctions": contractFunctions,
// 	}, "layouts/main")
// }

// func generateContractFunctions(methods abi.Methods) []ContractFunction {
// 	var contractFunctions []ContractFunction

// 	for _, method := range methods {
// 		function := ContractFunction{
// 			Name:   method.Name,
// 			Inputs: method.Inputs,
// 		}
// 		contractFunctions = append(contractFunctions, function)
// 	}

// 	return contractFunctions
// }
