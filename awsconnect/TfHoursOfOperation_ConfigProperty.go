package awsconnect


// Experimental.
type TfHoursOfOperation_ConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_hours_of_operation#day TfHoursOfOperation#day}.
	// Experimental.
	Day *string `field:"required" json:"day" yaml:"day"`
	// end_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_hours_of_operation#end_time TfHoursOfOperation#end_time}
	// Experimental.
	EndTime *TfHoursOfOperation_EndTimeProperty `field:"required" json:"endTime" yaml:"endTime"`
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_hours_of_operation#start_time TfHoursOfOperation#start_time}
	// Experimental.
	StartTime *TfHoursOfOperation_StartTimeProperty `field:"required" json:"startTime" yaml:"startTime"`
}

