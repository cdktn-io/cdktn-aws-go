package awsbedrockagents


// Experimental.
type TfKnowledgeBase_NeptuneAnalyticsConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#graph_arn TfKnowledgeBase#graph_arn}.
	// Experimental.
	GraphArn *string `field:"required" json:"graphArn" yaml:"graphArn"`
	// field_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#field_mapping TfKnowledgeBase#field_mapping}
	// Experimental.
	FieldMapping interface{} `field:"optional" json:"fieldMapping" yaml:"fieldMapping"`
}

