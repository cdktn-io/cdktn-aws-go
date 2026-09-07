package computeoptimizer


// Experimental.
type AwsRecommendationPreferences_UtilizationPreferenceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#metric_name AwsRecommendationPreferences#metric_name}.
	// Experimental.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// metric_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#metric_parameters AwsRecommendationPreferences#metric_parameters}
	// Experimental.
	MetricParameters interface{} `field:"optional" json:"metricParameters" yaml:"metricParameters"`
}

