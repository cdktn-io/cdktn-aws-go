package awsbedrockagents


// Experimental.
type AwsBedrockagentKnowledgeBase_TableProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#name AwsBedrockagentKnowledgeBase#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#column AwsBedrockagentKnowledgeBase#column}
	// Experimental.
	Column interface{} `field:"optional" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#description AwsBedrockagentKnowledgeBase#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#inclusion AwsBedrockagentKnowledgeBase#inclusion}.
	// Experimental.
	Inclusion *string `field:"optional" json:"inclusion" yaml:"inclusion"`
}

