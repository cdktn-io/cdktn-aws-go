package awselementalmediaconvert


// Experimental.
type AwsMediaConvertQueue_ReservationPlanSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/media_convert_queue#commitment AwsMediaConvertQueue#commitment}.
	// Experimental.
	Commitment *string `field:"required" json:"commitment" yaml:"commitment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/media_convert_queue#renewal_type AwsMediaConvertQueue#renewal_type}.
	// Experimental.
	RenewalType *string `field:"required" json:"renewalType" yaml:"renewalType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/media_convert_queue#reserved_slots AwsMediaConvertQueue#reserved_slots}.
	// Experimental.
	ReservedSlots *float64 `field:"required" json:"reservedSlots" yaml:"reservedSlots"`
}

