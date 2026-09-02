package awssagemakerai


// Experimental.
type TfWorkforce_OidcConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#authorization_endpoint TfWorkforce#authorization_endpoint}.
	// Experimental.
	AuthorizationEndpoint *string `field:"required" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#client_id TfWorkforce#client_id}.
	// Experimental.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#client_secret TfWorkforce#client_secret}.
	// Experimental.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#issuer TfWorkforce#issuer}.
	// Experimental.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#jwks_uri TfWorkforce#jwks_uri}.
	// Experimental.
	JwksUri *string `field:"required" json:"jwksUri" yaml:"jwksUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#logout_endpoint TfWorkforce#logout_endpoint}.
	// Experimental.
	LogoutEndpoint *string `field:"required" json:"logoutEndpoint" yaml:"logoutEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#token_endpoint TfWorkforce#token_endpoint}.
	// Experimental.
	TokenEndpoint *string `field:"required" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#user_info_endpoint TfWorkforce#user_info_endpoint}.
	// Experimental.
	UserInfoEndpoint *string `field:"required" json:"userInfoEndpoint" yaml:"userInfoEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#authentication_request_extra_params TfWorkforce#authentication_request_extra_params}.
	// Experimental.
	AuthenticationRequestExtraParams *map[string]*string `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#scope TfWorkforce#scope}.
	// Experimental.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

