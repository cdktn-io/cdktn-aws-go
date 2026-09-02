package awsappmesh


// Experimental.
type TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsProperty struct {
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#validation TfVirtualGateway#validation}
	// Experimental.
	Validation *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty `field:"required" json:"validation" yaml:"validation"`
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#certificate TfVirtualGateway#certificate}
	// Experimental.
	Certificate *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateProperty `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#enforce TfVirtualGateway#enforce}.
	// Experimental.
	Enforce interface{} `field:"optional" json:"enforce" yaml:"enforce"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#ports TfVirtualGateway#ports}.
	// Experimental.
	Ports *[]*float64 `field:"optional" json:"ports" yaml:"ports"`
}

