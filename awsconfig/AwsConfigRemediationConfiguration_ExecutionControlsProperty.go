package awsconfig


// Experimental.
type AwsConfigRemediationConfiguration_ExecutionControlsProperty struct {
	// ssm_controls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_remediation_configuration#ssm_controls AwsConfigRemediationConfiguration#ssm_controls}
	// Experimental.
	SsmControls *AwsConfigRemediationConfiguration_SsmControlsProperty `field:"optional" json:"ssmControls" yaml:"ssmControls"`
}

