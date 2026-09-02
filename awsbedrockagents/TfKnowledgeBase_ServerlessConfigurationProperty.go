package awsbedrockagents


// Experimental.
type TfKnowledgeBase_ServerlessConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#workgroup_arn TfKnowledgeBase#workgroup_arn}.
	// Experimental.
	WorkgroupArn *string `field:"required" json:"workgroupArn" yaml:"workgroupArn"`
	// auth_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#auth_configuration TfKnowledgeBase#auth_configuration}
	// Experimental.
	AuthConfiguration interface{} `field:"optional" json:"authConfiguration" yaml:"authConfiguration"`
}

