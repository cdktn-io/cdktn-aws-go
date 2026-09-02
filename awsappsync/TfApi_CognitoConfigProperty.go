package awsappsync


// Experimental.
type TfApi_CognitoConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#aws_region TfApi#aws_region}.
	// Experimental.
	AwsRegion *string `field:"required" json:"awsRegion" yaml:"awsRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#user_pool_id TfApi#user_pool_id}.
	// Experimental.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#app_id_client_regex TfApi#app_id_client_regex}.
	// Experimental.
	AppIdClientRegex *string `field:"optional" json:"appIdClientRegex" yaml:"appIdClientRegex"`
}

