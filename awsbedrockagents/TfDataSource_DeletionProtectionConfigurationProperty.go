package awsbedrockagents


// Experimental.
type TfDataSource_DeletionProtectionConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#deletion_protection_status TfDataSource#deletion_protection_status}.
	// Experimental.
	DeletionProtectionStatus *string `field:"required" json:"deletionProtectionStatus" yaml:"deletionProtectionStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#deletion_protection_threshold TfDataSource#deletion_protection_threshold}.
	// Experimental.
	DeletionProtectionThreshold *float64 `field:"optional" json:"deletionProtectionThreshold" yaml:"deletionProtectionThreshold"`
}

