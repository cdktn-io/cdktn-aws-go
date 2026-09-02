package awsglue


// Experimental.
type TfConnection_AuthorizationCodePropertiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#authorization_code TfConnection#authorization_code}.
	// Experimental.
	AuthorizationCode *string `field:"required" json:"authorizationCode" yaml:"authorizationCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#redirect_uri TfConnection#redirect_uri}.
	// Experimental.
	RedirectUri *string `field:"required" json:"redirectUri" yaml:"redirectUri"`
}

