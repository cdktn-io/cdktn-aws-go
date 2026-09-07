package appmesh


// Experimental.
type AwsGatewayRoute_SpecHttpRouteMatchHostnameProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#exact AwsGatewayRoute#exact}.
	// Experimental.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#suffix AwsGatewayRoute#suffix}.
	// Experimental.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

