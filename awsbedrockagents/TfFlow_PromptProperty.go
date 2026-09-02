package awsbedrockagents


// Experimental.
type TfFlow_PromptProperty struct {
	// guardrail_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#guardrail_configuration TfFlow#guardrail_configuration}
	// Experimental.
	GuardrailConfiguration interface{} `field:"optional" json:"guardrailConfiguration" yaml:"guardrailConfiguration"`
	// source_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#source_configuration TfFlow#source_configuration}
	// Experimental.
	SourceConfiguration interface{} `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
}

