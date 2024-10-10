package aws_test

import (
	"strings"
	"testing"

	"github.com/niraj8/cloud-resource-naming/pkg/aws"
)

// TestS3Validator tests the S3Validator
func TestS3Validator(t *testing.T) {

	tests := []struct {
		name      string
		bucket    string
		expectErr bool
	}{
		{"Valid bucket name", "valid-bucket-name", false},
		{"Too short", "ab", true},
		{"Too long", strings.Repeat("a", 64), true},
		{"Invalid characters", "Invalid_Bucket_Name", true},
		{"Adjacent periods", "invalid..bucket", true},
		{"IP address format", "192.168.0.1", true},
		{"Prefix xn--", "xn--bucket", true},
		{"Prefix sthree-", "sthree-bucket", true},
		{"Prefix sthree-configurator", "sthree-configurator-bucket", true},
		{"Prefix amzn-s3-demo-", "amzn-s3-demo-bucket", true},
		{"Suffix -s3alias", "bucket-s3alias", true},
		{"Suffix --ol-s3", "bucket--ol-s3", true},
		{"Suffix .mrap", "bucket.mrap", true},
		{"Suffix --x-s3", "bucket--x-s3", true},
		{"Contains dot", "bucket.name", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := aws.S3BucketName(tt.bucket)
			if (err != nil) != tt.expectErr {
				t.Errorf("Validate() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}
