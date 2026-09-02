package awsamazonqbusiness


// Experimental.
type TfApplication_AttachmentsConfigurationProperty struct {
	// Status information about whether file upload functionality is activated or deactivated for your end user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/qbusiness_application#attachments_control_mode TfApplication#attachments_control_mode}
	// Experimental.
	AttachmentsControlMode *string `field:"required" json:"attachmentsControlMode" yaml:"attachmentsControlMode"`
}

