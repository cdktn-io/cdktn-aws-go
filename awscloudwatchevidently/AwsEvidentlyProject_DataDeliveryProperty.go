package awscloudwatchevidently


// Experimental.
type AwsEvidentlyProject_DataDeliveryProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_project#cloudwatch_logs AwsEvidentlyProject#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *AwsEvidentlyProject_CloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// s3_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_project#s3_destination AwsEvidentlyProject#s3_destination}
	// Experimental.
	S3Destination *AwsEvidentlyProject_S3DestinationProperty `field:"optional" json:"s3Destination" yaml:"s3Destination"`
}

