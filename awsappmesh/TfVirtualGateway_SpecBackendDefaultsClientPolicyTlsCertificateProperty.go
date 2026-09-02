package awsappmesh


// Experimental.
type TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateProperty struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#file TfVirtualGateway#file}
	// Experimental.
	File *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateFileProperty `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#sds TfVirtualGateway#sds}
	// Experimental.
	Sds *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateSdsProperty `field:"optional" json:"sds" yaml:"sds"`
}

