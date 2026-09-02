package awscloudwatchlogs


// Experimental.
type DataTfDataProtectionPolicyDocument_StatementProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#data_identifiers DataTfDataProtectionPolicyDocument#data_identifiers}.
	// Experimental.
	DataIdentifiers *[]*string `field:"required" json:"dataIdentifiers" yaml:"dataIdentifiers"`
	// operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#operation DataTfDataProtectionPolicyDocument#operation}
	// Experimental.
	Operation *DataTfDataProtectionPolicyDocument_OperationProperty `field:"required" json:"operation" yaml:"operation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#sid DataTfDataProtectionPolicyDocument#sid}.
	// Experimental.
	Sid *string `field:"optional" json:"sid" yaml:"sid"`
}

