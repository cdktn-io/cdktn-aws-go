package awsappmesh


// Experimental.
type AwsAppmeshRoute_SpecHttp2RouteMatchHeaderMatchRangeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#end AwsAppmeshRoute#end}.
	// Experimental.
	End *float64 `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#start AwsAppmeshRoute#start}.
	// Experimental.
	Start *float64 `field:"required" json:"start" yaml:"start"`
}

