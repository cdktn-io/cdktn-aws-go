package awsappconfig


// Experimental.
type TfExtension_ActionPointProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appconfig_extension#action TfExtension#action}
	// Experimental.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appconfig_extension#point TfExtension#point}.
	// Experimental.
	Point *string `field:"required" json:"point" yaml:"point"`
}

