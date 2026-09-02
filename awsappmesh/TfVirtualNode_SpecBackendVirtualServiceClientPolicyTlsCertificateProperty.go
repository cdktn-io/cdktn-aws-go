package awsappmesh


// Experimental.
type TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsCertificateProperty struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#file TfVirtualNode#file}
	// Experimental.
	File *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsCertificateFileProperty `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#sds TfVirtualNode#sds}
	// Experimental.
	Sds *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsCertificateSdsProperty `field:"optional" json:"sds" yaml:"sds"`
}

