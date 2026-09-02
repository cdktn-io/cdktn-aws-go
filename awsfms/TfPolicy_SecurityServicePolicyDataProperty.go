package awsfms


// Experimental.
type TfPolicy_SecurityServicePolicyDataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#type TfPolicy#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#managed_service_data TfPolicy#managed_service_data}.
	// Experimental.
	ManagedServiceData *string `field:"optional" json:"managedServiceData" yaml:"managedServiceData"`
	// policy_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#policy_option TfPolicy#policy_option}
	// Experimental.
	PolicyOption *TfPolicy_PolicyOptionProperty `field:"optional" json:"policyOption" yaml:"policyOption"`
}

