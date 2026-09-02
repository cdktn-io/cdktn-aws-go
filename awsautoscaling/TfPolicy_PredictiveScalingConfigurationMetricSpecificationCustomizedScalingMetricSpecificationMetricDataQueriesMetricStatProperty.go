package awsautoscaling


// Experimental.
type TfPolicy_PredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueriesMetricStatProperty struct {
	// metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#metric TfPolicy#metric}
	// Experimental.
	Metric *TfPolicy_PredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueriesMetricStatMetricProperty `field:"required" json:"metric" yaml:"metric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#stat TfPolicy#stat}.
	// Experimental.
	Stat *string `field:"required" json:"stat" yaml:"stat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#unit TfPolicy#unit}.
	// Experimental.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

