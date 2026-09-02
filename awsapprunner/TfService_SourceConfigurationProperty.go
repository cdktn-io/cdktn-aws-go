package awsapprunner


// Experimental.
type TfService_SourceConfigurationProperty struct {
	// authentication_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#authentication_configuration TfService#authentication_configuration}
	// Experimental.
	AuthenticationConfiguration *TfService_AuthenticationConfigurationProperty `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#auto_deployments_enabled TfService#auto_deployments_enabled}.
	// Experimental.
	AutoDeploymentsEnabled interface{} `field:"optional" json:"autoDeploymentsEnabled" yaml:"autoDeploymentsEnabled"`
	// code_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#code_repository TfService#code_repository}
	// Experimental.
	CodeRepository *TfService_CodeRepositoryProperty `field:"optional" json:"codeRepository" yaml:"codeRepository"`
	// image_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#image_repository TfService#image_repository}
	// Experimental.
	ImageRepository *TfService_ImageRepositoryProperty `field:"optional" json:"imageRepository" yaml:"imageRepository"`
}

