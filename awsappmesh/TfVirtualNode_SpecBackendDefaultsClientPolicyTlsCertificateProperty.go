package awsappmesh


// Experimental.
type TfVirtualNode_SpecBackendDefaultsClientPolicyTlsCertificateProperty struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#file TfVirtualNode#file}
	// Experimental.
	File *TfVirtualNode_SpecBackendDefaultsClientPolicyTlsCertificateFileProperty `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#sds TfVirtualNode#sds}
	// Experimental.
	Sds *TfVirtualNode_SpecBackendDefaultsClientPolicyTlsCertificateSdsProperty `field:"optional" json:"sds" yaml:"sds"`
}

