package awsautoscalingplans


// Experimental.
type AwsAutoscalingplansScalingPlan_ScalingInstructionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#max_capacity AwsAutoscalingplansScalingPlan#max_capacity}.
	// Experimental.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#min_capacity AwsAutoscalingplansScalingPlan#min_capacity}.
	// Experimental.
	MinCapacity *float64 `field:"required" json:"minCapacity" yaml:"minCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#resource_id AwsAutoscalingplansScalingPlan#resource_id}.
	// Experimental.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#scalable_dimension AwsAutoscalingplansScalingPlan#scalable_dimension}.
	// Experimental.
	ScalableDimension *string `field:"required" json:"scalableDimension" yaml:"scalableDimension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#service_namespace AwsAutoscalingplansScalingPlan#service_namespace}.
	// Experimental.
	ServiceNamespace *string `field:"required" json:"serviceNamespace" yaml:"serviceNamespace"`
	// target_tracking_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#target_tracking_configuration AwsAutoscalingplansScalingPlan#target_tracking_configuration}
	// Experimental.
	TargetTrackingConfiguration interface{} `field:"required" json:"targetTrackingConfiguration" yaml:"targetTrackingConfiguration"`
	// customized_load_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#customized_load_metric_specification AwsAutoscalingplansScalingPlan#customized_load_metric_specification}
	// Experimental.
	CustomizedLoadMetricSpecification *AwsAutoscalingplansScalingPlan_CustomizedLoadMetricSpecificationProperty `field:"optional" json:"customizedLoadMetricSpecification" yaml:"customizedLoadMetricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#disable_dynamic_scaling AwsAutoscalingplansScalingPlan#disable_dynamic_scaling}.
	// Experimental.
	DisableDynamicScaling interface{} `field:"optional" json:"disableDynamicScaling" yaml:"disableDynamicScaling"`
	// predefined_load_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#predefined_load_metric_specification AwsAutoscalingplansScalingPlan#predefined_load_metric_specification}
	// Experimental.
	PredefinedLoadMetricSpecification *AwsAutoscalingplansScalingPlan_PredefinedLoadMetricSpecificationProperty `field:"optional" json:"predefinedLoadMetricSpecification" yaml:"predefinedLoadMetricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#predictive_scaling_max_capacity_behavior AwsAutoscalingplansScalingPlan#predictive_scaling_max_capacity_behavior}.
	// Experimental.
	PredictiveScalingMaxCapacityBehavior *string `field:"optional" json:"predictiveScalingMaxCapacityBehavior" yaml:"predictiveScalingMaxCapacityBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#predictive_scaling_max_capacity_buffer AwsAutoscalingplansScalingPlan#predictive_scaling_max_capacity_buffer}.
	// Experimental.
	PredictiveScalingMaxCapacityBuffer *float64 `field:"optional" json:"predictiveScalingMaxCapacityBuffer" yaml:"predictiveScalingMaxCapacityBuffer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#predictive_scaling_mode AwsAutoscalingplansScalingPlan#predictive_scaling_mode}.
	// Experimental.
	PredictiveScalingMode *string `field:"optional" json:"predictiveScalingMode" yaml:"predictiveScalingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#scaling_policy_update_behavior AwsAutoscalingplansScalingPlan#scaling_policy_update_behavior}.
	// Experimental.
	ScalingPolicyUpdateBehavior *string `field:"optional" json:"scalingPolicyUpdateBehavior" yaml:"scalingPolicyUpdateBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#scheduled_action_buffer_time AwsAutoscalingplansScalingPlan#scheduled_action_buffer_time}.
	// Experimental.
	ScheduledActionBufferTime *float64 `field:"optional" json:"scheduledActionBufferTime" yaml:"scheduledActionBufferTime"`
}

