package awscomputeoptimizer


// Experimental.
type TfRecommendationPreferences_PreferredResourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#name TfRecommendationPreferences#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#exclude_list TfRecommendationPreferences#exclude_list}.
	// Experimental.
	ExcludeList *[]*string `field:"optional" json:"excludeList" yaml:"excludeList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#include_list TfRecommendationPreferences#include_list}.
	// Experimental.
	IncludeList *[]*string `field:"optional" json:"includeList" yaml:"includeList"`
}

