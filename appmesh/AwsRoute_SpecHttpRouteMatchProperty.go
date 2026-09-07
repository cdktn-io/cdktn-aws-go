package appmesh


// Experimental.
type AwsRoute_SpecHttpRouteMatchProperty struct {
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#header AwsRoute#header}
	// Experimental.
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#method AwsRoute#method}.
	// Experimental.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#path AwsRoute#path}
	// Experimental.
	Path *AwsRoute_SpecHttpRouteMatchPathProperty `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#port AwsRoute#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#prefix AwsRoute#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// query_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#query_parameter AwsRoute#query_parameter}
	// Experimental.
	QueryParameter interface{} `field:"optional" json:"queryParameter" yaml:"queryParameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#scheme AwsRoute#scheme}.
	// Experimental.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

