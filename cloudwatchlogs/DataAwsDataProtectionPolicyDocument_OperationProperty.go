package cloudwatchlogs


// Experimental.
type DataAwsDataProtectionPolicyDocument_OperationProperty struct {
	// audit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#audit DataAwsDataProtectionPolicyDocument#audit}
	// Experimental.
	Audit *DataAwsDataProtectionPolicyDocument_AuditProperty `field:"optional" json:"audit" yaml:"audit"`
	// deidentify block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#deidentify DataAwsDataProtectionPolicyDocument#deidentify}
	// Experimental.
	Deidentify *DataAwsDataProtectionPolicyDocument_DeidentifyProperty `field:"optional" json:"deidentify" yaml:"deidentify"`
}

