package bedrock


// Experimental.
type AwsEvaluationJob_RetrieveAndGenerateConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#knowledge_base_id AwsEvaluationJob#knowledge_base_id}.
	// Experimental.
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#model_arn AwsEvaluationJob#model_arn}.
	// Experimental.
	ModelArn *string `field:"required" json:"modelArn" yaml:"modelArn"`
	// retrieval_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#retrieval_configuration AwsEvaluationJob#retrieval_configuration}
	// Experimental.
	RetrievalConfiguration interface{} `field:"optional" json:"retrievalConfiguration" yaml:"retrievalConfiguration"`
}

