package awswaf


// Experimental.
type TfWebAclRuleGroupAssociation_VisibilityConfigProperty struct {
	// Indicates whether the rule is available for use in the metrics for the web ACL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#cloudwatch_metrics_enabled TfWebAclRuleGroupAssociation#cloudwatch_metrics_enabled}
	// Experimental.
	CloudwatchMetricsEnabled interface{} `field:"required" json:"cloudwatchMetricsEnabled" yaml:"cloudwatchMetricsEnabled"`
	// A name for the metrics for this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#metric_name TfWebAclRuleGroupAssociation#metric_name}
	// Experimental.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Indicates whether to store a sampling of the web requests that match the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#sampled_requests_enabled TfWebAclRuleGroupAssociation#sampled_requests_enabled}
	// Experimental.
	SampledRequestsEnabled interface{} `field:"required" json:"sampledRequestsEnabled" yaml:"sampledRequestsEnabled"`
}

