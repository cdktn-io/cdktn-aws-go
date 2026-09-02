package awsbedrock


// Experimental.
type TfEvaluationJob_RetrieveConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#knowledge_base_id TfEvaluationJob#knowledge_base_id}.
	// Experimental.
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// knowledge_base_retrieval_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#knowledge_base_retrieval_configuration TfEvaluationJob#knowledge_base_retrieval_configuration}
	// Experimental.
	KnowledgeBaseRetrievalConfiguration interface{} `field:"optional" json:"knowledgeBaseRetrievalConfiguration" yaml:"knowledgeBaseRetrievalConfiguration"`
}

