package awscomputeoptimizer


// Experimental.
type AwsComputeoptimizerRecommendationPreferences_MetricParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#headroom AwsComputeoptimizerRecommendationPreferences#headroom}.
	// Experimental.
	Headroom *string `field:"required" json:"headroom" yaml:"headroom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#threshold AwsComputeoptimizerRecommendationPreferences#threshold}.
	// Experimental.
	Threshold *string `field:"optional" json:"threshold" yaml:"threshold"`
}

