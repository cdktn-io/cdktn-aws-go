package awsappmesh


// Experimental.
type TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustProperty struct {
	// acm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#acm TfVirtualNode#acm}
	// Experimental.
	Acm *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustAcmProperty `field:"optional" json:"acm" yaml:"acm"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#file TfVirtualNode#file}
	// Experimental.
	File *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFileProperty `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#sds TfVirtualNode#sds}
	// Experimental.
	Sds *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustSdsProperty `field:"optional" json:"sds" yaml:"sds"`
}

