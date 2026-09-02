package awsdocumentdb


// Experimental.
type TfEventSubscription_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdb_event_subscription#create TfEventSubscription#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdb_event_subscription#delete TfEventSubscription#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdb_event_subscription#update TfEventSubscription#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

