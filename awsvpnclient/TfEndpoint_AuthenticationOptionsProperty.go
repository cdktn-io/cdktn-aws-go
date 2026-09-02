package awsvpnclient


// Experimental.
type TfEndpoint_AuthenticationOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#type TfEndpoint#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#active_directory_id TfEndpoint#active_directory_id}.
	// Experimental.
	ActiveDirectoryId *string `field:"optional" json:"activeDirectoryId" yaml:"activeDirectoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#root_certificate_chain_arn TfEndpoint#root_certificate_chain_arn}.
	// Experimental.
	RootCertificateChainArn *string `field:"optional" json:"rootCertificateChainArn" yaml:"rootCertificateChainArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#saml_provider_arn TfEndpoint#saml_provider_arn}.
	// Experimental.
	SamlProviderArn *string `field:"optional" json:"samlProviderArn" yaml:"samlProviderArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#self_service_saml_provider_arn TfEndpoint#self_service_saml_provider_arn}.
	// Experimental.
	SelfServiceSamlProviderArn *string `field:"optional" json:"selfServiceSamlProviderArn" yaml:"selfServiceSamlProviderArn"`
}

