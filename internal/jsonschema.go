package internal

import "encoding/json"

type Schema interface {
	schemaType() string
}

type baseSchema struct {
	Type        string `json:"type"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

func (b baseSchema) schemaType() string {
	return b.Type
}

type ArraySchema struct {
	baseSchema
	Items    *Items `json:"items,omitempty"`
	MinItems *int   `json:"minItems,omitempty"`
	MaxItems *int   `json:"maxItems,omitempty"`
}
type BooleanSchema struct {
	baseSchema
}
type NullSchema struct {
	baseSchema
}
type StringSchema struct {
	baseSchema
	MinLength *int    `json:"minLength,omitempty"`
	MaxLength *int    `json:"maxLength,omitempty"`
	Pattern   *string `json:"pattern,omitempty"`
}

type Items struct {
	Schema Schema
	Bool   *bool
	Tuple  []Schema
}

func JsonSchema[T any](name, desc string) json.RawMessage {
	return json.RawMessage(jsonschema[T](name, desc))
}

func jsonschema[T any](name, desc string) []byte {
	return nil
}
