package awsappmesh


// Experimental.
type TfVirtualNode_SpecListenerTlsProperty struct {
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#certificate TfVirtualNode#certificate}
	// Experimental.
	Certificate *TfVirtualNode_SpecListenerTlsCertificateProperty `field:"required" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#mode TfVirtualNode#mode}.
	// Experimental.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#validation TfVirtualNode#validation}
	// Experimental.
	Validation *TfVirtualNode_SpecListenerTlsValidationProperty `field:"optional" json:"validation" yaml:"validation"`
}

