package s3


// Experimental.
type AwsBucketWebsiteConfiguration_ConditionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_website_configuration#http_error_code_returned_equals AwsBucketWebsiteConfiguration#http_error_code_returned_equals}.
	// Experimental.
	HttpErrorCodeReturnedEquals *string `field:"optional" json:"httpErrorCodeReturnedEquals" yaml:"httpErrorCodeReturnedEquals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_website_configuration#key_prefix_equals AwsBucketWebsiteConfiguration#key_prefix_equals}.
	// Experimental.
	KeyPrefixEquals *string `field:"optional" json:"keyPrefixEquals" yaml:"keyPrefixEquals"`
}

