package awsglue


// Experimental.
type TfConnection_Oauth2ClientApplicationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#aws_managed_client_application_reference TfConnection#aws_managed_client_application_reference}.
	// Experimental.
	AwsManagedClientApplicationReference *string `field:"optional" json:"awsManagedClientApplicationReference" yaml:"awsManagedClientApplicationReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#user_managed_client_application_client_id TfConnection#user_managed_client_application_client_id}.
	// Experimental.
	UserManagedClientApplicationClientId *string `field:"optional" json:"userManagedClientApplicationClientId" yaml:"userManagedClientApplicationClientId"`
}

