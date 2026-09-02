package awsapigatewayv2


// Experimental.
type TfApi_CorsConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_api#allow_credentials TfApi#allow_credentials}.
	// Experimental.
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_api#allow_headers TfApi#allow_headers}.
	// Experimental.
	AllowHeaders *[]*string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_api#allow_methods TfApi#allow_methods}.
	// Experimental.
	AllowMethods *[]*string `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_api#allow_origins TfApi#allow_origins}.
	// Experimental.
	AllowOrigins *[]*string `field:"optional" json:"allowOrigins" yaml:"allowOrigins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_api#expose_headers TfApi#expose_headers}.
	// Experimental.
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_api#max_age TfApi#max_age}.
	// Experimental.
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
}

