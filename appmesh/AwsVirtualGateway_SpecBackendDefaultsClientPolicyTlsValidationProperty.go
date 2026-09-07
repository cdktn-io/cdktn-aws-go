package appmesh


// Experimental.
type AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty struct {
	// trust block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#trust AwsVirtualGateway#trust}
	// Experimental.
	Trust *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty `field:"required" json:"trust" yaml:"trust"`
	// subject_alternative_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#subject_alternative_names AwsVirtualGateway#subject_alternative_names}
	// Experimental.
	SubjectAlternativeNames *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

