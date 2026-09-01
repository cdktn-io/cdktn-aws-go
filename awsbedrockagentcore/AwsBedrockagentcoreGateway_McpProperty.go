package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreGateway_McpProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#instructions AwsBedrockagentcoreGateway#instructions}.
	// Experimental.
	Instructions *string `field:"optional" json:"instructions" yaml:"instructions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#search_type AwsBedrockagentcoreGateway#search_type}.
	// Experimental.
	SearchType *string `field:"optional" json:"searchType" yaml:"searchType"`
	// session_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#session_configuration AwsBedrockagentcoreGateway#session_configuration}
	// Experimental.
	SessionConfiguration interface{} `field:"optional" json:"sessionConfiguration" yaml:"sessionConfiguration"`
	// streaming_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#streaming_configuration AwsBedrockagentcoreGateway#streaming_configuration}
	// Experimental.
	StreamingConfiguration interface{} `field:"optional" json:"streamingConfiguration" yaml:"streamingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#supported_versions AwsBedrockagentcoreGateway#supported_versions}.
	// Experimental.
	SupportedVersions *[]*string `field:"optional" json:"supportedVersions" yaml:"supportedVersions"`
}

