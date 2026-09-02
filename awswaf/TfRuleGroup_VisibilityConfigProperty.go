package awswaf


// Experimental.
type TfRuleGroup_VisibilityConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#cloudwatch_metrics_enabled TfRuleGroup#cloudwatch_metrics_enabled}.
	// Experimental.
	CloudwatchMetricsEnabled interface{} `field:"required" json:"cloudwatchMetricsEnabled" yaml:"cloudwatchMetricsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#metric_name TfRuleGroup#metric_name}.
	// Experimental.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#sampled_requests_enabled TfRuleGroup#sampled_requests_enabled}.
	// Experimental.
	SampledRequestsEnabled interface{} `field:"required" json:"sampledRequestsEnabled" yaml:"sampledRequestsEnabled"`
}

