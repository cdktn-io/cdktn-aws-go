package awsvpnsitetosite


// Experimental.
type AwsVpnGatewayRoutePropagation_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_gateway_route_propagation#create AwsVpnGatewayRoutePropagation#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_gateway_route_propagation#delete AwsVpnGatewayRoutePropagation#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

