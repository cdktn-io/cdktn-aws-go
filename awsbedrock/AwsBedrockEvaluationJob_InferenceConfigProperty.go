package awsbedrock


// Experimental.
type AwsBedrockEvaluationJob_InferenceConfigProperty struct {
	// model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#model AwsBedrockEvaluationJob#model}
	// Experimental.
	Model interface{} `field:"optional" json:"model" yaml:"model"`
	// rag_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#rag_config AwsBedrockEvaluationJob#rag_config}
	// Experimental.
	RagConfig interface{} `field:"optional" json:"ragConfig" yaml:"ragConfig"`
}

