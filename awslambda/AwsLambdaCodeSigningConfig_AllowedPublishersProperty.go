package awslambda


// Experimental.
type AwsLambdaCodeSigningConfig_AllowedPublishersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_code_signing_config#signing_profile_version_arns AwsLambdaCodeSigningConfig#signing_profile_version_arns}.
	// Experimental.
	SigningProfileVersionArns *[]*string `field:"required" json:"signingProfileVersionArns" yaml:"signingProfileVersionArns"`
}

