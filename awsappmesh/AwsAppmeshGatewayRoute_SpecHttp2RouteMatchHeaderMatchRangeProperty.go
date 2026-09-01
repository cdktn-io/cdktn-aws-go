package awsappmesh


// Experimental.
type AwsAppmeshGatewayRoute_SpecHttp2RouteMatchHeaderMatchRangeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#end AwsAppmeshGatewayRoute#end}.
	// Experimental.
	End *float64 `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#start AwsAppmeshGatewayRoute#start}.
	// Experimental.
	Start *float64 `field:"required" json:"start" yaml:"start"`
}

