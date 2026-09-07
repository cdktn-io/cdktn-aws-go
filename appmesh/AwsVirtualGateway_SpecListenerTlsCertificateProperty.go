package appmesh


// Experimental.
type AwsVirtualGateway_SpecListenerTlsCertificateProperty struct {
	// acm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#acm AwsVirtualGateway#acm}
	// Experimental.
	Acm *AwsVirtualGateway_SpecListenerTlsCertificateAcmProperty `field:"optional" json:"acm" yaml:"acm"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#file AwsVirtualGateway#file}
	// Experimental.
	File *AwsVirtualGateway_SpecListenerTlsCertificateFileProperty `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#sds AwsVirtualGateway#sds}
	// Experimental.
	Sds *AwsVirtualGateway_SpecListenerTlsCertificateSdsProperty `field:"optional" json:"sds" yaml:"sds"`
}

