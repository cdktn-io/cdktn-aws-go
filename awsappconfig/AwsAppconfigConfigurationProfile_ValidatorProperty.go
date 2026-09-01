package awsappconfig


// Experimental.
type AwsAppconfigConfigurationProfile_ValidatorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appconfig_configuration_profile#type AwsAppconfigConfigurationProfile#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appconfig_configuration_profile#content AwsAppconfigConfigurationProfile#content}.
	// Experimental.
	Content *string `field:"optional" json:"content" yaml:"content"`
}

