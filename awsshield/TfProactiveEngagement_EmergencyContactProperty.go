package awsshield


// Experimental.
type TfProactiveEngagement_EmergencyContactProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/shield_proactive_engagement#email_address TfProactiveEngagement#email_address}.
	// Experimental.
	EmailAddress *string `field:"required" json:"emailAddress" yaml:"emailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/shield_proactive_engagement#contact_notes TfProactiveEngagement#contact_notes}.
	// Experimental.
	ContactNotes *string `field:"optional" json:"contactNotes" yaml:"contactNotes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/shield_proactive_engagement#phone_number TfProactiveEngagement#phone_number}.
	// Experimental.
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
}

