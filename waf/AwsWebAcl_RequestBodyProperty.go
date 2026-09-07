package waf


// Experimental.
type AwsWebAcl_RequestBodyProperty struct {
	// api_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#api_gateway AwsWebAcl#api_gateway}
	// Experimental.
	ApiGateway *AwsWebAcl_ApiGatewayProperty `field:"optional" json:"apiGateway" yaml:"apiGateway"`
	// app_runner_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#app_runner_service AwsWebAcl#app_runner_service}
	// Experimental.
	AppRunnerService *AwsWebAcl_AppRunnerServiceProperty `field:"optional" json:"appRunnerService" yaml:"appRunnerService"`
	// cloudfront block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#cloudfront AwsWebAcl#cloudfront}
	// Experimental.
	Cloudfront *AwsWebAcl_CloudfrontProperty `field:"optional" json:"cloudfront" yaml:"cloudfront"`
	// cognito_user_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#cognito_user_pool AwsWebAcl#cognito_user_pool}
	// Experimental.
	CognitoUserPool *AwsWebAcl_CognitoUserPoolProperty `field:"optional" json:"cognitoUserPool" yaml:"cognitoUserPool"`
	// verified_access_instance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#verified_access_instance AwsWebAcl#verified_access_instance}
	// Experimental.
	VerifiedAccessInstance *AwsWebAcl_VerifiedAccessInstanceProperty `field:"optional" json:"verifiedAccessInstance" yaml:"verifiedAccessInstance"`
}

