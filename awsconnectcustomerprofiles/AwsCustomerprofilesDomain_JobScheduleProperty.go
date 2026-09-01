package awsconnectcustomerprofiles


// Experimental.
type AwsCustomerprofilesDomain_JobScheduleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#day_of_the_week AwsCustomerprofilesDomain#day_of_the_week}.
	// Experimental.
	DayOfTheWeek *string `field:"required" json:"dayOfTheWeek" yaml:"dayOfTheWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#time AwsCustomerprofilesDomain#time}.
	// Experimental.
	Time *string `field:"required" json:"time" yaml:"time"`
}

