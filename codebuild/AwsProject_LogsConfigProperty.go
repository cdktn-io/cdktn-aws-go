package codebuild


// Experimental.
type AwsProject_LogsConfigProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#cloudwatch_logs AwsProject#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *AwsProject_CloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// s3_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#s3_logs AwsProject#s3_logs}
	// Experimental.
	S3Logs *AwsProject_S3LogsProperty `field:"optional" json:"s3Logs" yaml:"s3Logs"`
}

