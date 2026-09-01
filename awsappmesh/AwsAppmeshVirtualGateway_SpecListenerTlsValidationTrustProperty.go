package awsappmesh


// Experimental.
type AwsAppmeshVirtualGateway_SpecListenerTlsValidationTrustProperty struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#file AwsAppmeshVirtualGateway#file}
	// Experimental.
	File *AwsAppmeshVirtualGateway_SpecListenerTlsValidationTrustFileProperty `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#sds AwsAppmeshVirtualGateway#sds}
	// Experimental.
	Sds *AwsAppmeshVirtualGateway_SpecListenerTlsValidationTrustSdsProperty `field:"optional" json:"sds" yaml:"sds"`
}

