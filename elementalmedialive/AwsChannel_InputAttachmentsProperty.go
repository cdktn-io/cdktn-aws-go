package elementalmedialive


// Experimental.
type AwsChannel_InputAttachmentsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_attachment_name AwsChannel#input_attachment_name}.
	// Experimental.
	InputAttachmentName *string `field:"required" json:"inputAttachmentName" yaml:"inputAttachmentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_id AwsChannel#input_id}.
	// Experimental.
	InputId *string `field:"required" json:"inputId" yaml:"inputId"`
	// automatic_input_failover_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#automatic_input_failover_settings AwsChannel#automatic_input_failover_settings}
	// Experimental.
	AutomaticInputFailoverSettings *AwsChannel_AutomaticInputFailoverSettingsProperty `field:"optional" json:"automaticInputFailoverSettings" yaml:"automaticInputFailoverSettings"`
	// input_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_settings AwsChannel#input_settings}
	// Experimental.
	InputSettings *AwsChannel_InputSettingsProperty `field:"optional" json:"inputSettings" yaml:"inputSettings"`
}

