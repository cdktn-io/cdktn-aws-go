package awscloudwatchlogs


// Experimental.
type DataTfDataProtectionPolicyDocument_CustomDataIdentifierProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#name DataTfDataProtectionPolicyDocument#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#regex DataTfDataProtectionPolicyDocument#regex}.
	// Experimental.
	Regex *string `field:"required" json:"regex" yaml:"regex"`
}

