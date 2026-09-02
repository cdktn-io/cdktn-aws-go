package awsbedrockagents


// Experimental.
type TfKnowledgeBase_CuratedQueryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#natural_language TfKnowledgeBase#natural_language}.
	// Experimental.
	NaturalLanguage *string `field:"required" json:"naturalLanguage" yaml:"naturalLanguage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#sql TfKnowledgeBase#sql}.
	// Experimental.
	Sql *string `field:"required" json:"sql" yaml:"sql"`
}

