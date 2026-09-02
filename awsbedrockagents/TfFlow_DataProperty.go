package awsbedrockagents


// Experimental.
type TfFlow_DataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#source_output TfFlow#source_output}.
	// Experimental.
	SourceOutput *string `field:"required" json:"sourceOutput" yaml:"sourceOutput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#target_input TfFlow#target_input}.
	// Experimental.
	TargetInput *string `field:"required" json:"targetInput" yaml:"targetInput"`
}

