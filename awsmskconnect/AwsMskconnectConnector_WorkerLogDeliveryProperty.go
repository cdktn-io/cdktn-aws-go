package awsmskconnect


// Experimental.
type AwsMskconnectConnector_WorkerLogDeliveryProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#cloudwatch_logs AwsMskconnectConnector#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *AwsMskconnectConnector_CloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#firehose AwsMskconnectConnector#firehose}
	// Experimental.
	Firehose *AwsMskconnectConnector_FirehoseProperty `field:"optional" json:"firehose" yaml:"firehose"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#s3 AwsMskconnectConnector#s3}
	// Experimental.
	S3 *AwsMskconnectConnector_S3Property `field:"optional" json:"s3" yaml:"s3"`
}

