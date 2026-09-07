package appmesh


// Experimental.
type AwsVirtualGateway_SpecListenerTlsValidationTrustProperty struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#file AwsVirtualGateway#file}
	// Experimental.
	File *AwsVirtualGateway_SpecListenerTlsValidationTrustFileProperty `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#sds AwsVirtualGateway#sds}
	// Experimental.
	Sds *AwsVirtualGateway_SpecListenerTlsValidationTrustSdsProperty `field:"optional" json:"sds" yaml:"sds"`
}

