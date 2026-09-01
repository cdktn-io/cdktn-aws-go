package awsstoragegateway


// Experimental.
type AwsStoragegatewayGateway_MaintenanceStartTimeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#hour_of_day AwsStoragegatewayGateway#hour_of_day}.
	// Experimental.
	HourOfDay *float64 `field:"required" json:"hourOfDay" yaml:"hourOfDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#day_of_month AwsStoragegatewayGateway#day_of_month}.
	// Experimental.
	DayOfMonth *string `field:"optional" json:"dayOfMonth" yaml:"dayOfMonth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#day_of_week AwsStoragegatewayGateway#day_of_week}.
	// Experimental.
	DayOfWeek *string `field:"optional" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#minute_of_hour AwsStoragegatewayGateway#minute_of_hour}.
	// Experimental.
	MinuteOfHour *float64 `field:"optional" json:"minuteOfHour" yaml:"minuteOfHour"`
}

