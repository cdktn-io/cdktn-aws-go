package awsapprunner


// Experimental.
type TfService_ImageRepositoryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#image_identifier TfService#image_identifier}.
	// Experimental.
	ImageIdentifier *string `field:"required" json:"imageIdentifier" yaml:"imageIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#image_repository_type TfService#image_repository_type}.
	// Experimental.
	ImageRepositoryType *string `field:"required" json:"imageRepositoryType" yaml:"imageRepositoryType"`
	// image_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#image_configuration TfService#image_configuration}
	// Experimental.
	ImageConfiguration *TfService_ImageConfigurationProperty `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
}

