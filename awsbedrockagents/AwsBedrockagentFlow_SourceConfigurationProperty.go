package awsbedrockagents


// Experimental.
type AwsBedrockagentFlow_SourceConfigurationProperty struct {
	// inline block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#inline AwsBedrockagentFlow#inline}
	// Experimental.
	Inline interface{} `field:"optional" json:"inline" yaml:"inline"`
	// resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#resource AwsBedrockagentFlow#resource}
	// Experimental.
	Resource interface{} `field:"optional" json:"resource" yaml:"resource"`
}

