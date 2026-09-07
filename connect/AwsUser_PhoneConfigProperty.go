package connect


// Experimental.
type AwsUser_PhoneConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user#phone_type AwsUser#phone_type}.
	// Experimental.
	PhoneType *string `field:"required" json:"phoneType" yaml:"phoneType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user#after_contact_work_time_limit AwsUser#after_contact_work_time_limit}.
	// Experimental.
	AfterContactWorkTimeLimit *float64 `field:"optional" json:"afterContactWorkTimeLimit" yaml:"afterContactWorkTimeLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user#auto_accept AwsUser#auto_accept}.
	// Experimental.
	AutoAccept interface{} `field:"optional" json:"autoAccept" yaml:"autoAccept"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user#desk_phone_number AwsUser#desk_phone_number}.
	// Experimental.
	DeskPhoneNumber *string `field:"optional" json:"deskPhoneNumber" yaml:"deskPhoneNumber"`
}

