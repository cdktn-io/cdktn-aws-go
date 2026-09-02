package awsmskconnect


// Experimental.
type TfConnector_WorkerLogDeliveryProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#cloudwatch_logs TfConnector#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *TfConnector_CloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#firehose TfConnector#firehose}
	// Experimental.
	Firehose *TfConnector_FirehoseProperty `field:"optional" json:"firehose" yaml:"firehose"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#s3 TfConnector#s3}
	// Experimental.
	S3 *TfConnector_S3Property `field:"optional" json:"s3" yaml:"s3"`
}

