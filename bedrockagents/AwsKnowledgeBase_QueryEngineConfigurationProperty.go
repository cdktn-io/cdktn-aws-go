package bedrockagents


// Experimental.
type AwsKnowledgeBase_QueryEngineConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#type AwsKnowledgeBase#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// provisioned_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#provisioned_configuration AwsKnowledgeBase#provisioned_configuration}
	// Experimental.
	ProvisionedConfiguration interface{} `field:"optional" json:"provisionedConfiguration" yaml:"provisionedConfiguration"`
	// serverless_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#serverless_configuration AwsKnowledgeBase#serverless_configuration}
	// Experimental.
	ServerlessConfiguration interface{} `field:"optional" json:"serverlessConfiguration" yaml:"serverlessConfiguration"`
}

