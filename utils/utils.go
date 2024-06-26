package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/mrz1836/go-sanitize"
	"golang.org/x/exp/rand"
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

func GenerateReferenceNumber(prefex string) string {
	timestamp := time.Now().UTC().Format("20060102150405")
	uniqueID := fmt.Sprintf("%s%04d", timestamp, time.Now().Nanosecond())
	generatedRefNumber := prefex + "_" + uniqueID[14:18] + strconv.Itoa(RandomInteger(3))

	return generatedRefNumber
}

func RandomInteger(n int) int {
	var sb strings.Builder
	ran := rand.Int63n(999999999999999999)
	strng := strconv.FormatInt(ran, 10)
	k := len(strng)
	for i := 0; i < n; i++ {
		c := strng[rand.Intn(k)]
		sb.WriteByte(c)
	}
	finalInt, _ := strconv.Atoi(sb.String())
	return finalInt
}
