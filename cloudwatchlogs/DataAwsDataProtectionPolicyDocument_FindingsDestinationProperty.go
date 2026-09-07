package cloudwatchlogs


// Experimental.
type DataAwsDataProtectionPolicyDocument_FindingsDestinationProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#cloudwatch_logs DataAwsDataProtectionPolicyDocument#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *DataAwsDataProtectionPolicyDocument_CloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#firehose DataAwsDataProtectionPolicyDocument#firehose}
	// Experimental.
	Firehose *DataAwsDataProtectionPolicyDocument_FirehoseProperty `field:"optional" json:"firehose" yaml:"firehose"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#s3 DataAwsDataProtectionPolicyDocument#s3}
	// Experimental.
	S3 *DataAwsDataProtectionPolicyDocument_S3Property `field:"optional" json:"s3" yaml:"s3"`
}

