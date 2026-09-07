package glue


// Experimental.
type AwsConnection_AuthorizationCodePropertiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#authorization_code AwsConnection#authorization_code}.
	// Experimental.
	AuthorizationCode *string `field:"required" json:"authorizationCode" yaml:"authorizationCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#redirect_uri AwsConnection#redirect_uri}.
	// Experimental.
	RedirectUri *string `field:"required" json:"redirectUri" yaml:"redirectUri"`
}

