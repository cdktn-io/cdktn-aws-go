package apigatewayv2


// Experimental.
type AwsIntegration_ResponseParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_integration#mappings AwsIntegration#mappings}.
	// Experimental.
	Mappings *map[string]*string `field:"required" json:"mappings" yaml:"mappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_integration#status_code AwsIntegration#status_code}.
	// Experimental.
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
}

