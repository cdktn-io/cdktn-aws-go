package awscloudwatchevidently


// Experimental.
type TfProject_DataDeliveryProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_project#cloudwatch_logs TfProject#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *TfProject_CloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// s3_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_project#s3_destination TfProject#s3_destination}
	// Experimental.
	S3Destination *TfProject_S3DestinationProperty `field:"optional" json:"s3Destination" yaml:"s3Destination"`
}

