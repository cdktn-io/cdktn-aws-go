package glue


// Experimental.
type AwsConnection_BasicAuthenticationCredentialsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#password AwsConnection#password}.
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#username AwsConnection#username}.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}

