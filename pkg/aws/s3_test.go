package aws_test

import (
	"strings"
	"testing"

	"github.com/niraj8/cloud-resource-naming/pkg/aws"
)

// TestS3BucketName tests the S3BucketName function
func TestS3BucketName(t *testing.T) {
	tests := []struct {
		name      string
		bucket    string
		opts      []aws.S3BucketNameOptions
		expectErr bool
	}{
		{"Valid bucket name", "valid-bucket-name", nil, false},
		{"Too short", "ab", nil, true},
		{"Too long", strings.Repeat("a", 64), nil, true},
		{"Uppercase characters", "InvalidBucketName", nil, true},
		{"doc_example_bucket", "doc_example_bucket", nil, true},
		{"Adjacent periods", "invalid..bucket", nil, true},
		{"IP address format", "192.168.0.1", nil, true},
		{"Prefix xn--", "xn--bucket", nil, true},
		{"Prefix sthree-", "sthree-bucket", nil, true},
		{"Prefix sthree-configurator", "sthree-configurator-bucket", nil, true},
		{"Prefix amzn-s3-demo-", "amzn-s3-demo-bucket", nil, true},
		{"Suffix -s3alias", "bucket-s3alias", nil, true},
		{"Suffix --ol-s3", "bucket--ol-s3", nil, true},
		{"Suffix .mrap", "bucket.mrap", nil, true},
		{"Suffix --x-s3", "bucket--x-s3", nil, true},
		{"Valid with dots", "my.bucket.name", nil, false},
		{"Invalid name", "amzn-s3-demo-bucket-", nil, true},
		{"Transfer acceleration enabled with dots", "my.bucket.name", []aws.S3BucketNameOptions{{TransferAccelerationEnabled: true}}, true},
		{"Transfer acceleration enabled without dots", "mybucketname", []aws.S3BucketNameOptions{{TransferAccelerationEnabled: true}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := aws.S3BucketName(tt.bucket, tt.opts...)
			if (err != nil) != tt.expectErr {
				t.Errorf("S3BucketName() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}

func TestS3ObjectNameDefaultOptions(t *testing.T) {
	tests := []struct {
		name      string
		object    string
		opts      []aws.S3ObjectNameOptions
		expectErr bool
	}{
		{"Valid object name", "valid-object-name.txt", nil, false},
		{"Too long", strings.Repeat("a", 1025), nil, true},
		{"Empty name", "", nil, true},
		{"Unsafe characters", "object/with/slash.txt", nil, true},
		{"Ending with dot", "object-name.", nil, true},
		{"Starting with ./", "./object-name.txt", nil, true},
		{"Starting with ../", "../object-name.txt", nil, true},
		{"Valid with special characters", "ABCabc123!_.*'()-", nil, false},
		{"Invalid character", "object name with space.txt", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := aws.S3ObjectName(tt.object, tt.opts...)
			if (err != nil) != tt.expectErr {
				t.Errorf("S3ObjectName() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}

func TestS3ObjectNameCustomOptions(t *testing.T) {
	tests := []struct {
		name      string
		object    string
		opts      []aws.S3ObjectNameOptions
		expectErr bool
	}{
		{"Unsafe characters allowed", "object/with/slash.txt", []aws.S3ObjectNameOptions{{SafeCharactersOnly: false}}, false},
		{"Console unsafe allowed", "object-name.", []aws.S3ObjectNameOptions{{AWSConsoleSafe: false}}, false},
		{"Starting with ./", "./object-name.txt", []aws.S3ObjectNameOptions{{AWSConsoleSafe: true}}, true},
		{"Programmatic unsafe not allowed", "../object-name.txt", []aws.S3ObjectNameOptions{{AWSProgrammaticSafe: true}}, true},
		{"All options disabled", "object name with space.txt", []aws.S3ObjectNameOptions{{SafeCharactersOnly: false, AWSConsoleSafe: false, AWSProgrammaticSafe: false}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := aws.S3ObjectName(tt.object, tt.opts...)
			if (err != nil) != tt.expectErr {
				t.Errorf("S3ObjectName() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}
