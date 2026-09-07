package bedrockagentcore


// Experimental.
type AwsGatewayRule_ActionProperty struct {
	// configuration_bundle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#configuration_bundle AwsGatewayRule#configuration_bundle}
	// Experimental.
	ConfigurationBundle interface{} `field:"optional" json:"configurationBundle" yaml:"configurationBundle"`
	// route_to_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#route_to_target AwsGatewayRule#route_to_target}
	// Experimental.
	RouteToTarget interface{} `field:"optional" json:"routeToTarget" yaml:"routeToTarget"`
}

