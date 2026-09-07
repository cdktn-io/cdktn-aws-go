package appmesh


// Experimental.
type AwsGatewayRoute_SpecHttpRouteActionRewriteProperty struct {
	// hostname block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#hostname AwsGatewayRoute#hostname}
	// Experimental.
	Hostname *AwsGatewayRoute_SpecHttpRouteActionRewriteHostnameProperty `field:"optional" json:"hostname" yaml:"hostname"`
	// path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#path AwsGatewayRoute#path}
	// Experimental.
	Path *AwsGatewayRoute_SpecHttpRouteActionRewritePathProperty `field:"optional" json:"path" yaml:"path"`
	// prefix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#prefix AwsGatewayRoute#prefix}
	// Experimental.
	Prefix *AwsGatewayRoute_SpecHttpRouteActionRewritePrefixProperty `field:"optional" json:"prefix" yaml:"prefix"`
}

