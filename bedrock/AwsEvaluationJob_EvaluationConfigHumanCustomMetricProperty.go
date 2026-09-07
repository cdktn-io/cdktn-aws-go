package bedrock


// Experimental.
type AwsEvaluationJob_EvaluationConfigHumanCustomMetricProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#name AwsEvaluationJob#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#rating_method AwsEvaluationJob#rating_method}.
	// Experimental.
	RatingMethod *string `field:"required" json:"ratingMethod" yaml:"ratingMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#description AwsEvaluationJob#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

