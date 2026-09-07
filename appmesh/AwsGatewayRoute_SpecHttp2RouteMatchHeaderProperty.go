package appmesh


// Experimental.
type AwsGatewayRoute_SpecHttp2RouteMatchHeaderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#name AwsGatewayRoute#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#invert AwsGatewayRoute#invert}.
	// Experimental.
	Invert interface{} `field:"optional" json:"invert" yaml:"invert"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#match AwsGatewayRoute#match}
	// Experimental.
	Match *AwsGatewayRoute_SpecHttp2RouteMatchHeaderMatchProperty `field:"optional" json:"match" yaml:"match"`
}

