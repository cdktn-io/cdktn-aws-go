package awsautoscaling


// Experimental.
type TfPolicy_PredictiveScalingConfigurationProperty struct {
	// metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#metric_specification TfPolicy#metric_specification}
	// Experimental.
	MetricSpecification *TfPolicy_MetricSpecificationProperty `field:"required" json:"metricSpecification" yaml:"metricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#max_capacity_breach_behavior TfPolicy#max_capacity_breach_behavior}.
	// Experimental.
	MaxCapacityBreachBehavior *string `field:"optional" json:"maxCapacityBreachBehavior" yaml:"maxCapacityBreachBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#max_capacity_buffer TfPolicy#max_capacity_buffer}.
	// Experimental.
	MaxCapacityBuffer *string `field:"optional" json:"maxCapacityBuffer" yaml:"maxCapacityBuffer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#mode TfPolicy#mode}.
	// Experimental.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#scheduling_buffer_time TfPolicy#scheduling_buffer_time}.
	// Experimental.
	SchedulingBufferTime *string `field:"optional" json:"schedulingBufferTime" yaml:"schedulingBufferTime"`
}

