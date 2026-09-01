package awswaf


// Experimental.
type AwsWafv2WebAclLoggingConfiguration_LoggingFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#default_behavior AwsWafv2WebAclLoggingConfiguration#default_behavior}.
	// Experimental.
	DefaultBehavior *string `field:"required" json:"defaultBehavior" yaml:"defaultBehavior"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_logging_configuration#filter AwsWafv2WebAclLoggingConfiguration#filter}
	// Experimental.
	Filter interface{} `field:"required" json:"filter" yaml:"filter"`
}

