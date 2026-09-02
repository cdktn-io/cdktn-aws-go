package awsbedrockagents


// Experimental.
type TfKnowledgeBase_MongoDbAtlasConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#collection_name TfKnowledgeBase#collection_name}.
	// Experimental.
	CollectionName *string `field:"required" json:"collectionName" yaml:"collectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#credentials_secret_arn TfKnowledgeBase#credentials_secret_arn}.
	// Experimental.
	CredentialsSecretArn *string `field:"required" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#database_name TfKnowledgeBase#database_name}.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#endpoint TfKnowledgeBase#endpoint}.
	// Experimental.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#vector_index_name TfKnowledgeBase#vector_index_name}.
	// Experimental.
	VectorIndexName *string `field:"required" json:"vectorIndexName" yaml:"vectorIndexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#endpoint_service_name TfKnowledgeBase#endpoint_service_name}.
	// Experimental.
	EndpointServiceName *string `field:"optional" json:"endpointServiceName" yaml:"endpointServiceName"`
	// field_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#field_mapping TfKnowledgeBase#field_mapping}
	// Experimental.
	FieldMapping interface{} `field:"optional" json:"fieldMapping" yaml:"fieldMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#text_index_name TfKnowledgeBase#text_index_name}.
	// Experimental.
	TextIndexName *string `field:"optional" json:"textIndexName" yaml:"textIndexName"`
}

