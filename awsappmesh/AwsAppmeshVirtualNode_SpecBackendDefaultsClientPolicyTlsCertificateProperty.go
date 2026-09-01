package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_SpecBackendDefaultsClientPolicyTlsCertificateProperty struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#file AwsAppmeshVirtualNode#file}
	// Experimental.
	File *AwsAppmeshVirtualNode_SpecBackendDefaultsClientPolicyTlsCertificateFileProperty `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#sds AwsAppmeshVirtualNode#sds}
	// Experimental.
	Sds *AwsAppmeshVirtualNode_SpecBackendDefaultsClientPolicyTlsCertificateSdsProperty `field:"optional" json:"sds" yaml:"sds"`
}

