package apprunner


// Experimental.
type AwsService_CodeConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#configuration_source AwsService#configuration_source}.
	// Experimental.
	ConfigurationSource *string `field:"required" json:"configurationSource" yaml:"configurationSource"`
	// code_configuration_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#code_configuration_values AwsService#code_configuration_values}
	// Experimental.
	CodeConfigurationValues *AwsService_CodeConfigurationValuesProperty `field:"optional" json:"codeConfigurationValues" yaml:"codeConfigurationValues"`
}

