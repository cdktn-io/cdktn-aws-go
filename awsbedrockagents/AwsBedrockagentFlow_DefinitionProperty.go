package awsbedrockagents


// Experimental.
type AwsBedrockagentFlow_DefinitionProperty struct {
	// connection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#connection AwsBedrockagentFlow#connection}
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#node AwsBedrockagentFlow#node}
	// Experimental.
	NodeAttribute interface{} `field:"optional" json:"nodeAttribute" yaml:"nodeAttribute"`
}

