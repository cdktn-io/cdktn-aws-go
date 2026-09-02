package awsredshiftserverless


// Experimental.
type TfWorkgroup_ConfigParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_workgroup#parameter_key TfWorkgroup#parameter_key}.
	// Experimental.
	ParameterKey *string `field:"required" json:"parameterKey" yaml:"parameterKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_workgroup#parameter_value TfWorkgroup#parameter_value}.
	// Experimental.
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

