package connect


// Experimental.
type AwsHoursOfOperation_ConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_hours_of_operation#day AwsHoursOfOperation#day}.
	// Experimental.
	Day *string `field:"required" json:"day" yaml:"day"`
	// end_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_hours_of_operation#end_time AwsHoursOfOperation#end_time}
	// Experimental.
	EndTime *AwsHoursOfOperation_EndTimeProperty `field:"required" json:"endTime" yaml:"endTime"`
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_hours_of_operation#start_time AwsHoursOfOperation#start_time}
	// Experimental.
	StartTime *AwsHoursOfOperation_StartTimeProperty `field:"required" json:"startTime" yaml:"startTime"`
}

