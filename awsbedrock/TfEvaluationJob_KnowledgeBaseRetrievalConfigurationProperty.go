package awsbedrock


// Experimental.
type TfEvaluationJob_KnowledgeBaseRetrievalConfigurationProperty struct {
	// vector_search_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#vector_search_configuration TfEvaluationJob#vector_search_configuration}
	// Experimental.
	VectorSearchConfiguration interface{} `field:"optional" json:"vectorSearchConfiguration" yaml:"vectorSearchConfiguration"`
}

