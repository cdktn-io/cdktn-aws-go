package awsbedrock


// Experimental.
type TfEvaluationJob_RatingScaleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#definition TfEvaluationJob#definition}.
	// Experimental.
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#value TfEvaluationJob#value}
	// Experimental.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

