package awscomputeoptimizer

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsComputeoptimizerRecommendationPreferencesConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#resource_type AwsComputeoptimizerRecommendationPreferences#resource_type}.
	// Experimental.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#enhanced_infrastructure_metrics AwsComputeoptimizerRecommendationPreferences#enhanced_infrastructure_metrics}.
	// Experimental.
	EnhancedInfrastructureMetrics *string `field:"optional" json:"enhancedInfrastructureMetrics" yaml:"enhancedInfrastructureMetrics"`
	// external_metrics_preference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#external_metrics_preference AwsComputeoptimizerRecommendationPreferences#external_metrics_preference}
	// Experimental.
	ExternalMetricsPreference interface{} `field:"optional" json:"externalMetricsPreference" yaml:"externalMetricsPreference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#inferred_workload_types AwsComputeoptimizerRecommendationPreferences#inferred_workload_types}.
	// Experimental.
	InferredWorkloadTypes *string `field:"optional" json:"inferredWorkloadTypes" yaml:"inferredWorkloadTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#look_back_period AwsComputeoptimizerRecommendationPreferences#look_back_period}.
	// Experimental.
	LookBackPeriod *string `field:"optional" json:"lookBackPeriod" yaml:"lookBackPeriod"`
	// preferred_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#preferred_resource AwsComputeoptimizerRecommendationPreferences#preferred_resource}
	// Experimental.
	PreferredResource interface{} `field:"optional" json:"preferredResource" yaml:"preferredResource"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#region AwsComputeoptimizerRecommendationPreferences#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#savings_estimation_mode AwsComputeoptimizerRecommendationPreferences#savings_estimation_mode}.
	// Experimental.
	SavingsEstimationMode *string `field:"optional" json:"savingsEstimationMode" yaml:"savingsEstimationMode"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#scope AwsComputeoptimizerRecommendationPreferences#scope}
	// Experimental.
	Scope interface{} `field:"optional" json:"scope" yaml:"scope"`
	// utilization_preference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences#utilization_preference AwsComputeoptimizerRecommendationPreferences#utilization_preference}
	// Experimental.
	UtilizationPreference interface{} `field:"optional" json:"utilizationPreference" yaml:"utilizationPreference"`
}

