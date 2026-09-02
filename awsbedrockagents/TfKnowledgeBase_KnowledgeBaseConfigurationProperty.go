package awsbedrockagents


// Experimental.
type TfKnowledgeBase_KnowledgeBaseConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#type TfKnowledgeBase#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// kendra_knowledge_base_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#kendra_knowledge_base_configuration TfKnowledgeBase#kendra_knowledge_base_configuration}
	// Experimental.
	KendraKnowledgeBaseConfiguration interface{} `field:"optional" json:"kendraKnowledgeBaseConfiguration" yaml:"kendraKnowledgeBaseConfiguration"`
	// managed_knowledge_base_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#managed_knowledge_base_configuration TfKnowledgeBase#managed_knowledge_base_configuration}
	// Experimental.
	ManagedKnowledgeBaseConfiguration interface{} `field:"optional" json:"managedKnowledgeBaseConfiguration" yaml:"managedKnowledgeBaseConfiguration"`
	// sql_knowledge_base_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#sql_knowledge_base_configuration TfKnowledgeBase#sql_knowledge_base_configuration}
	// Experimental.
	SqlKnowledgeBaseConfiguration interface{} `field:"optional" json:"sqlKnowledgeBaseConfiguration" yaml:"sqlKnowledgeBaseConfiguration"`
	// vector_knowledge_base_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#vector_knowledge_base_configuration TfKnowledgeBase#vector_knowledge_base_configuration}
	// Experimental.
	VectorKnowledgeBaseConfiguration interface{} `field:"optional" json:"vectorKnowledgeBaseConfiguration" yaml:"vectorKnowledgeBaseConfiguration"`
}

