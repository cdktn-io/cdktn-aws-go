package waf


// Experimental.
type AwsWebAclLoggingConfiguration_LoggingFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#default_behavior AwsWebAclLoggingConfiguration#default_behavior}.
	// Experimental.
	DefaultBehavior *string `field:"required" json:"defaultBehavior" yaml:"defaultBehavior"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#filter AwsWebAclLoggingConfiguration#filter}
	// Experimental.
	Filter interface{} `field:"required" json:"filter" yaml:"filter"`
}

