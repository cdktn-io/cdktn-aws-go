package awsbedrock


// Experimental.
type TfEvaluationJob_KnowledgeBaseConfigProperty struct {
	// retrieve_and_generate_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#retrieve_and_generate_config TfEvaluationJob#retrieve_and_generate_config}
	// Experimental.
	RetrieveAndGenerateConfig interface{} `field:"optional" json:"retrieveAndGenerateConfig" yaml:"retrieveAndGenerateConfig"`
	// retrieve_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#retrieve_config TfEvaluationJob#retrieve_config}
	// Experimental.
	RetrieveConfig interface{} `field:"optional" json:"retrieveConfig" yaml:"retrieveConfig"`
}

