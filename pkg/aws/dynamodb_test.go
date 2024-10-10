package aws_test

import (
	"testing"

	"github.com/niraj8/cloud-resource-naming/pkg/aws"
)

func TestDynamoDBTableName(t *testing.T) {
	tests := []struct {
		name    string
		table   string
		wantErr bool
	}{
		{"Valid table name", "validTableName", false},
		{"Too short", "ab", true},
		{"Too long", string(make([]byte, 256)), true},
		{"Contains space", "Invalid Table", true},
		{"Invalid characters", "Invalid🫥", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := aws.DynamoDBTableName(tt.table)
			if (err != nil) != tt.wantErr {
				t.Errorf("DynamoDBTableName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDynamoDBIndexName(t *testing.T) {
	tests := []struct {
		name    string
		index   string
		wantErr bool
	}{
		{"Valid index name", "validIndexName", false},
		{"Too short", "", true},
		{"Too long", string(make([]byte, 256)), true},
		{"Contains space", "Invalid Index", true},
		{"Invalid characters", "Invalid🫥", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := aws.DynamoDBIndexName(tt.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("DynamoDBIndexName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
