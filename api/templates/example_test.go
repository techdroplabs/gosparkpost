package templates_test

import (
	"encoding/json"
	"fmt"
	"log"

	"bitbucket.org/yargevad/go-sparkpost/api"
	"bitbucket.org/yargevad/go-sparkpost/api/templates"
)

// Build a native Go Template structure from a JSON string
func ExampleFromJSON() {
	cfg := &api.Config{BaseUrl: "https://example.com", ApiKey: "foo"}
	T, err := templates.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	template := &templates.Template{}
	jsonStr := `{
		"name": "testy template",
		"content": {
			"html": "this is a <b>test</b> email!",
			"subject": "test email",
			"from": {
				"name": "tester",
				"email": "tester@example.com"
			},
			"reply_to": "tester@example.com"
		}
	}`
	err = json.Unmarshal([]byte(jsonStr), template)
	if err != nil {
		log.Fatal(err)
	}

	id, err := T.Create(template)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created Template with id [%s]\n", id)
}
