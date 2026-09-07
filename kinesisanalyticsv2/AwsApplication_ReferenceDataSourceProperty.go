package kinesisanalyticsv2


// Experimental.
type AwsApplication_ReferenceDataSourceProperty struct {
	// reference_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#reference_schema AwsApplication#reference_schema}
	// Experimental.
	ReferenceSchema *AwsApplication_ReferenceSchemaProperty `field:"required" json:"referenceSchema" yaml:"referenceSchema"`
	// s3_reference_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#s3_reference_data_source AwsApplication#s3_reference_data_source}
	// Experimental.
	S3ReferenceDataSource *AwsApplication_S3ReferenceDataSourceProperty `field:"required" json:"s3ReferenceDataSource" yaml:"s3ReferenceDataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#table_name AwsApplication#table_name}.
	// Experimental.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

