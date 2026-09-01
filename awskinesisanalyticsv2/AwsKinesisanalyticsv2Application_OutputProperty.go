package awskinesisanalyticsv2


// Experimental.
type AwsKinesisanalyticsv2Application_OutputProperty struct {
	// destination_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#destination_schema AwsKinesisanalyticsv2Application#destination_schema}
	// Experimental.
	DestinationSchema *AwsKinesisanalyticsv2Application_DestinationSchemaProperty `field:"required" json:"destinationSchema" yaml:"destinationSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#name AwsKinesisanalyticsv2Application#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// kinesis_firehose_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#kinesis_firehose_output AwsKinesisanalyticsv2Application#kinesis_firehose_output}
	// Experimental.
	KinesisFirehoseOutput *AwsKinesisanalyticsv2Application_KinesisFirehoseOutputProperty `field:"optional" json:"kinesisFirehoseOutput" yaml:"kinesisFirehoseOutput"`
	// kinesis_streams_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#kinesis_streams_output AwsKinesisanalyticsv2Application#kinesis_streams_output}
	// Experimental.
	KinesisStreamsOutput *AwsKinesisanalyticsv2Application_KinesisStreamsOutputProperty `field:"optional" json:"kinesisStreamsOutput" yaml:"kinesisStreamsOutput"`
	// lambda_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#lambda_output AwsKinesisanalyticsv2Application#lambda_output}
	// Experimental.
	LambdaOutput *AwsKinesisanalyticsv2Application_LambdaOutputProperty `field:"optional" json:"lambdaOutput" yaml:"lambdaOutput"`
}

