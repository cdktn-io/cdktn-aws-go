package awsappmesh


// Experimental.
type TfRoute_SpecHttp2RouteTimeoutPerRequestProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#unit TfRoute#unit}.
	// Experimental.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#value TfRoute#value}.
	// Experimental.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

