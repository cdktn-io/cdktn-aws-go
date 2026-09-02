package awsappmesh


// Experimental.
type TfVirtualGateway_SpecListenerTlsCertificateProperty struct {
	// acm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#acm TfVirtualGateway#acm}
	// Experimental.
	Acm *TfVirtualGateway_SpecListenerTlsCertificateAcmProperty `field:"optional" json:"acm" yaml:"acm"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#file TfVirtualGateway#file}
	// Experimental.
	File *TfVirtualGateway_SpecListenerTlsCertificateFileProperty `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#sds TfVirtualGateway#sds}
	// Experimental.
	Sds *TfVirtualGateway_SpecListenerTlsCertificateSdsProperty `field:"optional" json:"sds" yaml:"sds"`
}

