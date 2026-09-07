package config


// Experimental.
type AwsRemediationConfiguration_ParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_remediation_configuration#name AwsRemediationConfiguration#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_remediation_configuration#resource_value AwsRemediationConfiguration#resource_value}.
	// Experimental.
	ResourceValue *string `field:"optional" json:"resourceValue" yaml:"resourceValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_remediation_configuration#static_value AwsRemediationConfiguration#static_value}.
	// Experimental.
	StaticValue *string `field:"optional" json:"staticValue" yaml:"staticValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_remediation_configuration#static_values AwsRemediationConfiguration#static_values}.
	// Experimental.
	StaticValues *[]*string `field:"optional" json:"staticValues" yaml:"staticValues"`
}

