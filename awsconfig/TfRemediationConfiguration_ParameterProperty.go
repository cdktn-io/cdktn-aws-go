package awsconfig


// Experimental.
type TfRemediationConfiguration_ParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_remediation_configuration#name TfRemediationConfiguration#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_remediation_configuration#resource_value TfRemediationConfiguration#resource_value}.
	// Experimental.
	ResourceValue *string `field:"optional" json:"resourceValue" yaml:"resourceValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_remediation_configuration#static_value TfRemediationConfiguration#static_value}.
	// Experimental.
	StaticValue *string `field:"optional" json:"staticValue" yaml:"staticValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_remediation_configuration#static_values TfRemediationConfiguration#static_values}.
	// Experimental.
	StaticValues *[]*string `field:"optional" json:"staticValues" yaml:"staticValues"`
}

