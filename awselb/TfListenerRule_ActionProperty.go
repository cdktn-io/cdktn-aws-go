package awselb


// Experimental.
type TfListenerRule_ActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#type TfListenerRule#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// authenticate_cognito block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#authenticate_cognito TfListenerRule#authenticate_cognito}
	// Experimental.
	AuthenticateCognito *TfListenerRule_AuthenticateCognitoProperty `field:"optional" json:"authenticateCognito" yaml:"authenticateCognito"`
	// authenticate_oidc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#authenticate_oidc TfListenerRule#authenticate_oidc}
	// Experimental.
	AuthenticateOidc *TfListenerRule_AuthenticateOidcProperty `field:"optional" json:"authenticateOidc" yaml:"authenticateOidc"`
	// fixed_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#fixed_response TfListenerRule#fixed_response}
	// Experimental.
	FixedResponse *TfListenerRule_FixedResponseProperty `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// forward block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#forward TfListenerRule#forward}
	// Experimental.
	Forward *TfListenerRule_ForwardProperty `field:"optional" json:"forward" yaml:"forward"`
	// jwt_validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#jwt_validation TfListenerRule#jwt_validation}
	// Experimental.
	JwtValidation *TfListenerRule_JwtValidationProperty `field:"optional" json:"jwtValidation" yaml:"jwtValidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#order TfListenerRule#order}.
	// Experimental.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#redirect TfListenerRule#redirect}
	// Experimental.
	Redirect *TfListenerRule_RedirectProperty `field:"optional" json:"redirect" yaml:"redirect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#target_group_arn TfListenerRule#target_group_arn}.
	// Experimental.
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}

