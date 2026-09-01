package awsappmesh


// Experimental.
type AwsAppmeshVirtualGateway_SpecListenerTlsValidationTrustSdsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#secret_name AwsAppmeshVirtualGateway#secret_name}.
	// Experimental.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

