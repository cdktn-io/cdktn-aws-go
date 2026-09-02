package awsbedrockagentcore


// Experimental.
type TfRegistry_CustomJwtAuthorizerProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_registry#discovery_url TfRegistry#discovery_url}.
	// Experimental.
	DiscoveryUrl *string `field:"required" json:"discoveryUrl" yaml:"discoveryUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_registry#allowed_audience TfRegistry#allowed_audience}.
	// Experimental.
	AllowedAudience *[]*string `field:"optional" json:"allowedAudience" yaml:"allowedAudience"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_registry#allowed_clients TfRegistry#allowed_clients}.
	// Experimental.
	AllowedClients *[]*string `field:"optional" json:"allowedClients" yaml:"allowedClients"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_registry#allowed_scopes TfRegistry#allowed_scopes}.
	// Experimental.
	AllowedScopes *[]*string `field:"optional" json:"allowedScopes" yaml:"allowedScopes"`
	// allowed_workload_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_registry#allowed_workload_configuration TfRegistry#allowed_workload_configuration}
	// Experimental.
	AllowedWorkloadConfiguration interface{} `field:"optional" json:"allowedWorkloadConfiguration" yaml:"allowedWorkloadConfiguration"`
	// custom_claim block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_registry#custom_claim TfRegistry#custom_claim}
	// Experimental.
	CustomClaim interface{} `field:"optional" json:"customClaim" yaml:"customClaim"`
	// private_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_registry#private_endpoint TfRegistry#private_endpoint}
	// Experimental.
	PrivateEndpoint interface{} `field:"optional" json:"privateEndpoint" yaml:"privateEndpoint"`
	// private_endpoint_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_registry#private_endpoint_overrides TfRegistry#private_endpoint_overrides}
	// Experimental.
	PrivateEndpointOverrides interface{} `field:"optional" json:"privateEndpointOverrides" yaml:"privateEndpointOverrides"`
}

