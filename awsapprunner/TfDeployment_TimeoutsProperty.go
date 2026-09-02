package awsapprunner


// Experimental.
type TfDeployment_TimeoutsProperty struct {
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_deployment#create TfDeployment#create}
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

