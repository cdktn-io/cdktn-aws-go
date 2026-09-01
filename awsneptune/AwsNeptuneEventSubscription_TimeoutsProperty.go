package awsneptune


// Experimental.
type AwsNeptuneEventSubscription_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_event_subscription#create AwsNeptuneEventSubscription#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_event_subscription#delete AwsNeptuneEventSubscription#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_event_subscription#update AwsNeptuneEventSubscription#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

