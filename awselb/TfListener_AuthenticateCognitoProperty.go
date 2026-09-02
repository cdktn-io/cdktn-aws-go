package awselb


// Experimental.
type TfListener_AuthenticateCognitoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#user_pool_arn TfListener#user_pool_arn}.
	// Experimental.
	UserPoolArn *string `field:"required" json:"userPoolArn" yaml:"userPoolArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#user_pool_client_id TfListener#user_pool_client_id}.
	// Experimental.
	UserPoolClientId *string `field:"required" json:"userPoolClientId" yaml:"userPoolClientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#user_pool_domain TfListener#user_pool_domain}.
	// Experimental.
	UserPoolDomain *string `field:"required" json:"userPoolDomain" yaml:"userPoolDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#authentication_request_extra_params TfListener#authentication_request_extra_params}.
	// Experimental.
	AuthenticationRequestExtraParams *map[string]*string `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#on_unauthenticated_request TfListener#on_unauthenticated_request}.
	// Experimental.
	OnUnauthenticatedRequest *string `field:"optional" json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#scope TfListener#scope}.
	// Experimental.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#session_cookie_name TfListener#session_cookie_name}.
	// Experimental.
	SessionCookieName *string `field:"optional" json:"sessionCookieName" yaml:"sessionCookieName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#session_timeout TfListener#session_timeout}.
	// Experimental.
	SessionTimeout *float64 `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
}

