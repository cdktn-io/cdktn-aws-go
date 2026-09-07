package mskconnect


// Experimental.
type AwsConnector_WorkerLogDeliveryProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#cloudwatch_logs AwsConnector#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *AwsConnector_CloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#firehose AwsConnector#firehose}
	// Experimental.
	Firehose *AwsConnector_FirehoseProperty `field:"optional" json:"firehose" yaml:"firehose"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#s3 AwsConnector#s3}
	// Experimental.
	S3 *AwsConnector_S3Property `field:"optional" json:"s3" yaml:"s3"`
}

