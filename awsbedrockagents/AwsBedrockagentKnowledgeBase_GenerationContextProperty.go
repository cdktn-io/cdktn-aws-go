package awsbedrockagents


// Experimental.
type AwsBedrockagentKnowledgeBase_GenerationContextProperty struct {
	// curated_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#curated_query AwsBedrockagentKnowledgeBase#curated_query}
	// Experimental.
	CuratedQuery interface{} `field:"optional" json:"curatedQuery" yaml:"curatedQuery"`
	// table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#table AwsBedrockagentKnowledgeBase#table}
	// Experimental.
	Table interface{} `field:"optional" json:"table" yaml:"table"`
}

