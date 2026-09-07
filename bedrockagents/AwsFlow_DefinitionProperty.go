package bedrockagents


// Experimental.
type AwsFlow_DefinitionProperty struct {
	// connection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#connection AwsFlow#connection}
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#node AwsFlow#node}
	// Experimental.
	NodeAttribute interface{} `field:"optional" json:"nodeAttribute" yaml:"nodeAttribute"`
}

