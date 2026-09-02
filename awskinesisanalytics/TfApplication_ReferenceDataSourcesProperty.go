package awskinesisanalytics


// Experimental.
type TfApplication_ReferenceDataSourcesProperty struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#s3 TfApplication#s3}
	// Experimental.
	S3 *TfApplication_S3Property `field:"required" json:"s3" yaml:"s3"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#schema TfApplication#schema}
	// Experimental.
	Schema *TfApplication_ReferenceDataSourcesSchemaProperty `field:"required" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#table_name TfApplication#table_name}.
	// Experimental.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

