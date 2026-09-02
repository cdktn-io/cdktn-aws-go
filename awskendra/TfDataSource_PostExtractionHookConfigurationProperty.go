package awskendra


// Experimental.
type TfDataSource_PostExtractionHookConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#lambda_arn TfDataSource#lambda_arn}.
	// Experimental.
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#s3_bucket TfDataSource#s3_bucket}.
	// Experimental.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// invocation_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#invocation_condition TfDataSource#invocation_condition}
	// Experimental.
	InvocationCondition *TfDataSource_CustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionProperty `field:"optional" json:"invocationCondition" yaml:"invocationCondition"`
}

