package awsappmesh


// Experimental.
type AwsAppmeshGatewayRoute_SpecHttp2RouteActionRewriteProperty struct {
	// hostname block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#hostname AwsAppmeshGatewayRoute#hostname}
	// Experimental.
	Hostname *AwsAppmeshGatewayRoute_SpecHttp2RouteActionRewriteHostnameProperty `field:"optional" json:"hostname" yaml:"hostname"`
	// path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#path AwsAppmeshGatewayRoute#path}
	// Experimental.
	Path *AwsAppmeshGatewayRoute_SpecHttp2RouteActionRewritePathProperty `field:"optional" json:"path" yaml:"path"`
	// prefix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#prefix AwsAppmeshGatewayRoute#prefix}
	// Experimental.
	Prefix *AwsAppmeshGatewayRoute_SpecHttp2RouteActionRewritePrefixProperty `field:"optional" json:"prefix" yaml:"prefix"`
}

