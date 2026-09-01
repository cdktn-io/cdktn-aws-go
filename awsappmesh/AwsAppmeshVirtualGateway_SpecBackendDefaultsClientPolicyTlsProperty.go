package awsappmesh


// Experimental.
type AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsProperty struct {
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#validation AwsAppmeshVirtualGateway#validation}
	// Experimental.
	Validation *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty `field:"required" json:"validation" yaml:"validation"`
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#certificate AwsAppmeshVirtualGateway#certificate}
	// Experimental.
	Certificate *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateProperty `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#enforce AwsAppmeshVirtualGateway#enforce}.
	// Experimental.
	Enforce interface{} `field:"optional" json:"enforce" yaml:"enforce"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#ports AwsAppmeshVirtualGateway#ports}.
	// Experimental.
	Ports *[]*float64 `field:"optional" json:"ports" yaml:"ports"`
}

