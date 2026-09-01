package awsdlm


// Experimental.
type AwsDlmLifecyclePolicy_CrossRegionCopyProperty struct {
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#encryption_configuration AwsDlmLifecyclePolicy#encryption_configuration}
	// Experimental.
	EncryptionConfiguration *AwsDlmLifecyclePolicy_EncryptionConfigurationProperty `field:"required" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#target AwsDlmLifecyclePolicy#target}.
	// Experimental.
	Target *string `field:"required" json:"target" yaml:"target"`
	// retain_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#retain_rule AwsDlmLifecyclePolicy#retain_rule}
	// Experimental.
	RetainRule *AwsDlmLifecyclePolicy_PolicyDetailsActionCrossRegionCopyRetainRuleProperty `field:"optional" json:"retainRule" yaml:"retainRule"`
}

