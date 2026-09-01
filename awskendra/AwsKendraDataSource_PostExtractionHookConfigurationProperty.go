package awskendra


// Experimental.
type AwsKendraDataSource_PostExtractionHookConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#lambda_arn AwsKendraDataSource#lambda_arn}.
	// Experimental.
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#s3_bucket AwsKendraDataSource#s3_bucket}.
	// Experimental.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// invocation_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#invocation_condition AwsKendraDataSource#invocation_condition}
	// Experimental.
	InvocationCondition *AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionProperty `field:"optional" json:"invocationCondition" yaml:"invocationCondition"`
}

