package awssecurityhub


// Experimental.
type TfConfigurationPolicy_SecurityControlCustomParameterProperty struct {
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#parameter TfConfigurationPolicy#parameter}
	// Experimental.
	Parameter interface{} `field:"required" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#security_control_id TfConfigurationPolicy#security_control_id}.
	// Experimental.
	SecurityControlId *string `field:"required" json:"securityControlId" yaml:"securityControlId"`
}

