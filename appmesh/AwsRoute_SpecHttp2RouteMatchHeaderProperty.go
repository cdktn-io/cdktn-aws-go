package appmesh


// Experimental.
type AwsRoute_SpecHttp2RouteMatchHeaderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#name AwsRoute#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#invert AwsRoute#invert}.
	// Experimental.
	Invert interface{} `field:"optional" json:"invert" yaml:"invert"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#match AwsRoute#match}
	// Experimental.
	Match *AwsRoute_SpecHttp2RouteMatchHeaderMatchProperty `field:"optional" json:"match" yaml:"match"`
}

