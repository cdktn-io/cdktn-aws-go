package awselb


// Experimental.
type TfListenerRule_AuthenticateCognitoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#user_pool_arn TfListenerRule#user_pool_arn}.
	// Experimental.
	UserPoolArn *string `field:"required" json:"userPoolArn" yaml:"userPoolArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#user_pool_client_id TfListenerRule#user_pool_client_id}.
	// Experimental.
	UserPoolClientId *string `field:"required" json:"userPoolClientId" yaml:"userPoolClientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#user_pool_domain TfListenerRule#user_pool_domain}.
	// Experimental.
	UserPoolDomain *string `field:"required" json:"userPoolDomain" yaml:"userPoolDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#authentication_request_extra_params TfListenerRule#authentication_request_extra_params}.
	// Experimental.
	AuthenticationRequestExtraParams *map[string]*string `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#on_unauthenticated_request TfListenerRule#on_unauthenticated_request}.
	// Experimental.
	OnUnauthenticatedRequest *string `field:"optional" json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#scope TfListenerRule#scope}.
	// Experimental.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#session_cookie_name TfListenerRule#session_cookie_name}.
	// Experimental.
	SessionCookieName *string `field:"optional" json:"sessionCookieName" yaml:"sessionCookieName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#session_timeout TfListenerRule#session_timeout}.
	// Experimental.
	SessionTimeout *float64 `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
}

