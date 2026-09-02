package awselb


// Experimental.
type TfAlbListener_DefaultActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#type TfAlbListener#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// authenticate_cognito block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#authenticate_cognito TfAlbListener#authenticate_cognito}
	// Experimental.
	AuthenticateCognito *TfAlbListener_AuthenticateCognitoProperty `field:"optional" json:"authenticateCognito" yaml:"authenticateCognito"`
	// authenticate_oidc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#authenticate_oidc TfAlbListener#authenticate_oidc}
	// Experimental.
	AuthenticateOidc *TfAlbListener_AuthenticateOidcProperty `field:"optional" json:"authenticateOidc" yaml:"authenticateOidc"`
	// fixed_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#fixed_response TfAlbListener#fixed_response}
	// Experimental.
	FixedResponse *TfAlbListener_FixedResponseProperty `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// forward block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#forward TfAlbListener#forward}
	// Experimental.
	Forward *TfAlbListener_ForwardProperty `field:"optional" json:"forward" yaml:"forward"`
	// jwt_validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#jwt_validation TfAlbListener#jwt_validation}
	// Experimental.
	JwtValidation *TfAlbListener_JwtValidationProperty `field:"optional" json:"jwtValidation" yaml:"jwtValidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#order TfAlbListener#order}.
	// Experimental.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#redirect TfAlbListener#redirect}
	// Experimental.
	Redirect *TfAlbListener_RedirectProperty `field:"optional" json:"redirect" yaml:"redirect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#target_group_arn TfAlbListener#target_group_arn}.
	// Experimental.
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}

