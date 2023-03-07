/*
Package orquesta is the official Orquesta SDK for Go.

Use this package to integrate your Go application with Orquesta and evaluate your rules.


# Basic Usage

The first step is to initialize the SDK, providing at a minimum the ApiKey of your Orquesta workspace

	func main() {
		err := orquesta.Init(...)
		...
	}

# Rule evaluation

The Query function is used to evaluate a rule with a given context.
	var value string
	value, err := client.Query("rule_id", orquesta.RuleContext{
		"key": "value",
	}, value)
*/

package orquesta
