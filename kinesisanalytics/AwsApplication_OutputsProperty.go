package kinesisanalytics


// Experimental.
type AwsApplication_OutputsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#name AwsApplication#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#schema AwsApplication#schema}
	// Experimental.
	Schema *AwsApplication_OutputsSchemaProperty `field:"required" json:"schema" yaml:"schema"`
	// kinesis_firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#kinesis_firehose AwsApplication#kinesis_firehose}
	// Experimental.
	KinesisFirehose *AwsApplication_OutputsKinesisFirehoseProperty `field:"optional" json:"kinesisFirehose" yaml:"kinesisFirehose"`
	// kinesis_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#kinesis_stream AwsApplication#kinesis_stream}
	// Experimental.
	KinesisStream *AwsApplication_OutputsKinesisStreamProperty `field:"optional" json:"kinesisStream" yaml:"kinesisStream"`
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#lambda AwsApplication#lambda}
	// Experimental.
	Lambda *AwsApplication_OutputsLambdaProperty `field:"optional" json:"lambda" yaml:"lambda"`
}

