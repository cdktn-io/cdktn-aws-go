package awsapplicationautoscaling


// Experimental.
type TfPolicy_PredictiveScalingPolicyConfigurationMetricSpecificationCustomizedLoadMetricSpecificationMetricDataQueryMetricStatMetricDimensionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#name TfPolicy#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#value TfPolicy#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

