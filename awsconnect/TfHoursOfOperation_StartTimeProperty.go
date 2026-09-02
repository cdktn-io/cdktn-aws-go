package awsconnect


// Experimental.
type TfHoursOfOperation_StartTimeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_hours_of_operation#hours TfHoursOfOperation#hours}.
	// Experimental.
	Hours *float64 `field:"required" json:"hours" yaml:"hours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_hours_of_operation#minutes TfHoursOfOperation#minutes}.
	// Experimental.
	Minutes *float64 `field:"required" json:"minutes" yaml:"minutes"`
}

