package aws

import (
	"errors"
	"fmt"
	"regexp"
)

func DynamoDBTableName(name string) error {
	if len(name) < 3 || len(name) > 255 {
		return errors.New("dynamodb table name must be between 3 and 255 characters long")
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9_.-]+$`).MatchString(name) {
		return errors.New("dynamodb table name can only contain a-z, A-Z, 0-9, -, _, . ")
	}

	return nil
}

func DynamoDBIndexName(name string) error {
	if len(name) < 1 || len(name) > 255 {
		return errors.New("dynamodb index name must be between 1 and 255 characters long")
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(name) {
		return fmt.Errorf("dynamodb index name can only contain a-z, A-Z, 0-9, _")
	}

	return nil
}
