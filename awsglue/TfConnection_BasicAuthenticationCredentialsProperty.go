package awsglue


// Experimental.
type TfConnection_BasicAuthenticationCredentialsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#password TfConnection#password}.
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#username TfConnection#username}.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}

