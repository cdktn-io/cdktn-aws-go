package awskinesisanalytics


// Experimental.
type AwsKinesisAnalyticsApplication_ReferenceDataSourcesProperty struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#s3 AwsKinesisAnalyticsApplication#s3}
	// Experimental.
	S3 *AwsKinesisAnalyticsApplication_S3Property `field:"required" json:"s3" yaml:"s3"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#schema AwsKinesisAnalyticsApplication#schema}
	// Experimental.
	Schema *AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaProperty `field:"required" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#table_name AwsKinesisAnalyticsApplication#table_name}.
	// Experimental.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

