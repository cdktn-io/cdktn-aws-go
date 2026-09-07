package ec2imagebuilder


// Experimental.
type AwsImage_ImageScanningConfigurationProperty struct {
	// ecr_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image#ecr_configuration AwsImage#ecr_configuration}
	// Experimental.
	EcrConfiguration *AwsImage_EcrConfigurationProperty `field:"optional" json:"ecrConfiguration" yaml:"ecrConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image#image_scanning_enabled AwsImage#image_scanning_enabled}.
	// Experimental.
	ImageScanningEnabled interface{} `field:"optional" json:"imageScanningEnabled" yaml:"imageScanningEnabled"`
}

