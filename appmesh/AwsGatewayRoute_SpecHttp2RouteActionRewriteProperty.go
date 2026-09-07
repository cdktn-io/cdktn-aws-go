package appmesh


// Experimental.
type AwsGatewayRoute_SpecHttp2RouteActionRewriteProperty struct {
	// hostname block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#hostname AwsGatewayRoute#hostname}
	// Experimental.
	Hostname *AwsGatewayRoute_SpecHttp2RouteActionRewriteHostnameProperty `field:"optional" json:"hostname" yaml:"hostname"`
	// path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#path AwsGatewayRoute#path}
	// Experimental.
	Path *AwsGatewayRoute_SpecHttp2RouteActionRewritePathProperty `field:"optional" json:"path" yaml:"path"`
	// prefix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#prefix AwsGatewayRoute#prefix}
	// Experimental.
	Prefix *AwsGatewayRoute_SpecHttp2RouteActionRewritePrefixProperty `field:"optional" json:"prefix" yaml:"prefix"`
}

