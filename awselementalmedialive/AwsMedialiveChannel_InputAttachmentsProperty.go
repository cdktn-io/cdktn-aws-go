package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_InputAttachmentsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_attachment_name AwsMedialiveChannel#input_attachment_name}.
	// Experimental.
	InputAttachmentName *string `field:"required" json:"inputAttachmentName" yaml:"inputAttachmentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_id AwsMedialiveChannel#input_id}.
	// Experimental.
	InputId *string `field:"required" json:"inputId" yaml:"inputId"`
	// automatic_input_failover_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#automatic_input_failover_settings AwsMedialiveChannel#automatic_input_failover_settings}
	// Experimental.
	AutomaticInputFailoverSettings *AwsMedialiveChannel_AutomaticInputFailoverSettingsProperty `field:"optional" json:"automaticInputFailoverSettings" yaml:"automaticInputFailoverSettings"`
	// input_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_settings AwsMedialiveChannel#input_settings}
	// Experimental.
	InputSettings *AwsMedialiveChannel_InputSettingsProperty `field:"optional" json:"inputSettings" yaml:"inputSettings"`
}

