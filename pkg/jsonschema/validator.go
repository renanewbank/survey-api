package jsonschema

import (
	"github.com/xeipuuv/gojsonschema"
)

func ValidateQuestionJSON(jsonData []byte, schemaPath string) ([]string, error) {
	schemaLoader := gojsonschema.NewReferenceLoader("file://" + schemaPath)
	documentLoader := gojsonschema.NewBytesLoader(jsonData)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return nil, err
	}

	if result.Valid() {
		return nil, nil
	}

	var errors []string
	for _, err := range result.Errors() {
		errors = append(errors, err.String())
	}
	return errors, nil
}
