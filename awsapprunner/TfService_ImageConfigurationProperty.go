package awsapprunner


// Experimental.
type TfService_ImageConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#port TfService#port}.
	// Experimental.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#runtime_environment_secrets TfService#runtime_environment_secrets}.
	// Experimental.
	RuntimeEnvironmentSecrets *map[string]*string `field:"optional" json:"runtimeEnvironmentSecrets" yaml:"runtimeEnvironmentSecrets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#runtime_environment_variables TfService#runtime_environment_variables}.
	// Experimental.
	RuntimeEnvironmentVariables *map[string]*string `field:"optional" json:"runtimeEnvironmentVariables" yaml:"runtimeEnvironmentVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#start_command TfService#start_command}.
	// Experimental.
	StartCommand *string `field:"optional" json:"startCommand" yaml:"startCommand"`
}

