package appmesh


// Experimental.
type AwsVirtualNode_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty struct {
	// acm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#acm AwsVirtualNode#acm}
	// Experimental.
	Acm *AwsVirtualNode_SpecBackendDefaultsClientPolicyTlsValidationTrustAcmProperty `field:"optional" json:"acm" yaml:"acm"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#file AwsVirtualNode#file}
	// Experimental.
	File *AwsVirtualNode_SpecBackendDefaultsClientPolicyTlsValidationTrustFileProperty `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#sds AwsVirtualNode#sds}
	// Experimental.
	Sds *AwsVirtualNode_SpecBackendDefaultsClientPolicyTlsValidationTrustSdsProperty `field:"optional" json:"sds" yaml:"sds"`
}

