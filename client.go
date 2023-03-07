package orquesta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// ApiUrl is the URL of the Orquesta API.
const ApiUrl = "https://api.orquesta.dev/evaluate"

var Logger = log.New(io.Discard, "[Orquesta] ", log.LstdFlags)

type Dictionary map[string]interface{}
type RuleContext Dictionary

type ClientOptions struct {
	// ApiKey is the ApiKey to use for the Orquesta API authentication.
	ApiKey string

	// In debug mode, the debug information is printed to stdout to help you
	// understand what orquesta is doing.
	// Debug bool

	// io.Writer implementation that should be used with the Debug mode.
	// DebugWriter io.Writer
}

type Client struct {
	options ClientOptions
}

func NewClient(options ClientOptions) (*Client, error) {
	// if options.Debug {
	// 	debugWriter := options.DebugWriter
	// 	if debugWriter == nil {
	// 		debugWriter = os.Stderr
	// 	}
	// 	Logger.SetOutput(debugWriter)
	// }

	if options.ApiKey == "" {
		options.ApiKey = os.Getenv("ORQUESTA_API_KEY")
	}

	client := Client{
		options: options,
	}

	return &client, nil
}

// Options return ClientOptions for the current Client.
func (client Client) Options() ClientOptions {
	return client.options
}

/*
Query a rule with the given context.

We recommend to use generics so we can infer the type of the result.

`client.Query<int>(ruleKey, context)`
*/
func (c *Client) Query(ruleKey string, context RuleContext, value interface{}) error {
	body, err := json.Marshal(Dictionary{
		"rule_key": ruleKey,
		"context":  context,
	})

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", ApiUrl, bytes.NewBuffer(body))

	if err != nil {
		return err
	}

	options := c.Options()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+options.ApiKey)
	req.Header.Set("X-SDK-Version", userAgent)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var data map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&data)

		if err != nil {
			return fmt.Errorf(resp.Status)
		}

		return fmt.Errorf("%s: %s ", data["detail"], resp.Status)
	}

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return err
	}

	result, ok := data[ruleKey]

	if !ok {
		return fmt.Errorf("%s property cannot be evaluated", ruleKey)
	}

	resultBytes, err := json.Marshal(result)

	if err != nil {
		return err
	}

	err = json.Unmarshal(resultBytes, &value)

	if err != nil {
		return err
	}

	return nil
}
