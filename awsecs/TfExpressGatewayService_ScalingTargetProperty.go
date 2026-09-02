package awsecs


// Experimental.
type TfExpressGatewayService_ScalingTargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_express_gateway_service#auto_scaling_metric TfExpressGatewayService#auto_scaling_metric}.
	// Experimental.
	AutoScalingMetric *string `field:"optional" json:"autoScalingMetric" yaml:"autoScalingMetric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_express_gateway_service#auto_scaling_target_value TfExpressGatewayService#auto_scaling_target_value}.
	// Experimental.
	AutoScalingTargetValue *float64 `field:"optional" json:"autoScalingTargetValue" yaml:"autoScalingTargetValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_express_gateway_service#max_task_count TfExpressGatewayService#max_task_count}.
	// Experimental.
	MaxTaskCount *float64 `field:"optional" json:"maxTaskCount" yaml:"maxTaskCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_express_gateway_service#min_task_count TfExpressGatewayService#min_task_count}.
	// Experimental.
	MinTaskCount *float64 `field:"optional" json:"minTaskCount" yaml:"minTaskCount"`
}

