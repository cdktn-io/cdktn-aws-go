package bedrock


// Experimental.
type AwsEvaluationJob_CustomMetricDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#instructions AwsEvaluationJob#instructions}.
	// Experimental.
	Instructions *string `field:"required" json:"instructions" yaml:"instructions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#name AwsEvaluationJob#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// rating_scale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#rating_scale AwsEvaluationJob#rating_scale}
	// Experimental.
	RatingScale interface{} `field:"optional" json:"ratingScale" yaml:"ratingScale"`
}

