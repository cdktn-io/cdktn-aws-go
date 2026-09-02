package awscloudwatchlogs


// Experimental.
type DataTfDataProtectionPolicyDocument_AuditProperty struct {
	// findings_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#findings_destination DataTfDataProtectionPolicyDocument#findings_destination}
	// Experimental.
	FindingsDestination *DataTfDataProtectionPolicyDocument_FindingsDestinationProperty `field:"required" json:"findingsDestination" yaml:"findingsDestination"`
}

