package awseventbridge


// Experimental.
type TfConnection_BasicProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#password TfConnection#password}.
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#username TfConnection#username}.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}

