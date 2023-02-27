package validation

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/qri-io/jsonschema"
	apperror "github.com/samriddhadev/go-acme-flights/api/error"
	"github.com/samriddhadev/go-acme-flights/core/config"
	"github.com/samriddhadev/go-acme-flights/core/logger"
)

const (
	SCHEMA_GET_FLIGHTS   = "get-flights"
	SCHEMA_CREATE_FLIGHT = "create-flight"
	SCHEMA_GET_FLIGHT    = "get-flight"
	SCHEMA_UPDATE_FLIGHT = "update-flight"
	SCHEMA_DELETE_FLIGHT = "delete-flight"
)

type ValidationModel struct {
	Header *jsonschema.Schema `json:"header"`
	Query  *jsonschema.Schema `json:"query"`
	URI    *jsonschema.Schema `json:"uri"`
	Body   *jsonschema.Schema `json:"body"`
}

func NewValidator(logger *logger.AcmeLogger) Validator {
	return Validator{
		logger: logger,
	}
}

type Validator struct {
	schema map[string]*ValidationModel
	logger *logger.AcmeLogger
}

func (validator *Validator) loadSchema(key string, cfg *config.Config) (*ValidationModel, error) {
	if validator.schema == nil {
		validator.schema = map[string]*ValidationModel{}
	}
	value, ok := validator.schema[key]
	if !ok {
		validator.logger.Infof("loading schema - %s", key)
		path, ok := cfg.APIValidationSchema[key]
		if ok {
			path = filepath.Join("./resources", path)
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return &ValidationModel{}, fmt.Errorf("missing schema definition file - %s", key)
			} else {
				schema := &ValidationModel{}
				err = json.Unmarshal(content, schema)
				if err != nil {
					return &ValidationModel{}, fmt.Errorf("invalid schema definition - %s", key)
				}
				value = schema
				validator.schema[key] = schema
			}
		} else {
			return &ValidationModel{}, fmt.Errorf("missing schema definition key - %s", key)
		}
	}
	return value, nil
}

func (validator *Validator) validateSchema(schema *ValidationModel, ctx *gin.Context) (bool, *apperror.ValidationError) {
	context := context.Background()
	validationFields := []apperror.ValidationField{}
	if schema != nil {
		if schema.Header != nil {
			ok, err := validator.validateRequestParts(&context, schema.Header, ctx.Request.Header, &validationFields, "header")
			if !ok {
				return false, &apperror.ValidationError{
					Errors: []apperror.ValidationField{{FieldName: "NA", Message: err.Error()}},
				}
			}
		}
		if schema.Query != nil {
			ok, err := validator.validateRequestParts(&context, schema.Query, ctx.Request.URL.Query(), &validationFields, "query")
			if !ok {
				return false, &apperror.ValidationError{
					Errors: []apperror.ValidationField{{FieldName: "NA", Message: err.Error()}},
				}
			}
		}
		if schema.URI != nil {
			ok, err := validator.validateRequestParts(&context, schema.URI, validator.normalizeUri(ctx), &validationFields, "uri")
			if !ok {
				return false, &apperror.ValidationError{
					Errors: []apperror.ValidationField{{FieldName: "NA", Message: err.Error()}},
				}
			}
		}
		if schema.Body != nil {
			bodyData, err := ioutil.ReadAll(ctx.Request.Body)
			if err != nil {
				return false, &apperror.ValidationError{
					Errors: []apperror.ValidationField{{FieldName: "NA", Message: err.Error()}},
				}
			}
			var body interface{}
			err = json.Unmarshal(bodyData, &body)
			if err != nil {
				return false, &apperror.ValidationError{
					Errors: []apperror.ValidationField{{FieldName: "NA", Message: err.Error()}},
				}
			}
			ok, err := validator.validateRequestParts(&context, schema.Body, body, &validationFields, "body")
			if !ok {
				return false, &apperror.ValidationError{
					Errors: []apperror.ValidationField{{FieldName: "NA", Message: err.Error()}},
				}
			}
			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyData))
		}
	}

	if len(validationFields) > 0 {
		return false, &apperror.ValidationError{
			Errors: validationFields,
		}
	}
	return true, &apperror.ValidationError{}
}

func (validator *Validator) validateRequestParts(
	context *context.Context,
	schema *jsonschema.Schema,
	data interface{},
	fields *[]apperror.ValidationField,
	fieldName string) (bool, error) {
	if schema != nil {
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return false, fmt.Errorf("invalid request %s", fieldName)
		}
		keyErrors, err := schema.ValidateBytes(*context, dataBytes)
		if err != nil {
			return false, fmt.Errorf("request %s validation failed", fieldName)
		}
		if len(keyErrors) > 0 {
			for _, keyErr := range keyErrors {
				validationField := apperror.ValidationField{
					FieldName: keyErr.PropertyPath,
					Message:   keyErr.Message,
					Details:   keyErr.InvalidValue,
				}
				*fields = append(*fields, validationField)
			}
		}
	}
	return true, nil
}

func (validator *Validator) normalizeUri(ctx *gin.Context) map[string]interface{} {
	outparam := map[string]interface{}{}
	if ctx.Params != nil {
		for _, param := range ctx.Params {
			outparam[param.Key] = param.Value
		}
	}
	return outparam
}

func (validator *Validator) Validate(key string, cfg *config.Config, handler gin.HandlerFunc) gin.HandlerFunc {
	schema, err := validator.loadSchema(key, cfg)
	return func(ctx *gin.Context) {
		if err != nil {
			ctx.AbortWithError(http.StatusPreconditionFailed, err)
		}
		ok, validationError := validator.validateSchema(schema, ctx)
		if ok {
			handler(ctx)
		} else {
			ctx.AbortWithError(http.StatusBadRequest, validationError)
		}
	}
}
