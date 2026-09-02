package awsappmesh


// Experimental.
type TfGatewayRoute_SpecHttp2RouteMatchHeaderMatchProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#exact TfGatewayRoute#exact}.
	// Experimental.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#prefix TfGatewayRoute#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#range TfGatewayRoute#range}
	// Experimental.
	Range *TfGatewayRoute_SpecHttp2RouteMatchHeaderMatchRangeProperty `field:"optional" json:"range" yaml:"range"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#regex TfGatewayRoute#regex}.
	// Experimental.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#suffix TfGatewayRoute#suffix}.
	// Experimental.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

