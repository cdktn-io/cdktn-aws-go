package awsapprunner


// Experimental.
type AwsApprunnerService_SourceConfigurationProperty struct {
	// authentication_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#authentication_configuration AwsApprunnerService#authentication_configuration}
	// Experimental.
	AuthenticationConfiguration *AwsApprunnerService_AuthenticationConfigurationProperty `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#auto_deployments_enabled AwsApprunnerService#auto_deployments_enabled}.
	// Experimental.
	AutoDeploymentsEnabled interface{} `field:"optional" json:"autoDeploymentsEnabled" yaml:"autoDeploymentsEnabled"`
	// code_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#code_repository AwsApprunnerService#code_repository}
	// Experimental.
	CodeRepository *AwsApprunnerService_CodeRepositoryProperty `field:"optional" json:"codeRepository" yaml:"codeRepository"`
	// image_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#image_repository AwsApprunnerService#image_repository}
	// Experimental.
	ImageRepository *AwsApprunnerService_ImageRepositoryProperty `field:"optional" json:"imageRepository" yaml:"imageRepository"`
}

