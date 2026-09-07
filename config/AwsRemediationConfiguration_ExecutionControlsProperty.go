package config


// Experimental.
type AwsRemediationConfiguration_ExecutionControlsProperty struct {
	// ssm_controls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_remediation_configuration#ssm_controls AwsRemediationConfiguration#ssm_controls}
	// Experimental.
	SsmControls *AwsRemediationConfiguration_SsmControlsProperty `field:"optional" json:"ssmControls" yaml:"ssmControls"`
}

