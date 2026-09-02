package awswaf


// Experimental.
type TfWebAclRule_RateBasedStatementProperty struct {
	// Setting that indicates how to aggregate the request counts. Valid values: IP, FORWARDED_IP, CUSTOM_KEYS, CONSTANT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#aggregate_key_type TfWebAclRule#aggregate_key_type}
	// Experimental.
	AggregateKeyType *string `field:"required" json:"aggregateKeyType" yaml:"aggregateKeyType"`
	// Rate limit threshold (10-2000000000).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#limit TfWebAclRule#limit}
	// Experimental.
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// custom_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#custom_keys TfWebAclRule#custom_keys}
	// Experimental.
	CustomKeys interface{} `field:"optional" json:"customKeys" yaml:"customKeys"`
	// Time window for AWS WAF to use to check the rate (60, 120, 300, 600).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#evaluation_window_sec TfWebAclRule#evaluation_window_sec}
	// Experimental.
	EvaluationWindowSec *float64 `field:"optional" json:"evaluationWindowSec" yaml:"evaluationWindowSec"`
	// forwarded_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#forwarded_ip_config TfWebAclRule#forwarded_ip_config}
	// Experimental.
	ForwardedIpConfig interface{} `field:"optional" json:"forwardedIpConfig" yaml:"forwardedIpConfig"`
	// scope_down_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#scope_down_statement TfWebAclRule#scope_down_statement}
	// Experimental.
	ScopeDownStatement interface{} `field:"optional" json:"scopeDownStatement" yaml:"scopeDownStatement"`
}

