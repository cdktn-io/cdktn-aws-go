package awswaf


// Experimental.
type AwsWafv2WebAcl_RequestBodyProperty struct {
	// api_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#api_gateway AwsWafv2WebAcl#api_gateway}
	// Experimental.
	ApiGateway *AwsWafv2WebAcl_ApiGatewayProperty `field:"optional" json:"apiGateway" yaml:"apiGateway"`
	// app_runner_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#app_runner_service AwsWafv2WebAcl#app_runner_service}
	// Experimental.
	AppRunnerService *AwsWafv2WebAcl_AppRunnerServiceProperty `field:"optional" json:"appRunnerService" yaml:"appRunnerService"`
	// cloudfront block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#cloudfront AwsWafv2WebAcl#cloudfront}
	// Experimental.
	Cloudfront *AwsWafv2WebAcl_CloudfrontProperty `field:"optional" json:"cloudfront" yaml:"cloudfront"`
	// cognito_user_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#cognito_user_pool AwsWafv2WebAcl#cognito_user_pool}
	// Experimental.
	CognitoUserPool *AwsWafv2WebAcl_CognitoUserPoolProperty `field:"optional" json:"cognitoUserPool" yaml:"cognitoUserPool"`
	// verified_access_instance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#verified_access_instance AwsWafv2WebAcl#verified_access_instance}
	// Experimental.
	VerifiedAccessInstance *AwsWafv2WebAcl_VerifiedAccessInstanceProperty `field:"optional" json:"verifiedAccessInstance" yaml:"verifiedAccessInstance"`
}

