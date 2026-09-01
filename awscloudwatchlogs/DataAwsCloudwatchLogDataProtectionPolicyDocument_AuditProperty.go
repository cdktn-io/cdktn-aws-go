package awscloudwatchlogs


// Experimental.
type DataAwsCloudwatchLogDataProtectionPolicyDocument_AuditProperty struct {
	// findings_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#findings_destination DataAwsCloudwatchLogDataProtectionPolicyDocument#findings_destination}
	// Experimental.
	FindingsDestination *DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationProperty `field:"required" json:"findingsDestination" yaml:"findingsDestination"`
}

