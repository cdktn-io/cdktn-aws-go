package awsappconfig


// Experimental.
type TfExtension_ParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appconfig_extension#name TfExtension#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appconfig_extension#description TfExtension#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appconfig_extension#required TfExtension#required}.
	// Experimental.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}

