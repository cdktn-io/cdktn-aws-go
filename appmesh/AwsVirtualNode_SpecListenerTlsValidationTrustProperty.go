package appmesh


// Experimental.
type AwsVirtualNode_SpecListenerTlsValidationTrustProperty struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#file AwsVirtualNode#file}
	// Experimental.
	File *AwsVirtualNode_SpecListenerTlsValidationTrustFileProperty `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#sds AwsVirtualNode#sds}
	// Experimental.
	Sds *AwsVirtualNode_SpecListenerTlsValidationTrustSdsProperty `field:"optional" json:"sds" yaml:"sds"`
}

