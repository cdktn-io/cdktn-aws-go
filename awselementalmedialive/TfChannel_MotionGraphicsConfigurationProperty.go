package awselementalmedialive


// Experimental.
type TfChannel_MotionGraphicsConfigurationProperty struct {
	// motion_graphics_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#motion_graphics_settings TfChannel#motion_graphics_settings}
	// Experimental.
	MotionGraphicsSettings *TfChannel_MotionGraphicsSettingsProperty `field:"required" json:"motionGraphicsSettings" yaml:"motionGraphicsSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#motion_graphics_insertion TfChannel#motion_graphics_insertion}.
	// Experimental.
	MotionGraphicsInsertion *string `field:"optional" json:"motionGraphicsInsertion" yaml:"motionGraphicsInsertion"`
}

