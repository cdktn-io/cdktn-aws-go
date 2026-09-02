package awsbedrockagentcore


// Experimental.
type TfGateway_McpProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#instructions TfGateway#instructions}.
	// Experimental.
	Instructions *string `field:"optional" json:"instructions" yaml:"instructions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#search_type TfGateway#search_type}.
	// Experimental.
	SearchType *string `field:"optional" json:"searchType" yaml:"searchType"`
	// session_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#session_configuration TfGateway#session_configuration}
	// Experimental.
	SessionConfiguration interface{} `field:"optional" json:"sessionConfiguration" yaml:"sessionConfiguration"`
	// streaming_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#streaming_configuration TfGateway#streaming_configuration}
	// Experimental.
	StreamingConfiguration interface{} `field:"optional" json:"streamingConfiguration" yaml:"streamingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#supported_versions TfGateway#supported_versions}.
	// Experimental.
	SupportedVersions *[]*string `field:"optional" json:"supportedVersions" yaml:"supportedVersions"`
}

