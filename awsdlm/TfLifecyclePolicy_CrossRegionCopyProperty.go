package awsdlm


// Experimental.
type TfLifecyclePolicy_CrossRegionCopyProperty struct {
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#encryption_configuration TfLifecyclePolicy#encryption_configuration}
	// Experimental.
	EncryptionConfiguration *TfLifecyclePolicy_EncryptionConfigurationProperty `field:"required" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#target TfLifecyclePolicy#target}.
	// Experimental.
	Target *string `field:"required" json:"target" yaml:"target"`
	// retain_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#retain_rule TfLifecyclePolicy#retain_rule}
	// Experimental.
	RetainRule *TfLifecyclePolicy_PolicyDetailsActionCrossRegionCopyRetainRuleProperty `field:"optional" json:"retainRule" yaml:"retainRule"`
}

