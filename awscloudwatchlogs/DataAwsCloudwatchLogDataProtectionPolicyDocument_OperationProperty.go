package awscloudwatchlogs


// Experimental.
type DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationProperty struct {
	// audit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#audit DataAwsCloudwatchLogDataProtectionPolicyDocument#audit}
	// Experimental.
	Audit *DataAwsCloudwatchLogDataProtectionPolicyDocument_AuditProperty `field:"optional" json:"audit" yaml:"audit"`
	// deidentify block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#deidentify DataAwsCloudwatchLogDataProtectionPolicyDocument#deidentify}
	// Experimental.
	Deidentify *DataAwsCloudwatchLogDataProtectionPolicyDocument_DeidentifyProperty `field:"optional" json:"deidentify" yaml:"deidentify"`
}

