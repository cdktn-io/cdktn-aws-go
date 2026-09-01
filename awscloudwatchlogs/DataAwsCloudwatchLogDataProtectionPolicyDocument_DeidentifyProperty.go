package awscloudwatchlogs


// Experimental.
type DataAwsCloudwatchLogDataProtectionPolicyDocument_DeidentifyProperty struct {
	// mask_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#mask_config DataAwsCloudwatchLogDataProtectionPolicyDocument#mask_config}
	// Experimental.
	MaskConfig *DataAwsCloudwatchLogDataProtectionPolicyDocument_MaskConfigProperty `field:"required" json:"maskConfig" yaml:"maskConfig"`
}

