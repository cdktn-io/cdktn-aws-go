package awskinesisanalyticsv2


// Experimental.
type TfApplication_ReferenceDataSourceProperty struct {
	// reference_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#reference_schema TfApplication#reference_schema}
	// Experimental.
	ReferenceSchema *TfApplication_ReferenceSchemaProperty `field:"required" json:"referenceSchema" yaml:"referenceSchema"`
	// s3_reference_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#s3_reference_data_source TfApplication#s3_reference_data_source}
	// Experimental.
	S3ReferenceDataSource *TfApplication_S3ReferenceDataSourceProperty `field:"required" json:"s3ReferenceDataSource" yaml:"s3ReferenceDataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#table_name TfApplication#table_name}.
	// Experimental.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

