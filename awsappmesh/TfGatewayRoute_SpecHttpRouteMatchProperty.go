package awsappmesh


// Experimental.
type TfGatewayRoute_SpecHttpRouteMatchProperty struct {
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#header TfGatewayRoute#header}
	// Experimental.
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// hostname block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#hostname TfGatewayRoute#hostname}
	// Experimental.
	Hostname *TfGatewayRoute_SpecHttpRouteMatchHostnameProperty `field:"optional" json:"hostname" yaml:"hostname"`
	// path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#path TfGatewayRoute#path}
	// Experimental.
	Path *TfGatewayRoute_SpecHttpRouteMatchPathProperty `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#port TfGatewayRoute#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#prefix TfGatewayRoute#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// query_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#query_parameter TfGatewayRoute#query_parameter}
	// Experimental.
	QueryParameter interface{} `field:"optional" json:"queryParameter" yaml:"queryParameter"`
}

