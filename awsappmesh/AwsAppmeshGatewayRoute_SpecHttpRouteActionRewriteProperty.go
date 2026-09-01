package awsappmesh


// Experimental.
type AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteProperty struct {
	// hostname block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#hostname AwsAppmeshGatewayRoute#hostname}
	// Experimental.
	Hostname *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteHostnameProperty `field:"optional" json:"hostname" yaml:"hostname"`
	// path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#path AwsAppmeshGatewayRoute#path}
	// Experimental.
	Path *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePathProperty `field:"optional" json:"path" yaml:"path"`
	// prefix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#prefix AwsAppmeshGatewayRoute#prefix}
	// Experimental.
	Prefix *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePrefixProperty `field:"optional" json:"prefix" yaml:"prefix"`
}

