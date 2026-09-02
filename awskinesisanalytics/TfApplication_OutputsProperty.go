package awskinesisanalytics


// Experimental.
type TfApplication_OutputsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#name TfApplication#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#schema TfApplication#schema}
	// Experimental.
	Schema *TfApplication_OutputsSchemaProperty `field:"required" json:"schema" yaml:"schema"`
	// kinesis_firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#kinesis_firehose TfApplication#kinesis_firehose}
	// Experimental.
	KinesisFirehose *TfApplication_OutputsKinesisFirehoseProperty `field:"optional" json:"kinesisFirehose" yaml:"kinesisFirehose"`
	// kinesis_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#kinesis_stream TfApplication#kinesis_stream}
	// Experimental.
	KinesisStream *TfApplication_OutputsKinesisStreamProperty `field:"optional" json:"kinesisStream" yaml:"kinesisStream"`
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#lambda TfApplication#lambda}
	// Experimental.
	Lambda *TfApplication_OutputsLambdaProperty `field:"optional" json:"lambda" yaml:"lambda"`
}

