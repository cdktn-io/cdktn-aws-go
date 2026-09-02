package awsbedrock


// Experimental.
type TfEvaluationJob_InferenceConfigProperty struct {
	// model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#model TfEvaluationJob#model}
	// Experimental.
	Model interface{} `field:"optional" json:"model" yaml:"model"`
	// rag_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#rag_config TfEvaluationJob#rag_config}
	// Experimental.
	RagConfig interface{} `field:"optional" json:"ragConfig" yaml:"ragConfig"`
}

