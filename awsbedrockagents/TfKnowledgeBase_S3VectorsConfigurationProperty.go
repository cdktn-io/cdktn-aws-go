package awsbedrockagents


// Experimental.
type TfKnowledgeBase_S3VectorsConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#index_arn TfKnowledgeBase#index_arn}.
	// Experimental.
	IndexArn *string `field:"optional" json:"indexArn" yaml:"indexArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#index_name TfKnowledgeBase#index_name}.
	// Experimental.
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#vector_bucket_arn TfKnowledgeBase#vector_bucket_arn}.
	// Experimental.
	VectorBucketArn *string `field:"optional" json:"vectorBucketArn" yaml:"vectorBucketArn"`
}

