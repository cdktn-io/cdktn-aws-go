package awsbedrockagents


// Experimental.
type TfKnowledgeBase_SqlKnowledgeBaseConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#type TfKnowledgeBase#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// redshift_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#redshift_configuration TfKnowledgeBase#redshift_configuration}
	// Experimental.
	RedshiftConfiguration interface{} `field:"optional" json:"redshiftConfiguration" yaml:"redshiftConfiguration"`
}

