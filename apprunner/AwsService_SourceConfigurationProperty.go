package apprunner


// Experimental.
type AwsService_SourceConfigurationProperty struct {
	// authentication_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#authentication_configuration AwsService#authentication_configuration}
	// Experimental.
	AuthenticationConfiguration *AwsService_AuthenticationConfigurationProperty `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#auto_deployments_enabled AwsService#auto_deployments_enabled}.
	// Experimental.
	AutoDeploymentsEnabled interface{} `field:"optional" json:"autoDeploymentsEnabled" yaml:"autoDeploymentsEnabled"`
	// code_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#code_repository AwsService#code_repository}
	// Experimental.
	CodeRepository *AwsService_CodeRepositoryProperty `field:"optional" json:"codeRepository" yaml:"codeRepository"`
	// image_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#image_repository AwsService#image_repository}
	// Experimental.
	ImageRepository *AwsService_ImageRepositoryProperty `field:"optional" json:"imageRepository" yaml:"imageRepository"`
}

