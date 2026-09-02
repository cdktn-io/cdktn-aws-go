package awskinesisfirehose


// Experimental.
type TfDeliveryStream_DestinationTableConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#database_name TfDeliveryStream#database_name}.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#table_name TfDeliveryStream#table_name}.
	// Experimental.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#s3_error_output_prefix TfDeliveryStream#s3_error_output_prefix}.
	// Experimental.
	S3ErrorOutputPrefix *string `field:"optional" json:"s3ErrorOutputPrefix" yaml:"s3ErrorOutputPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#unique_keys TfDeliveryStream#unique_keys}.
	// Experimental.
	UniqueKeys *[]*string `field:"optional" json:"uniqueKeys" yaml:"uniqueKeys"`
}

