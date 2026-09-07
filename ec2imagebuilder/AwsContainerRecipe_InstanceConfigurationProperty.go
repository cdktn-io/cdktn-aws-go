package ec2imagebuilder


// Experimental.
type AwsContainerRecipe_InstanceConfigurationProperty struct {
	// block_device_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#block_device_mapping AwsContainerRecipe#block_device_mapping}
	// Experimental.
	BlockDeviceMapping interface{} `field:"optional" json:"blockDeviceMapping" yaml:"blockDeviceMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#image AwsContainerRecipe#image}.
	// Experimental.
	Image *string `field:"optional" json:"image" yaml:"image"`
}

