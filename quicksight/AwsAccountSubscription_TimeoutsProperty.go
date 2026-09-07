package quicksight


// Experimental.
type AwsAccountSubscription_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#create AwsAccountSubscription#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#delete AwsAccountSubscription#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription#read AwsAccountSubscription#read}.
	// Experimental.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

