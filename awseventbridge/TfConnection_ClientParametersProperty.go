package awseventbridge


// Experimental.
type TfConnection_ClientParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#client_id TfConnection#client_id}.
	// Experimental.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#client_secret TfConnection#client_secret}.
	// Experimental.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
}

