package awslexv2models


// Experimental.
type AwsLexv2ModelsIntent_DataSourceConfigurationProperty struct {
	// bedrock_knowledge_store_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#bedrock_knowledge_store_configuration AwsLexv2ModelsIntent#bedrock_knowledge_store_configuration}
	// Experimental.
	BedrockKnowledgeStoreConfiguration interface{} `field:"optional" json:"bedrockKnowledgeStoreConfiguration" yaml:"bedrockKnowledgeStoreConfiguration"`
	// kendra_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#kendra_configuration AwsLexv2ModelsIntent#kendra_configuration}
	// Experimental.
	KendraConfiguration interface{} `field:"optional" json:"kendraConfiguration" yaml:"kendraConfiguration"`
	// opensearch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#opensearch_configuration AwsLexv2ModelsIntent#opensearch_configuration}
	// Experimental.
	OpensearchConfiguration interface{} `field:"optional" json:"opensearchConfiguration" yaml:"opensearchConfiguration"`
}

