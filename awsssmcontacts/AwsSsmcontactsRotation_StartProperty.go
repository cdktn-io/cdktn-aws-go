package awsssmcontacts


// Experimental.
type AwsSsmcontactsRotation_StartProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_rotation#hour_of_day AwsSsmcontactsRotation#hour_of_day}.
	// Experimental.
	HourOfDay *float64 `field:"required" json:"hourOfDay" yaml:"hourOfDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_rotation#minute_of_hour AwsSsmcontactsRotation#minute_of_hour}.
	// Experimental.
	MinuteOfHour *float64 `field:"required" json:"minuteOfHour" yaml:"minuteOfHour"`
}

