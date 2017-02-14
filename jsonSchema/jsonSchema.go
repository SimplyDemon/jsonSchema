package jsonSchema

import (
	"github.com/tidwall/gjson"
	"github.com/xeipuuv/gojsonschema"
)

type Json struct {
	String string
}

func (json *Json) SetString(setString string) {
	json.String = setString
}

func (json *Json) IsContain(jsonPath string, value string) bool {
	err := gjson.Get(json.String, jsonPath)     // json path its where we finding ("err")
	for _, name := range err.Array() {
		if name.String() == value {
			// value its which error we expecting ("PLATFORM_REQUIRED")
			return true
		}
	}
	return false
}

func (json *Json) IsValidBySchema(jsonSchemaPath string) bool {

	schemaLoader := gojsonschema.NewReferenceLoader(jsonSchemaPath)            //json schema its absolutely path to jsonSchema (like: "file:///C:/Users/sd/IdeaProjects/untitled1/jsonWrongSchemaTest.json")
	documentLoader := gojsonschema.NewStringLoader(json.String)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return false
	}
	return result.Valid()
}