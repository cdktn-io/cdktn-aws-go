package awsappmesh


// Experimental.
type AwsAppmeshRoute_SpecHttp2RouteMatchProperty struct {
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#header AwsAppmeshRoute#header}
	// Experimental.
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#method AwsAppmeshRoute#method}.
	// Experimental.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#path AwsAppmeshRoute#path}
	// Experimental.
	Path *AwsAppmeshRoute_SpecHttp2RouteMatchPathProperty `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#port AwsAppmeshRoute#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#prefix AwsAppmeshRoute#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// query_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#query_parameter AwsAppmeshRoute#query_parameter}
	// Experimental.
	QueryParameter interface{} `field:"optional" json:"queryParameter" yaml:"queryParameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#scheme AwsAppmeshRoute#scheme}.
	// Experimental.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

