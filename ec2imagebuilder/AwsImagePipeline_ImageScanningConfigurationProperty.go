package ec2imagebuilder


// Experimental.
type AwsImagePipeline_ImageScanningConfigurationProperty struct {
	// ecr_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_pipeline#ecr_configuration AwsImagePipeline#ecr_configuration}
	// Experimental.
	EcrConfiguration *AwsImagePipeline_EcrConfigurationProperty `field:"optional" json:"ecrConfiguration" yaml:"ecrConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_pipeline#image_scanning_enabled AwsImagePipeline#image_scanning_enabled}.
	// Experimental.
	ImageScanningEnabled interface{} `field:"optional" json:"imageScanningEnabled" yaml:"imageScanningEnabled"`
}

