# flip - Unofficial Flip Golang Interface

# TODO
- Validation before send to Flip
- Documentations
- Unit Tests

# How To Use

### Getting Flip

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```
import "github.com/fari-99/flip"
```

to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the `flip` package:

```sh
$ go get -u github.com/fari-99/flip
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

	"github.com/fari-99/flip"
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

	"github.com/fari-99/flip"
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
