package awsconnect


// Experimental.
type AwsConnectHoursOfOperation_StartTimeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_hours_of_operation#hours AwsConnectHoursOfOperation#hours}.
	// Experimental.
	Hours *float64 `field:"required" json:"hours" yaml:"hours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_hours_of_operation#minutes AwsConnectHoursOfOperation#minutes}.
	// Experimental.
	Minutes *float64 `field:"required" json:"minutes" yaml:"minutes"`
}

