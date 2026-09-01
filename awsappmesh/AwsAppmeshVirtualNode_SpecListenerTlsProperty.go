package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_SpecListenerTlsProperty struct {
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#certificate AwsAppmeshVirtualNode#certificate}
	// Experimental.
	Certificate *AwsAppmeshVirtualNode_SpecListenerTlsCertificateProperty `field:"required" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#mode AwsAppmeshVirtualNode#mode}.
	// Experimental.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#validation AwsAppmeshVirtualNode#validation}
	// Experimental.
	Validation *AwsAppmeshVirtualNode_SpecListenerTlsValidationProperty `field:"optional" json:"validation" yaml:"validation"`
}

