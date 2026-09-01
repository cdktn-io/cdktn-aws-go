package awsapprunner


// Experimental.
type AwsApprunnerService_CodeConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#configuration_source AwsApprunnerService#configuration_source}.
	// Experimental.
	ConfigurationSource *string `field:"required" json:"configurationSource" yaml:"configurationSource"`
	// code_configuration_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#code_configuration_values AwsApprunnerService#code_configuration_values}
	// Experimental.
	CodeConfigurationValues *AwsApprunnerService_CodeConfigurationValuesProperty `field:"optional" json:"codeConfigurationValues" yaml:"codeConfigurationValues"`
}

