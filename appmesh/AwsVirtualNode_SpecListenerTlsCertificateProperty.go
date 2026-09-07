package appmesh


// Experimental.
type AwsVirtualNode_SpecListenerTlsCertificateProperty struct {
	// acm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#acm AwsVirtualNode#acm}
	// Experimental.
	Acm *AwsVirtualNode_SpecListenerTlsCertificateAcmProperty `field:"optional" json:"acm" yaml:"acm"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#file AwsVirtualNode#file}
	// Experimental.
	File *AwsVirtualNode_SpecListenerTlsCertificateFileProperty `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#sds AwsVirtualNode#sds}
	// Experimental.
	Sds *AwsVirtualNode_SpecListenerTlsCertificateSdsProperty `field:"optional" json:"sds" yaml:"sds"`
}

