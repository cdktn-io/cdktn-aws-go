package apigatewayv2


// Experimental.
type AwsRoutingRule_AnyOfProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_routing_rule#header AwsRoutingRule#header}.
	// Experimental.
	Header *string `field:"required" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_routing_rule#value_glob AwsRoutingRule#value_glob}.
	// Experimental.
	ValueGlob *string `field:"required" json:"valueGlob" yaml:"valueGlob"`
}

