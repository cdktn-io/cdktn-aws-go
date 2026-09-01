package awsconnect


// Experimental.
type AwsConnectHoursOfOperation_ConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_hours_of_operation#day AwsConnectHoursOfOperation#day}.
	// Experimental.
	Day *string `field:"required" json:"day" yaml:"day"`
	// end_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_hours_of_operation#end_time AwsConnectHoursOfOperation#end_time}
	// Experimental.
	EndTime *AwsConnectHoursOfOperation_EndTimeProperty `field:"required" json:"endTime" yaml:"endTime"`
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_hours_of_operation#start_time AwsConnectHoursOfOperation#start_time}
	// Experimental.
	StartTime *AwsConnectHoursOfOperation_StartTimeProperty `field:"required" json:"startTime" yaml:"startTime"`
}

