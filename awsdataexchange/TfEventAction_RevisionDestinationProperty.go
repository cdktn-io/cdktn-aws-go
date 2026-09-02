package awsdataexchange


// Experimental.
type TfEventAction_RevisionDestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dataexchange_event_action#bucket TfEventAction#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dataexchange_event_action#key_pattern TfEventAction#key_pattern}.
	// Experimental.
	KeyPattern *string `field:"optional" json:"keyPattern" yaml:"keyPattern"`
}

