package awswaf


// Experimental.
type TfWebAcl_RequestBodyProperty struct {
	// api_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#api_gateway TfWebAcl#api_gateway}
	// Experimental.
	ApiGateway *TfWebAcl_ApiGatewayProperty `field:"optional" json:"apiGateway" yaml:"apiGateway"`
	// app_runner_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#app_runner_service TfWebAcl#app_runner_service}
	// Experimental.
	AppRunnerService *TfWebAcl_AppRunnerServiceProperty `field:"optional" json:"appRunnerService" yaml:"appRunnerService"`
	// cloudfront block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#cloudfront TfWebAcl#cloudfront}
	// Experimental.
	Cloudfront *TfWebAcl_CloudfrontProperty `field:"optional" json:"cloudfront" yaml:"cloudfront"`
	// cognito_user_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#cognito_user_pool TfWebAcl#cognito_user_pool}
	// Experimental.
	CognitoUserPool *TfWebAcl_CognitoUserPoolProperty `field:"optional" json:"cognitoUserPool" yaml:"cognitoUserPool"`
	// verified_access_instance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#verified_access_instance TfWebAcl#verified_access_instance}
	// Experimental.
	VerifiedAccessInstance *TfWebAcl_VerifiedAccessInstanceProperty `field:"optional" json:"verifiedAccessInstance" yaml:"verifiedAccessInstance"`
}

