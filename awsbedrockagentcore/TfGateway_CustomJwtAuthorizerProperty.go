package awsbedrockagentcore


// Experimental.
type TfGateway_CustomJwtAuthorizerProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#discovery_url TfGateway#discovery_url}.
	// Experimental.
	DiscoveryUrl *string `field:"required" json:"discoveryUrl" yaml:"discoveryUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#allowed_audience TfGateway#allowed_audience}.
	// Experimental.
	AllowedAudience *[]*string `field:"optional" json:"allowedAudience" yaml:"allowedAudience"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#allowed_clients TfGateway#allowed_clients}.
	// Experimental.
	AllowedClients *[]*string `field:"optional" json:"allowedClients" yaml:"allowedClients"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#allowed_scopes TfGateway#allowed_scopes}.
	// Experimental.
	AllowedScopes *[]*string `field:"optional" json:"allowedScopes" yaml:"allowedScopes"`
	// allowed_workload_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#allowed_workload_configuration TfGateway#allowed_workload_configuration}
	// Experimental.
	AllowedWorkloadConfiguration interface{} `field:"optional" json:"allowedWorkloadConfiguration" yaml:"allowedWorkloadConfiguration"`
	// custom_claim block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#custom_claim TfGateway#custom_claim}
	// Experimental.
	CustomClaim interface{} `field:"optional" json:"customClaim" yaml:"customClaim"`
	// private_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#private_endpoint TfGateway#private_endpoint}
	// Experimental.
	PrivateEndpoint interface{} `field:"optional" json:"privateEndpoint" yaml:"privateEndpoint"`
	// private_endpoint_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#private_endpoint_overrides TfGateway#private_endpoint_overrides}
	// Experimental.
	PrivateEndpointOverrides interface{} `field:"optional" json:"privateEndpointOverrides" yaml:"privateEndpointOverrides"`
}

