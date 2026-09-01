package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustProperty struct {
	// acm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#acm AwsAppmeshVirtualNode#acm}
	// Experimental.
	Acm *AwsAppmeshVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustAcmProperty `field:"optional" json:"acm" yaml:"acm"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#file AwsAppmeshVirtualNode#file}
	// Experimental.
	File *AwsAppmeshVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFileProperty `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#sds AwsAppmeshVirtualNode#sds}
	// Experimental.
	Sds *AwsAppmeshVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustSdsProperty `field:"optional" json:"sds" yaml:"sds"`
}

