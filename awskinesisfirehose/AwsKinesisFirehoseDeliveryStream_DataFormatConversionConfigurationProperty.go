package awskinesisfirehose


// Experimental.
type AwsKinesisFirehoseDeliveryStream_DataFormatConversionConfigurationProperty struct {
	// input_format_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#input_format_configuration AwsKinesisFirehoseDeliveryStream#input_format_configuration}
	// Experimental.
	InputFormatConfiguration *AwsKinesisFirehoseDeliveryStream_InputFormatConfigurationProperty `field:"required" json:"inputFormatConfiguration" yaml:"inputFormatConfiguration"`
	// output_format_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#output_format_configuration AwsKinesisFirehoseDeliveryStream#output_format_configuration}
	// Experimental.
	OutputFormatConfiguration *AwsKinesisFirehoseDeliveryStream_OutputFormatConfigurationProperty `field:"required" json:"outputFormatConfiguration" yaml:"outputFormatConfiguration"`
	// schema_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#schema_configuration AwsKinesisFirehoseDeliveryStream#schema_configuration}
	// Experimental.
	SchemaConfiguration *AwsKinesisFirehoseDeliveryStream_SchemaConfigurationProperty `field:"required" json:"schemaConfiguration" yaml:"schemaConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#enabled AwsKinesisFirehoseDeliveryStream#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

