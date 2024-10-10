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
// https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html
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

type S3ObjectNameOptions struct {
	SafeCharactersOnly  bool
	AWSConsoleSafe      bool
	AWSProgrammaticSafe bool
}

// Default S3ObjectNameOptions
var DefaultS3ObjectNameOptions = S3ObjectNameOptions{
	SafeCharactersOnly:  true,
	AWSConsoleSafe:      true,
	AWSProgrammaticSafe: true,
}

// S3ObjectName validates an S3 object name
// The options provided encode generally accepted best practices for S3 object naming
// https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html
func S3ObjectName(name string, opts ...S3ObjectNameOptions) error {
	opts = append(opts, DefaultS3ObjectNameOptions)
	if len(name) < 1 || len(name) > 1024 {
		return errors.New("object name must be between 1 and 1024 characters long")
	}

	// safe_characters_set := []string{"0-9", "a-z", "A-Z", "!_.*'()-"}
	safe_characters_regex := regexp.MustCompile(`^[0-9a-zA-Z!_.*'()-]+$`)
	if len(opts) > 0 && opts[0].SafeCharactersOnly && !safe_characters_regex.MatchString(name) {
		return errors.New("object name can only contain the following characters: 0-9, a-z, A-Z, !, _, ., *, ', (, ) and -")
	}

	// console safe validations
	// 1. should not end with a dot
	// 2. should not have a prefix of ./ or ../
	if len(opts) > 0 && opts[0].AWSConsoleSafe {
		if strings.HasSuffix(name, ".") {
			return errors.New("object name should not end with a dot")
		}
		if strings.HasPrefix(name, "./") || strings.HasPrefix(name, "../") {
			return errors.New("object name should not have a prefix of ./ or ../")
		}
	}

	// validations for SDK and CLI downloading of objects
	// 1. should not have a prefix of ../
	if len(opts) > 0 && opts[0].AWSProgrammaticSafe {
		if strings.HasPrefix(name, "../") {
			return errors.New("object name should not have a prefix of ../")
		}
	}

	return nil
}
