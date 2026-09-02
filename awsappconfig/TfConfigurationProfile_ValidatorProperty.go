package awsappconfig


// Experimental.
type TfConfigurationProfile_ValidatorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appconfig_configuration_profile#type TfConfigurationProfile#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appconfig_configuration_profile#content TfConfigurationProfile#content}.
	// Experimental.
	Content *string `field:"optional" json:"content" yaml:"content"`
}

