package awsssmcontacts


// Experimental.
type TfRotation_RecurrenceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_rotation#number_of_on_calls TfRotation#number_of_on_calls}.
	// Experimental.
	NumberOfOnCalls *float64 `field:"required" json:"numberOfOnCalls" yaml:"numberOfOnCalls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_rotation#recurrence_multiplier TfRotation#recurrence_multiplier}.
	// Experimental.
	RecurrenceMultiplier *float64 `field:"required" json:"recurrenceMultiplier" yaml:"recurrenceMultiplier"`
	// daily_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_rotation#daily_settings TfRotation#daily_settings}
	// Experimental.
	DailySettings interface{} `field:"optional" json:"dailySettings" yaml:"dailySettings"`
	// monthly_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_rotation#monthly_settings TfRotation#monthly_settings}
	// Experimental.
	MonthlySettings interface{} `field:"optional" json:"monthlySettings" yaml:"monthlySettings"`
	// shift_coverages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_rotation#shift_coverages TfRotation#shift_coverages}
	// Experimental.
	ShiftCoverages interface{} `field:"optional" json:"shiftCoverages" yaml:"shiftCoverages"`
	// weekly_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_rotation#weekly_settings TfRotation#weekly_settings}
	// Experimental.
	WeeklySettings interface{} `field:"optional" json:"weeklySettings" yaml:"weeklySettings"`
}

