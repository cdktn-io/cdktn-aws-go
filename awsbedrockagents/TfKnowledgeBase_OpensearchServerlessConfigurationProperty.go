package awsbedrockagents


// Experimental.
type TfKnowledgeBase_OpensearchServerlessConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#collection_arn TfKnowledgeBase#collection_arn}.
	// Experimental.
	CollectionArn *string `field:"required" json:"collectionArn" yaml:"collectionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#vector_index_name TfKnowledgeBase#vector_index_name}.
	// Experimental.
	VectorIndexName *string `field:"required" json:"vectorIndexName" yaml:"vectorIndexName"`
	// field_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#field_mapping TfKnowledgeBase#field_mapping}
	// Experimental.
	FieldMapping interface{} `field:"optional" json:"fieldMapping" yaml:"fieldMapping"`
}

