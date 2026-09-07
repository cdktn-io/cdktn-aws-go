package appmesh


// Experimental.
type AwsRoute_SpecHttp2RouteMatchQueryParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#name AwsRoute#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#match AwsRoute#match}
	// Experimental.
	Match *AwsRoute_SpecHttp2RouteMatchQueryParameterMatchProperty `field:"optional" json:"match" yaml:"match"`
}

