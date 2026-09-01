package awsbedrockagents


// Experimental.
type AwsBedrockagentKnowledgeBase_ServerlessConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#workgroup_arn AwsBedrockagentKnowledgeBase#workgroup_arn}.
	// Experimental.
	WorkgroupArn *string `field:"required" json:"workgroupArn" yaml:"workgroupArn"`
	// auth_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#auth_configuration AwsBedrockagentKnowledgeBase#auth_configuration}
	// Experimental.
	AuthConfiguration interface{} `field:"optional" json:"authConfiguration" yaml:"authConfiguration"`
}

