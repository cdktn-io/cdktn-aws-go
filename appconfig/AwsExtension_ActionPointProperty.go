package appconfig


// Experimental.
type AwsExtension_ActionPointProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appconfig_extension#action AwsExtension#action}
	// Experimental.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appconfig_extension#point AwsExtension#point}.
	// Experimental.
	Point *string `field:"required" json:"point" yaml:"point"`
}

