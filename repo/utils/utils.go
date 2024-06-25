package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/mrz1836/go-sanitize"
)

func BindingFormError(err error, model any) string {
	errMessage := err.Error()
	if parts := strings.Split(errMessage, " "); len(parts) >= 2 {
		field := parts[1]
		fieldList := strings.Split(field, ".")
		if len(fieldList) > 1 {
			fieldExtracted := sanitize.Alpha(fieldList[1], true)
			ref, found := reflect.TypeOf(model).FieldByName(fieldExtracted)
			if !found {
				return "Invalid input"
			}

			form := ref.Tag.Get("form")
			value := strings.ReplaceAll(form, "_", " ")
			return fmt.Sprintf("Field '" + value + "' is required")
		} else {
			return errMessage
		}
	}

	return "Invalid input"
}
