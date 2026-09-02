package awsappmesh


// Experimental.
type TfVirtualGateway_SpecListenerTlsProperty struct {
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#certificate TfVirtualGateway#certificate}
	// Experimental.
	Certificate *TfVirtualGateway_SpecListenerTlsCertificateProperty `field:"required" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#mode TfVirtualGateway#mode}.
	// Experimental.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#validation TfVirtualGateway#validation}
	// Experimental.
	Validation *TfVirtualGateway_SpecListenerTlsValidationProperty `field:"optional" json:"validation" yaml:"validation"`
}

