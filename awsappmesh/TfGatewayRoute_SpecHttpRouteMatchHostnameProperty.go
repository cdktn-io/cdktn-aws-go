package awsappmesh


// Experimental.
type TfGatewayRoute_SpecHttpRouteMatchHostnameProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#exact TfGatewayRoute#exact}.
	// Experimental.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#suffix TfGatewayRoute#suffix}.
	// Experimental.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

