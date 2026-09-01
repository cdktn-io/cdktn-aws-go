package awskinesisanalyticsv2


// Experimental.
type AwsKinesisanalyticsv2Application_ReferenceDataSourceProperty struct {
	// reference_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#reference_schema AwsKinesisanalyticsv2Application#reference_schema}
	// Experimental.
	ReferenceSchema *AwsKinesisanalyticsv2Application_ReferenceSchemaProperty `field:"required" json:"referenceSchema" yaml:"referenceSchema"`
	// s3_reference_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#s3_reference_data_source AwsKinesisanalyticsv2Application#s3_reference_data_source}
	// Experimental.
	S3ReferenceDataSource *AwsKinesisanalyticsv2Application_S3ReferenceDataSourceProperty `field:"required" json:"s3ReferenceDataSource" yaml:"s3ReferenceDataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#table_name AwsKinesisanalyticsv2Application#table_name}.
	// Experimental.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

