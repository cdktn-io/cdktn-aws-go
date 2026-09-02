package awsappmesh


// Experimental.
type TfRoute_SpecHttpRouteMatchHeaderMatchRangeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#end TfRoute#end}.
	// Experimental.
	End *float64 `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#start TfRoute#start}.
	// Experimental.
	Start *float64 `field:"required" json:"start" yaml:"start"`
}

