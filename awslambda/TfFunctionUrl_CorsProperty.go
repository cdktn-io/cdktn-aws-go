package awslambda


// Experimental.
type TfFunctionUrl_CorsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function_url#allow_credentials TfFunctionUrl#allow_credentials}.
	// Experimental.
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function_url#allow_headers TfFunctionUrl#allow_headers}.
	// Experimental.
	AllowHeaders *[]*string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function_url#allow_methods TfFunctionUrl#allow_methods}.
	// Experimental.
	AllowMethods *[]*string `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function_url#allow_origins TfFunctionUrl#allow_origins}.
	// Experimental.
	AllowOrigins *[]*string `field:"optional" json:"allowOrigins" yaml:"allowOrigins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function_url#expose_headers TfFunctionUrl#expose_headers}.
	// Experimental.
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function_url#max_age TfFunctionUrl#max_age}.
	// Experimental.
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
}

