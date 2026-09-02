package awsappmesh


// Experimental.
type TfGatewayRoute_SpecHttp2RouteActionRewriteProperty struct {
	// hostname block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#hostname TfGatewayRoute#hostname}
	// Experimental.
	Hostname *TfGatewayRoute_SpecHttp2RouteActionRewriteHostnameProperty `field:"optional" json:"hostname" yaml:"hostname"`
	// path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#path TfGatewayRoute#path}
	// Experimental.
	Path *TfGatewayRoute_SpecHttp2RouteActionRewritePathProperty `field:"optional" json:"path" yaml:"path"`
	// prefix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#prefix TfGatewayRoute#prefix}
	// Experimental.
	Prefix *TfGatewayRoute_SpecHttp2RouteActionRewritePrefixProperty `field:"optional" json:"prefix" yaml:"prefix"`
}

