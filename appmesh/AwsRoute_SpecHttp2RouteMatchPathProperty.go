package appmesh


// Experimental.
type AwsRoute_SpecHttp2RouteMatchPathProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#exact AwsRoute#exact}.
	// Experimental.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#regex AwsRoute#regex}.
	// Experimental.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
}

