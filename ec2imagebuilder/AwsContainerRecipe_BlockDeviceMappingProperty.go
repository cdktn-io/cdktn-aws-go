package ec2imagebuilder


// Experimental.
type AwsContainerRecipe_BlockDeviceMappingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#device_name AwsContainerRecipe#device_name}.
	// Experimental.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// ebs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#ebs AwsContainerRecipe#ebs}
	// Experimental.
	Ebs *AwsContainerRecipe_EbsProperty `field:"optional" json:"ebs" yaml:"ebs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#no_device AwsContainerRecipe#no_device}.
	// Experimental.
	NoDevice interface{} `field:"optional" json:"noDevice" yaml:"noDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#virtual_name AwsContainerRecipe#virtual_name}.
	// Experimental.
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}

