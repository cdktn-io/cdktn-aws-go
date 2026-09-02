package awsconnectcustomerprofiles


// Experimental.
type TfDomain_JobScheduleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#day_of_the_week TfDomain#day_of_the_week}.
	// Experimental.
	DayOfTheWeek *string `field:"required" json:"dayOfTheWeek" yaml:"dayOfTheWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#time TfDomain#time}.
	// Experimental.
	Time *string `field:"required" json:"time" yaml:"time"`
}

