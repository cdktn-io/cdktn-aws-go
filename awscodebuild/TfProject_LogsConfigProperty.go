package awscodebuild


// Experimental.
type TfProject_LogsConfigProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#cloudwatch_logs TfProject#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *TfProject_CloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// s3_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#s3_logs TfProject#s3_logs}
	// Experimental.
	S3Logs *TfProject_S3LogsProperty `field:"optional" json:"s3Logs" yaml:"s3Logs"`
}

