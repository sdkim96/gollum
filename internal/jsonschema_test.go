package internal

import (
	"fmt"
	"testing"
)

func TestJsonSchema(t *testing.T) {
	type Sample struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	jsonschema := JsonSchema[Sample]("sample_tool", "A tool for parsing the model's response into a structured format. The model should use this tool to return the final output in a structured way.")
	fmt.Print(jsonschema)
}
