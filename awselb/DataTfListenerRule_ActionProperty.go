package awselb


// Experimental.
type DataTfListenerRule_ActionProperty struct {
	// authenticate_cognito block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#authenticate_cognito DataTfListenerRule#authenticate_cognito}
	// Experimental.
	AuthenticateCognito interface{} `field:"optional" json:"authenticateCognito" yaml:"authenticateCognito"`
	// authenticate_oidc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#authenticate_oidc DataTfListenerRule#authenticate_oidc}
	// Experimental.
	AuthenticateOidc interface{} `field:"optional" json:"authenticateOidc" yaml:"authenticateOidc"`
	// fixed_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#fixed_response DataTfListenerRule#fixed_response}
	// Experimental.
	FixedResponse interface{} `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// forward block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#forward DataTfListenerRule#forward}
	// Experimental.
	Forward interface{} `field:"optional" json:"forward" yaml:"forward"`
	// jwt_validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#jwt_validation DataTfListenerRule#jwt_validation}
	// Experimental.
	JwtValidation interface{} `field:"optional" json:"jwtValidation" yaml:"jwtValidation"`
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#redirect DataTfListenerRule#redirect}
	// Experimental.
	Redirect interface{} `field:"optional" json:"redirect" yaml:"redirect"`
}

