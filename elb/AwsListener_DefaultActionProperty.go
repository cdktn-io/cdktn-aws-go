package elb


// Experimental.
type AwsListener_DefaultActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#type AwsListener#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// authenticate_cognito block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#authenticate_cognito AwsListener#authenticate_cognito}
	// Experimental.
	AuthenticateCognito *AwsListener_AuthenticateCognitoProperty `field:"optional" json:"authenticateCognito" yaml:"authenticateCognito"`
	// authenticate_oidc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#authenticate_oidc AwsListener#authenticate_oidc}
	// Experimental.
	AuthenticateOidc *AwsListener_AuthenticateOidcProperty `field:"optional" json:"authenticateOidc" yaml:"authenticateOidc"`
	// fixed_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#fixed_response AwsListener#fixed_response}
	// Experimental.
	FixedResponse *AwsListener_FixedResponseProperty `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// forward block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#forward AwsListener#forward}
	// Experimental.
	Forward *AwsListener_ForwardProperty `field:"optional" json:"forward" yaml:"forward"`
	// jwt_validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#jwt_validation AwsListener#jwt_validation}
	// Experimental.
	JwtValidation *AwsListener_JwtValidationProperty `field:"optional" json:"jwtValidation" yaml:"jwtValidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#order AwsListener#order}.
	// Experimental.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#redirect AwsListener#redirect}
	// Experimental.
	Redirect *AwsListener_RedirectProperty `field:"optional" json:"redirect" yaml:"redirect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#target_group_arn AwsListener#target_group_arn}.
	// Experimental.
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}

