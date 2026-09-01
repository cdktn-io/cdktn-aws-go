package awssecurityhub


// Experimental.
type AwsSecurityhubConfigurationPolicy_SecurityControlsConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#disabled_control_identifiers AwsSecurityhubConfigurationPolicy#disabled_control_identifiers}.
	// Experimental.
	DisabledControlIdentifiers *[]*string `field:"optional" json:"disabledControlIdentifiers" yaml:"disabledControlIdentifiers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#enabled_control_identifiers AwsSecurityhubConfigurationPolicy#enabled_control_identifiers}.
	// Experimental.
	EnabledControlIdentifiers *[]*string `field:"optional" json:"enabledControlIdentifiers" yaml:"enabledControlIdentifiers"`
	// security_control_custom_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#security_control_custom_parameter AwsSecurityhubConfigurationPolicy#security_control_custom_parameter}
	// Experimental.
	SecurityControlCustomParameter interface{} `field:"optional" json:"securityControlCustomParameter" yaml:"securityControlCustomParameter"`
}

