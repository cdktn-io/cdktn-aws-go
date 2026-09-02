package awskinesisfirehose


// Experimental.
type TfDeliveryStream_DataFormatConversionConfigurationProperty struct {
	// input_format_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#input_format_configuration TfDeliveryStream#input_format_configuration}
	// Experimental.
	InputFormatConfiguration *TfDeliveryStream_InputFormatConfigurationProperty `field:"required" json:"inputFormatConfiguration" yaml:"inputFormatConfiguration"`
	// output_format_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#output_format_configuration TfDeliveryStream#output_format_configuration}
	// Experimental.
	OutputFormatConfiguration *TfDeliveryStream_OutputFormatConfigurationProperty `field:"required" json:"outputFormatConfiguration" yaml:"outputFormatConfiguration"`
	// schema_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#schema_configuration TfDeliveryStream#schema_configuration}
	// Experimental.
	SchemaConfiguration *TfDeliveryStream_SchemaConfigurationProperty `field:"required" json:"schemaConfiguration" yaml:"schemaConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#enabled TfDeliveryStream#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

