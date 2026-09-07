package appconfig


// Experimental.
type AwsConfigurationProfile_ValidatorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appconfig_configuration_profile#type AwsConfigurationProfile#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appconfig_configuration_profile#content AwsConfigurationProfile#content}.
	// Experimental.
	Content *string `field:"optional" json:"content" yaml:"content"`
}

