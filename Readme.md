<p align="left">
  <a href="https://orquesta.dev" target="_blank">
    <img src="https://static.wixstatic.com/media/e063e5_4f60988535a643218a02ad84cf60b7cd~mv2.png/v1/fill/w_130,h_108,al_c,q_85,usm_0.66_1.00_0.01,enc_auto/Logo%2001.png" alt="Orquesta"  height="84">
  </a>
</p>

# Official Orquesta SDK for GO

## Installation

```bash
go github.com/orquestadev/orquesta-go
```

## Usage

_You can get your workspace API key from the settings section in your workspace._

```go
package main

import (
    "fmt"
    "github.com/orquestadev/orquesta-go"
)

func main() {
	client, err := orquesta.Init(orquesta.ClientOptions{
		ApiKey: "ORQUESTA_API_KEY",
	})

    if err != nil {
		  // ...
	  }

    var kill_switch_enabled bool
    err = client.Query(
    	"kill_switch",
    	orquesta.RuleContext{"environments": "production"},
    	&kill_switch_enabled
    )


    if err != nil {
		  // ...
    }

    fmt.Printf("Result: %v\n", kill_switch_enabled)
}

```

## Notes

Is important to note that the value provided as the third parameter to the `Query` method must be a pointer to a variable of the type you want to receive the result in. This is because the SDK uses reflection to determine the type of the variable and then unmarshal the result into it.

It is easy to find the type if you look at the definition of the rule in the Orquesta Dashboard.

For example, if you want to receive the result as a `string`, you must pass a pointer to a `string` variable, like this:

```go
var string_var string

err = client.Query("string_rule", orquesta.RuleContext{"environments": "production"}, &string_var)

if err != nil {
    // ...
}
```
