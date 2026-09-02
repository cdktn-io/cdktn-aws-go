package awsbedrockagents


// Experimental.
type TfFlow_NodeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#name TfFlow#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#type TfFlow#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#configuration TfFlow#configuration}
	// Experimental.
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#input TfFlow#input}
	// Experimental.
	Input interface{} `field:"optional" json:"input" yaml:"input"`
	// output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#output TfFlow#output}
	// Experimental.
	Output interface{} `field:"optional" json:"output" yaml:"output"`
}

