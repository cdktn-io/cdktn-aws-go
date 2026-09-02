package awsbedrock


// Experimental.
type TfEvaluationJob_RetrieveAndGenerateConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#knowledge_base_id TfEvaluationJob#knowledge_base_id}.
	// Experimental.
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#model_arn TfEvaluationJob#model_arn}.
	// Experimental.
	ModelArn *string `field:"required" json:"modelArn" yaml:"modelArn"`
	// retrieval_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#retrieval_configuration TfEvaluationJob#retrieval_configuration}
	// Experimental.
	RetrievalConfiguration interface{} `field:"optional" json:"retrievalConfiguration" yaml:"retrievalConfiguration"`
}

