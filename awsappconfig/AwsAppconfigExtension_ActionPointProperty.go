package awsappconfig


// Experimental.
type AwsAppconfigExtension_ActionPointProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appconfig_extension#action AwsAppconfigExtension#action}
	// Experimental.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appconfig_extension#point AwsAppconfigExtension#point}.
	// Experimental.
	Point *string `field:"required" json:"point" yaml:"point"`
}

