package awsappmesh


// Experimental.
type TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty struct {
	// trust block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#trust TfVirtualGateway#trust}
	// Experimental.
	Trust *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty `field:"required" json:"trust" yaml:"trust"`
	// subject_alternative_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#subject_alternative_names TfVirtualGateway#subject_alternative_names}
	// Experimental.
	SubjectAlternativeNames *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

