package awsappmesh


// Experimental.
type TfRoute_SpecHttp2RouteMatchHeaderMatchProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#exact TfRoute#exact}.
	// Experimental.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#prefix TfRoute#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#range TfRoute#range}
	// Experimental.
	Range *TfRoute_SpecHttp2RouteMatchHeaderMatchRangeProperty `field:"optional" json:"range" yaml:"range"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#regex TfRoute#regex}.
	// Experimental.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#suffix TfRoute#suffix}.
	// Experimental.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

