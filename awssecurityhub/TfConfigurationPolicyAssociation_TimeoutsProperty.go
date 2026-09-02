package awssecurityhub


// Experimental.
type TfConfigurationPolicyAssociation_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy_association#create TfConfigurationPolicyAssociation#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy_association#update TfConfigurationPolicyAssociation#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

