package awsappmesh


// Experimental.
type AwsAppmeshVirtualGateway_SpecListenerTlsCertificateProperty struct {
	// acm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#acm AwsAppmeshVirtualGateway#acm}
	// Experimental.
	Acm *AwsAppmeshVirtualGateway_SpecListenerTlsCertificateAcmProperty `field:"optional" json:"acm" yaml:"acm"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#file AwsAppmeshVirtualGateway#file}
	// Experimental.
	File *AwsAppmeshVirtualGateway_SpecListenerTlsCertificateFileProperty `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#sds AwsAppmeshVirtualGateway#sds}
	// Experimental.
	Sds *AwsAppmeshVirtualGateway_SpecListenerTlsCertificateSdsProperty `field:"optional" json:"sds" yaml:"sds"`
}

