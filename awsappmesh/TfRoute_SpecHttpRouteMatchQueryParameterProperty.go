package awsappmesh


// Experimental.
type TfRoute_SpecHttpRouteMatchQueryParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#name TfRoute#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#match TfRoute#match}
	// Experimental.
	Match *TfRoute_SpecHttpRouteMatchQueryParameterMatchProperty `field:"optional" json:"match" yaml:"match"`
}

