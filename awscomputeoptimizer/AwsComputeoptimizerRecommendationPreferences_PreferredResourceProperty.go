package awscomputeoptimizer


// Experimental.
type AwsComputeoptimizerRecommendationPreferences_PreferredResourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#name AwsComputeoptimizerRecommendationPreferences#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#exclude_list AwsComputeoptimizerRecommendationPreferences#exclude_list}.
	// Experimental.
	ExcludeList *[]*string `field:"optional" json:"excludeList" yaml:"excludeList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#include_list AwsComputeoptimizerRecommendationPreferences#include_list}.
	// Experimental.
	IncludeList *[]*string `field:"optional" json:"includeList" yaml:"includeList"`
}

