package mq


// Experimental.
type AwsBroker_MaintenanceWindowStartTimeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#day_of_week AwsBroker#day_of_week}.
	// Experimental.
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#time_of_day AwsBroker#time_of_day}.
	// Experimental.
	TimeOfDay *string `field:"required" json:"timeOfDay" yaml:"timeOfDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#time_zone AwsBroker#time_zone}.
	// Experimental.
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
}

