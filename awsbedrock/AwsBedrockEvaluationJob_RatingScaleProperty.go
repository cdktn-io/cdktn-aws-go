package awsbedrock


// Experimental.
type AwsBedrockEvaluationJob_RatingScaleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#definition AwsBedrockEvaluationJob#definition}.
	// Experimental.
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#value AwsBedrockEvaluationJob#value}
	// Experimental.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

