package bedrockagents


// Experimental.
type AwsKnowledgeBase_KnowledgeBaseConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#type AwsKnowledgeBase#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// kendra_knowledge_base_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#kendra_knowledge_base_configuration AwsKnowledgeBase#kendra_knowledge_base_configuration}
	// Experimental.
	KendraKnowledgeBaseConfiguration interface{} `field:"optional" json:"kendraKnowledgeBaseConfiguration" yaml:"kendraKnowledgeBaseConfiguration"`
	// managed_knowledge_base_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#managed_knowledge_base_configuration AwsKnowledgeBase#managed_knowledge_base_configuration}
	// Experimental.
	ManagedKnowledgeBaseConfiguration interface{} `field:"optional" json:"managedKnowledgeBaseConfiguration" yaml:"managedKnowledgeBaseConfiguration"`
	// sql_knowledge_base_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#sql_knowledge_base_configuration AwsKnowledgeBase#sql_knowledge_base_configuration}
	// Experimental.
	SqlKnowledgeBaseConfiguration interface{} `field:"optional" json:"sqlKnowledgeBaseConfiguration" yaml:"sqlKnowledgeBaseConfiguration"`
	// vector_knowledge_base_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#vector_knowledge_base_configuration AwsKnowledgeBase#vector_knowledge_base_configuration}
	// Experimental.
	VectorKnowledgeBaseConfiguration interface{} `field:"optional" json:"vectorKnowledgeBaseConfiguration" yaml:"vectorKnowledgeBaseConfiguration"`
}

