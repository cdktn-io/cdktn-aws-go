package awsec2imagebuilder


// Experimental.
type TfImage_ImageScanningConfigurationProperty struct {
	// ecr_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image#ecr_configuration TfImage#ecr_configuration}
	// Experimental.
	EcrConfiguration *TfImage_EcrConfigurationProperty `field:"optional" json:"ecrConfiguration" yaml:"ecrConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image#image_scanning_enabled TfImage#image_scanning_enabled}.
	// Experimental.
	ImageScanningEnabled interface{} `field:"optional" json:"imageScanningEnabled" yaml:"imageScanningEnabled"`
}

