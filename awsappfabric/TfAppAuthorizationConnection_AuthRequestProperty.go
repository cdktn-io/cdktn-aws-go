package awsappfabric


// Experimental.
type TfAppAuthorizationConnection_AuthRequestProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_app_authorization_connection#code TfAppAuthorizationConnection#code}.
	// Experimental.
	Code *string `field:"required" json:"code" yaml:"code"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_app_authorization_connection#redirect_uri TfAppAuthorizationConnection#redirect_uri}.
	// Experimental.
	RedirectUri *string `field:"required" json:"redirectUri" yaml:"redirectUri"`
}

