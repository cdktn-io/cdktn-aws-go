package awsapigatewayv2


// Experimental.
type TfRoutingRule_InvokeApiProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_routing_rule#api_id TfRoutingRule#api_id}.
	// Experimental.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_routing_rule#stage TfRoutingRule#stage}.
	// Experimental.
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_routing_rule#strip_base_path TfRoutingRule#strip_base_path}.
	// Experimental.
	StripBasePath interface{} `field:"optional" json:"stripBasePath" yaml:"stripBasePath"`
}

