package awsconfig


// Experimental.
type TfOrganizationConformancePack_InputParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_conformance_pack#parameter_name TfOrganizationConformancePack#parameter_name}.
	// Experimental.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_conformance_pack#parameter_value TfOrganizationConformancePack#parameter_value}.
	// Experimental.
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

