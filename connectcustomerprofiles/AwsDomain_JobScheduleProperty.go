package connectcustomerprofiles


// Experimental.
type AwsDomain_JobScheduleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#day_of_the_week AwsDomain#day_of_the_week}.
	// Experimental.
	DayOfTheWeek *string `field:"required" json:"dayOfTheWeek" yaml:"dayOfTheWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#time AwsDomain#time}.
	// Experimental.
	Time *string `field:"required" json:"time" yaml:"time"`
}

