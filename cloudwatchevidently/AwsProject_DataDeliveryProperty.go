package cloudwatchevidently


// Experimental.
type AwsProject_DataDeliveryProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_project#cloudwatch_logs AwsProject#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *AwsProject_CloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// s3_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_project#s3_destination AwsProject#s3_destination}
	// Experimental.
	S3Destination *AwsProject_S3DestinationProperty `field:"optional" json:"s3Destination" yaml:"s3Destination"`
}

