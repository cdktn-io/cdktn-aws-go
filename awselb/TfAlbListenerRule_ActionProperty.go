package awselb


// Experimental.
type TfAlbListenerRule_ActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#type TfAlbListenerRule#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// authenticate_cognito block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#authenticate_cognito TfAlbListenerRule#authenticate_cognito}
	// Experimental.
	AuthenticateCognito *TfAlbListenerRule_AuthenticateCognitoProperty `field:"optional" json:"authenticateCognito" yaml:"authenticateCognito"`
	// authenticate_oidc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#authenticate_oidc TfAlbListenerRule#authenticate_oidc}
	// Experimental.
	AuthenticateOidc *TfAlbListenerRule_AuthenticateOidcProperty `field:"optional" json:"authenticateOidc" yaml:"authenticateOidc"`
	// fixed_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#fixed_response TfAlbListenerRule#fixed_response}
	// Experimental.
	FixedResponse *TfAlbListenerRule_FixedResponseProperty `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// forward block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#forward TfAlbListenerRule#forward}
	// Experimental.
	Forward *TfAlbListenerRule_ForwardProperty `field:"optional" json:"forward" yaml:"forward"`
	// jwt_validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#jwt_validation TfAlbListenerRule#jwt_validation}
	// Experimental.
	JwtValidation *TfAlbListenerRule_JwtValidationProperty `field:"optional" json:"jwtValidation" yaml:"jwtValidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#order TfAlbListenerRule#order}.
	// Experimental.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#redirect TfAlbListenerRule#redirect}
	// Experimental.
	Redirect *TfAlbListenerRule_RedirectProperty `field:"optional" json:"redirect" yaml:"redirect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#target_group_arn TfAlbListenerRule#target_group_arn}.
	// Experimental.
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}

