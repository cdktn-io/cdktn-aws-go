package awsbedrockagents


// Experimental.
type TfFlow_DefinitionNodeConfigurationProperty struct {
	// agent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#agent TfFlow#agent}
	// Experimental.
	Agent interface{} `field:"optional" json:"agent" yaml:"agent"`
	// collector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#collector TfFlow#collector}
	// Experimental.
	Collector interface{} `field:"optional" json:"collector" yaml:"collector"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#condition TfFlow#condition}
	// Experimental.
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// inline_code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#inline_code TfFlow#inline_code}
	// Experimental.
	InlineCode interface{} `field:"optional" json:"inlineCode" yaml:"inlineCode"`
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#input TfFlow#input}
	// Experimental.
	Input interface{} `field:"optional" json:"input" yaml:"input"`
	// iterator block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#iterator TfFlow#iterator}
	// Experimental.
	Iterator interface{} `field:"optional" json:"iterator" yaml:"iterator"`
	// knowledge_base block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#knowledge_base TfFlow#knowledge_base}
	// Experimental.
	KnowledgeBase interface{} `field:"optional" json:"knowledgeBase" yaml:"knowledgeBase"`
	// lambda_function block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#lambda_function TfFlow#lambda_function}
	// Experimental.
	LambdaFunction interface{} `field:"optional" json:"lambdaFunction" yaml:"lambdaFunction"`
	// lex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#lex TfFlow#lex}
	// Experimental.
	Lex interface{} `field:"optional" json:"lex" yaml:"lex"`
	// output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#output TfFlow#output}
	// Experimental.
	Output interface{} `field:"optional" json:"output" yaml:"output"`
	// prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#prompt TfFlow#prompt}
	// Experimental.
	Prompt interface{} `field:"optional" json:"prompt" yaml:"prompt"`
	// retrieval block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#retrieval TfFlow#retrieval}
	// Experimental.
	Retrieval interface{} `field:"optional" json:"retrieval" yaml:"retrieval"`
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#storage TfFlow#storage}
	// Experimental.
	Storage interface{} `field:"optional" json:"storage" yaml:"storage"`
}

