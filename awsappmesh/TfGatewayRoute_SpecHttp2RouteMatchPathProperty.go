package awsappmesh


// Experimental.
type TfGatewayRoute_SpecHttp2RouteMatchPathProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#exact TfGatewayRoute#exact}.
	// Experimental.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#regex TfGatewayRoute#regex}.
	// Experimental.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
}

