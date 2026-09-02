package awswaf


// Experimental.
type TfWebAclRule_VisibilityConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#cloudwatch_metrics_enabled TfWebAclRule#cloudwatch_metrics_enabled}.
	// Experimental.
	CloudwatchMetricsEnabled interface{} `field:"required" json:"cloudwatchMetricsEnabled" yaml:"cloudwatchMetricsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#metric_name TfWebAclRule#metric_name}.
	// Experimental.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#sampled_requests_enabled TfWebAclRule#sampled_requests_enabled}.
	// Experimental.
	SampledRequestsEnabled interface{} `field:"required" json:"sampledRequestsEnabled" yaml:"sampledRequestsEnabled"`
}

