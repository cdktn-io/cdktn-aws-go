package awselb


// Experimental.
type TfListener_DefaultActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#type TfListener#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// authenticate_cognito block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#authenticate_cognito TfListener#authenticate_cognito}
	// Experimental.
	AuthenticateCognito *TfListener_AuthenticateCognitoProperty `field:"optional" json:"authenticateCognito" yaml:"authenticateCognito"`
	// authenticate_oidc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#authenticate_oidc TfListener#authenticate_oidc}
	// Experimental.
	AuthenticateOidc *TfListener_AuthenticateOidcProperty `field:"optional" json:"authenticateOidc" yaml:"authenticateOidc"`
	// fixed_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#fixed_response TfListener#fixed_response}
	// Experimental.
	FixedResponse *TfListener_FixedResponseProperty `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// forward block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#forward TfListener#forward}
	// Experimental.
	Forward *TfListener_ForwardProperty `field:"optional" json:"forward" yaml:"forward"`
	// jwt_validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#jwt_validation TfListener#jwt_validation}
	// Experimental.
	JwtValidation *TfListener_JwtValidationProperty `field:"optional" json:"jwtValidation" yaml:"jwtValidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#order TfListener#order}.
	// Experimental.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#redirect TfListener#redirect}
	// Experimental.
	Redirect *TfListener_RedirectProperty `field:"optional" json:"redirect" yaml:"redirect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#target_group_arn TfListener#target_group_arn}.
	// Experimental.
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}

