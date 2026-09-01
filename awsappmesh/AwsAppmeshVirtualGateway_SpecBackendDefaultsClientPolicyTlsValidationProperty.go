package awsappmesh


// Experimental.
type AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty struct {
	// trust block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#trust AwsAppmeshVirtualGateway#trust}
	// Experimental.
	Trust *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty `field:"required" json:"trust" yaml:"trust"`
	// subject_alternative_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#subject_alternative_names AwsAppmeshVirtualGateway#subject_alternative_names}
	// Experimental.
	SubjectAlternativeNames *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

