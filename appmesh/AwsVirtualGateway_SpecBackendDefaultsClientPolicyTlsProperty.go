package appmesh


// Experimental.
type AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsProperty struct {
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#validation AwsVirtualGateway#validation}
	// Experimental.
	Validation *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty `field:"required" json:"validation" yaml:"validation"`
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#certificate AwsVirtualGateway#certificate}
	// Experimental.
	Certificate *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateProperty `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#enforce AwsVirtualGateway#enforce}.
	// Experimental.
	Enforce interface{} `field:"optional" json:"enforce" yaml:"enforce"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#ports AwsVirtualGateway#ports}.
	// Experimental.
	Ports *[]*float64 `field:"optional" json:"ports" yaml:"ports"`
}

