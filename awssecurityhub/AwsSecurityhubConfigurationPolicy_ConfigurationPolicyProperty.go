package awssecurityhub


// Experimental.
type AwsSecurityhubConfigurationPolicy_ConfigurationPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#service_enabled AwsSecurityhubConfigurationPolicy#service_enabled}.
	// Experimental.
	ServiceEnabled interface{} `field:"required" json:"serviceEnabled" yaml:"serviceEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#enabled_standard_arns AwsSecurityhubConfigurationPolicy#enabled_standard_arns}.
	// Experimental.
	EnabledStandardArns *[]*string `field:"optional" json:"enabledStandardArns" yaml:"enabledStandardArns"`
	// security_controls_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_configuration_policy#security_controls_configuration AwsSecurityhubConfigurationPolicy#security_controls_configuration}
	// Experimental.
	SecurityControlsConfiguration *AwsSecurityhubConfigurationPolicy_SecurityControlsConfigurationProperty `field:"optional" json:"securityControlsConfiguration" yaml:"securityControlsConfiguration"`
}

