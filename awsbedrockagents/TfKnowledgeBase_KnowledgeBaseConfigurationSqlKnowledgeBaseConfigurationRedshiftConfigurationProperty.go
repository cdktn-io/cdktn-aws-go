package awsbedrockagents


// Experimental.
type TfKnowledgeBase_KnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationProperty struct {
	// query_engine_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#query_engine_configuration TfKnowledgeBase#query_engine_configuration}
	// Experimental.
	QueryEngineConfiguration interface{} `field:"optional" json:"queryEngineConfiguration" yaml:"queryEngineConfiguration"`
	// query_generation_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#query_generation_configuration TfKnowledgeBase#query_generation_configuration}
	// Experimental.
	QueryGenerationConfiguration interface{} `field:"optional" json:"queryGenerationConfiguration" yaml:"queryGenerationConfiguration"`
	// storage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#storage_configuration TfKnowledgeBase#storage_configuration}
	// Experimental.
	StorageConfiguration interface{} `field:"optional" json:"storageConfiguration" yaml:"storageConfiguration"`
}

