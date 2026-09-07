package bedrockagents


// Experimental.
type AwsKnowledgeBase_S3VectorsConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#index_arn AwsKnowledgeBase#index_arn}.
	// Experimental.
	IndexArn *string `field:"optional" json:"indexArn" yaml:"indexArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#index_name AwsKnowledgeBase#index_name}.
	// Experimental.
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#vector_bucket_arn AwsKnowledgeBase#vector_bucket_arn}.
	// Experimental.
	VectorBucketArn *string `field:"optional" json:"vectorBucketArn" yaml:"vectorBucketArn"`
}

