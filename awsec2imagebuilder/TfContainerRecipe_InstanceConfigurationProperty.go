package awsec2imagebuilder


// Experimental.
type TfContainerRecipe_InstanceConfigurationProperty struct {
	// block_device_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#block_device_mapping TfContainerRecipe#block_device_mapping}
	// Experimental.
	BlockDeviceMapping interface{} `field:"optional" json:"blockDeviceMapping" yaml:"blockDeviceMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#image TfContainerRecipe#image}.
	// Experimental.
	Image *string `field:"optional" json:"image" yaml:"image"`
}

