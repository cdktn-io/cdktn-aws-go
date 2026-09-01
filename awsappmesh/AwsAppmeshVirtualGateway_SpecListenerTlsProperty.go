package awsappmesh


// Experimental.
type AwsAppmeshVirtualGateway_SpecListenerTlsProperty struct {
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#certificate AwsAppmeshVirtualGateway#certificate}
	// Experimental.
	Certificate *AwsAppmeshVirtualGateway_SpecListenerTlsCertificateProperty `field:"required" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#mode AwsAppmeshVirtualGateway#mode}.
	// Experimental.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#validation AwsAppmeshVirtualGateway#validation}
	// Experimental.
	Validation *AwsAppmeshVirtualGateway_SpecListenerTlsValidationProperty `field:"optional" json:"validation" yaml:"validation"`
}

