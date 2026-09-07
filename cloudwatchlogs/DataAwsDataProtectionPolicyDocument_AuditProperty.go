package cloudwatchlogs


// Experimental.
type DataAwsDataProtectionPolicyDocument_AuditProperty struct {
	// findings_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#findings_destination DataAwsDataProtectionPolicyDocument#findings_destination}
	// Experimental.
	FindingsDestination *DataAwsDataProtectionPolicyDocument_FindingsDestinationProperty `field:"required" json:"findingsDestination" yaml:"findingsDestination"`
}

