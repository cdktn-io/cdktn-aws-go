package awscloudwatchlogs


// Experimental.
type DataAwsCloudwatchLogDataProtectionPolicyDocument_StatementProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#data_identifiers DataAwsCloudwatchLogDataProtectionPolicyDocument#data_identifiers}.
	// Experimental.
	DataIdentifiers *[]*string `field:"required" json:"dataIdentifiers" yaml:"dataIdentifiers"`
	// operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#operation DataAwsCloudwatchLogDataProtectionPolicyDocument#operation}
	// Experimental.
	Operation *DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationProperty `field:"required" json:"operation" yaml:"operation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#sid DataAwsCloudwatchLogDataProtectionPolicyDocument#sid}.
	// Experimental.
	Sid *string `field:"optional" json:"sid" yaml:"sid"`
}

