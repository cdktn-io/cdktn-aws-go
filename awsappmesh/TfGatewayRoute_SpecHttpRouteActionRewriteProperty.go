package awsappmesh


// Experimental.
type TfGatewayRoute_SpecHttpRouteActionRewriteProperty struct {
	// hostname block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#hostname TfGatewayRoute#hostname}
	// Experimental.
	Hostname *TfGatewayRoute_SpecHttpRouteActionRewriteHostnameProperty `field:"optional" json:"hostname" yaml:"hostname"`
	// path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#path TfGatewayRoute#path}
	// Experimental.
	Path *TfGatewayRoute_SpecHttpRouteActionRewritePathProperty `field:"optional" json:"path" yaml:"path"`
	// prefix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#prefix TfGatewayRoute#prefix}
	// Experimental.
	Prefix *TfGatewayRoute_SpecHttpRouteActionRewritePrefixProperty `field:"optional" json:"prefix" yaml:"prefix"`
}

