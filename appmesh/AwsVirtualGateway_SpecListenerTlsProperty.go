package appmesh


// Experimental.
type AwsVirtualGateway_SpecListenerTlsProperty struct {
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#certificate AwsVirtualGateway#certificate}
	// Experimental.
	Certificate *AwsVirtualGateway_SpecListenerTlsCertificateProperty `field:"required" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#mode AwsVirtualGateway#mode}.
	// Experimental.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#validation AwsVirtualGateway#validation}
	// Experimental.
	Validation *AwsVirtualGateway_SpecListenerTlsValidationProperty `field:"optional" json:"validation" yaml:"validation"`
}

