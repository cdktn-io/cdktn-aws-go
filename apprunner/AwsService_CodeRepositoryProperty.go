package apprunner


// Experimental.
type AwsService_CodeRepositoryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#repository_url AwsService#repository_url}.
	// Experimental.
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
	// source_code_version block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#source_code_version AwsService#source_code_version}
	// Experimental.
	SourceCodeVersion *AwsService_SourceCodeVersionProperty `field:"required" json:"sourceCodeVersion" yaml:"sourceCodeVersion"`
	// code_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#code_configuration AwsService#code_configuration}
	// Experimental.
	CodeConfiguration *AwsService_CodeConfigurationProperty `field:"optional" json:"codeConfiguration" yaml:"codeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#source_directory AwsService#source_directory}.
	// Experimental.
	SourceDirectory *string `field:"optional" json:"sourceDirectory" yaml:"sourceDirectory"`
}

