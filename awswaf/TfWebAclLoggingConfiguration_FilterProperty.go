package awswaf


// Experimental.
type TfWebAclLoggingConfiguration_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#behavior TfWebAclLoggingConfiguration#behavior}.
	// Experimental.
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#condition TfWebAclLoggingConfiguration#condition}
	// Experimental.
	Condition interface{} `field:"required" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#requirement TfWebAclLoggingConfiguration#requirement}.
	// Experimental.
	Requirement *string `field:"required" json:"requirement" yaml:"requirement"`
}

