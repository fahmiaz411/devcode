package message

import (
	"fmt"
	"strings"
)

const (
	Success string = "Success"
	InvalidRequestBody = "Invalid Request Body"
)

func NotFound(name, property, value string) string {
	return fmt.Sprintf("%s with %s %s Not Found", name, property, value)
}

func CanotNull(property string) string {
	return fmt.Sprintf("%s cannot be null", property)
}

func InvalidId(name string) string {
	return fmt.Sprintf("Invalid %s Id", name)
}

func ShoudMatchEnum(property string, enums []string) string {
	return fmt.Sprintf("field %s should match one of: %s", property, strings.Join(enums, ", "))
}