package appmesh


// Experimental.
type AwsVirtualNode_SpecListenerTlsProperty struct {
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#certificate AwsVirtualNode#certificate}
	// Experimental.
	Certificate *AwsVirtualNode_SpecListenerTlsCertificateProperty `field:"required" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#mode AwsVirtualNode#mode}.
	// Experimental.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#validation AwsVirtualNode#validation}
	// Experimental.
	Validation *AwsVirtualNode_SpecListenerTlsValidationProperty `field:"optional" json:"validation" yaml:"validation"`
}

