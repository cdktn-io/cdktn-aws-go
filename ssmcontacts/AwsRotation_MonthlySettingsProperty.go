package ssmcontacts


// Experimental.
type AwsRotation_MonthlySettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_rotation#day_of_month AwsRotation#day_of_month}.
	// Experimental.
	DayOfMonth *float64 `field:"required" json:"dayOfMonth" yaml:"dayOfMonth"`
	// hand_off_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_rotation#hand_off_time AwsRotation#hand_off_time}
	// Experimental.
	HandOffTime interface{} `field:"optional" json:"handOffTime" yaml:"handOffTime"`
}

