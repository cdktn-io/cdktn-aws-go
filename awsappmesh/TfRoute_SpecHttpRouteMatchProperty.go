package awsappmesh


// Experimental.
type TfRoute_SpecHttpRouteMatchProperty struct {
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#header TfRoute#header}
	// Experimental.
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#method TfRoute#method}.
	// Experimental.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#path TfRoute#path}
	// Experimental.
	Path *TfRoute_SpecHttpRouteMatchPathProperty `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#port TfRoute#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#prefix TfRoute#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// query_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#query_parameter TfRoute#query_parameter}
	// Experimental.
	QueryParameter interface{} `field:"optional" json:"queryParameter" yaml:"queryParameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#scheme TfRoute#scheme}.
	// Experimental.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

