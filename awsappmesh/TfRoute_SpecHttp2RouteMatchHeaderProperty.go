package awsappmesh


// Experimental.
type TfRoute_SpecHttp2RouteMatchHeaderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#name TfRoute#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#invert TfRoute#invert}.
	// Experimental.
	Invert interface{} `field:"optional" json:"invert" yaml:"invert"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#match TfRoute#match}
	// Experimental.
	Match *TfRoute_SpecHttp2RouteMatchHeaderMatchProperty `field:"optional" json:"match" yaml:"match"`
}

