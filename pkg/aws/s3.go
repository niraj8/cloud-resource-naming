package aws

import (
	"errors"
	"regexp"
	"strings"
)

type S3BucketNameOptions struct {
	TransferAccelerationEnabled bool
}

// S3BucketName validates an S3 bucket name for general purpose buckets
func S3BucketName(name string, opts ...S3BucketNameOptions) error {
	if len(name) < 3 || len(name) > 63 {
		return errors.New("bucket name must be between 3 and 63 characters long")
	}

	if !regexp.MustCompile(`^[a-z0-9][a-z0-9.-]*[a-z0-9]$`).MatchString(name) {
		return errors.New("bucket name can consist only of lowercase letters, numbers, dots (.), and hyphens (-), and must begin and end with a letter or number")
	}

	if strings.Contains(name, "..") {
		return errors.New("bucket name must not contain two adjacent periods")
	}

	if regexp.MustCompile(`^\d+\.\d+\.\d+\.\d+$`).MatchString(name) {
		return errors.New("bucket name must not be formatted as an IP address")
	}

	if strings.HasPrefix(name, "xn--") {
		return errors.New("bucket name must not start with the prefix xn--")
	}

	if strings.HasPrefix(name, "sthree-") {
		return errors.New("bucket name must not start with the prefix sthree-")
	}

	if strings.HasPrefix(name, "amzn-s3-demo-") {
		return errors.New("bucket name must not start with the prefix amzn-s3-demo-")
	}

	if strings.HasSuffix(name, "-s3alias") {
		return errors.New("bucket name must not end with the suffix -s3alias")
	}

	if strings.HasSuffix(name, "--ol-s3") {
		return errors.New("bucket name must not end with the suffix --ol-s3")
	}

	if strings.HasSuffix(name, ".mrap") {
		return errors.New("bucket name must not end with the suffix .mrap")
	}

	if strings.HasSuffix(name, "--x-s3") {
		return errors.New("bucket name must not end with the suffix --x-s3")
	}

	if len(opts) > 0 && opts[0].TransferAccelerationEnabled {
		if strings.Contains(name, ".") {
			return errors.New("buckets used with Amazon S3 Transfer Acceleration can't have dots (.) in their names")
		}
	}

	return nil
}
