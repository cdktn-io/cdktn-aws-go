package fms


// Experimental.
type AwsPolicy_SecurityServicePolicyDataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#type AwsPolicy#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#managed_service_data AwsPolicy#managed_service_data}.
	// Experimental.
	ManagedServiceData *string `field:"optional" json:"managedServiceData" yaml:"managedServiceData"`
	// policy_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#policy_option AwsPolicy#policy_option}
	// Experimental.
	PolicyOption *AwsPolicy_PolicyOptionProperty `field:"optional" json:"policyOption" yaml:"policyOption"`
}

