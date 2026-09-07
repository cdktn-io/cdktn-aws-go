package computeoptimizer


// Experimental.
type AwsRecommendationPreferences_MetricParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#headroom AwsRecommendationPreferences#headroom}.
	// Experimental.
	Headroom *string `field:"required" json:"headroom" yaml:"headroom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#threshold AwsRecommendationPreferences#threshold}.
	// Experimental.
	Threshold *string `field:"optional" json:"threshold" yaml:"threshold"`
}

