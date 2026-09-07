package appmesh


// Experimental.
type AwsGatewayRoute_SpecHttp2RouteMatchProperty struct {
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#header AwsGatewayRoute#header}
	// Experimental.
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// hostname block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#hostname AwsGatewayRoute#hostname}
	// Experimental.
	Hostname *AwsGatewayRoute_SpecHttp2RouteMatchHostnameProperty `field:"optional" json:"hostname" yaml:"hostname"`
	// path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#path AwsGatewayRoute#path}
	// Experimental.
	Path *AwsGatewayRoute_SpecHttp2RouteMatchPathProperty `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#port AwsGatewayRoute#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#prefix AwsGatewayRoute#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// query_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#query_parameter AwsGatewayRoute#query_parameter}
	// Experimental.
	QueryParameter interface{} `field:"optional" json:"queryParameter" yaml:"queryParameter"`
}

