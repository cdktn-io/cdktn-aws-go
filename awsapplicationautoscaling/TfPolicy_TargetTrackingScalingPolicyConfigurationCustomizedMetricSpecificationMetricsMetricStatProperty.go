package awsapplicationautoscaling


// Experimental.
type TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatProperty struct {
	// metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#metric TfPolicy#metric}
	// Experimental.
	Metric *TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricProperty `field:"required" json:"metric" yaml:"metric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#stat TfPolicy#stat}.
	// Experimental.
	Stat *string `field:"required" json:"stat" yaml:"stat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#unit TfPolicy#unit}.
	// Experimental.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

