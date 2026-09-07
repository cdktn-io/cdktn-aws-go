package neptune


// Experimental.
type AwsEventSubscription_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_event_subscription#create AwsEventSubscription#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_event_subscription#delete AwsEventSubscription#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_event_subscription#update AwsEventSubscription#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

