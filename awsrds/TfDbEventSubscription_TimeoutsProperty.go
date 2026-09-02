package awsrds


// Experimental.
type TfDbEventSubscription_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_event_subscription#create TfDbEventSubscription#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_event_subscription#delete TfDbEventSubscription#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_event_subscription#update TfDbEventSubscription#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

