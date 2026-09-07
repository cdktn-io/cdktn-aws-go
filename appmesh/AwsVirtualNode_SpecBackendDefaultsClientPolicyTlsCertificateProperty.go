package appmesh


// Experimental.
type AwsVirtualNode_SpecBackendDefaultsClientPolicyTlsCertificateProperty struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#file AwsVirtualNode#file}
	// Experimental.
	File *AwsVirtualNode_SpecBackendDefaultsClientPolicyTlsCertificateFileProperty `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#sds AwsVirtualNode#sds}
	// Experimental.
	Sds *AwsVirtualNode_SpecBackendDefaultsClientPolicyTlsCertificateSdsProperty `field:"optional" json:"sds" yaml:"sds"`
}

