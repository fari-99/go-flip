# Go-Flip - Unofficial Flip Golang Interface

# Descriptions
This package is interface for api https://docs.flip.id/, you only need to add ENV and add this package to your project.

# Features
1. Have collections of constants that used in the flip environment
2. Easy to use, doesn't need to add authentication every function, the package handle it by getting it from ENV.
# Endpoints
- [Generals](https://docs.flip.id/#general)
	- Get Balance
	- Get Bank Info
	- Is Maintenance
	- Bank Account Inquiry (failed testing)
- [Money Transfer](https://docs.flip.id/#money-transfer)
	- Create Disbursement
	- Get All Disbursement
	- Get Disbursement By Idempotency Key (failed testing)
	- Get Disbursement By ID (failed testing)
- [Special Money Transfer](https://docs.flip.id/#special-money-transfer)
	- Create Special Disbursement
	- City List
	- Country List
	- City and Country List
- [Agent Money Transfer](https://docs.flip.id/#agent-money-transfer)
	- Create Disbursement for Agent
	- Get Agent Disbursement by ID
	- Get Agent Disbursement List
- [Agent Verification](https://docs.flip.id/#agent-verification)
	- Create Agent Identity
	- Update Agent Identity
	- Get Agent Identity
	- Upload Agent Identity Image
	- Upload Supporting Documents
	- KYC Submission
	- Repair Data
	- Repair Identity Image
	- Repair Identity Selfie Image
	- Get Country List
	- Get Province List
	- Get City List
	- Get District List
- [Accept Payment](https://docs.flip.id/#accept-payment)
	- Create Bill
	- Edit Bill
	- Get Bill
	- Get All Bills
	- Get Payment
	- Get All Payment
	- Confirm Bill Payment (**ONLY FOR PRODUCTION**)
- [International Transfer](https://docs.flip.id/#international-transfer)
	- Get Exchange Rates
	- Get Form Data
	- Create International Transfer C2C/C2B
	- Create International Transfer B2C/B2B (**NOT FOUND** need confirmation)
	- Get International Transfer
	- Get All International Transfer

# TODO
- Validation, flip already have validation on their end, but it will be more helpfull if we validate first before sending to flip
- Documentations
- Unit Tests

# How To Use

### Getting Flip

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```
import "github.com/fari-99/go-flip"
```

to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the `flip` package:

```sh
$ go get -u github.com/fari-99/go-flip
```

Add Environment variable to your project
```dotenv
FLIP_ENVIRONMENT=dev
FLIP_SECRET_TOKEN="YOUR FLIP SECRET TOKEN"
FLIP_VALIDATION_TOKEN="YOUR FLIP VALIDATION TOKEN"
```

### Example
#### Check if maintenance

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/fari-99/go-flip"
)

func main() {
	baseFlip := flip.NewBaseFlip()
	isMaintenance, err := baseFlip.IsMaintenance()
	if err != nil {
		fmt.Println(err.Error())
		errorMarshal, _ := json.MarshalIndent(baseFlip.GetErrorDetails(), "", " ")
		fmt.Println(string(errorMarshal))
		return
	}
	
	fmt.Printf("%+v", isMaintenance)
	return
}
```

#### Get Balance

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/fari-99/go-flip"
)

func main() {
	baseFlip := flip.NewBaseFlip()
	balance, err := baseFlip.GetBalance()
	if err != nil {
		fmt.Println(err.Error())
		errorMarshal, _ := json.MarshalIndent(baseFlip.GetErrorDetails(), "", " ")
		fmt.Println(string(errorMarshal))
		return
	}
	
	fmt.Printf("%+v", balance)
	return
}
```
