package awsfms


// Experimental.
type AwsFmsPolicy_SecurityServicePolicyDataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#type AwsFmsPolicy#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#managed_service_data AwsFmsPolicy#managed_service_data}.
	// Experimental.
	ManagedServiceData *string `field:"optional" json:"managedServiceData" yaml:"managedServiceData"`
	// policy_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#policy_option AwsFmsPolicy#policy_option}
	// Experimental.
	PolicyOption *AwsFmsPolicy_PolicyOptionProperty `field:"optional" json:"policyOption" yaml:"policyOption"`
}

