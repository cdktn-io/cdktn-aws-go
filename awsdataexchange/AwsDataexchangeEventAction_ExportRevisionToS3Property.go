package awsdataexchange


// Experimental.
type AwsDataexchangeEventAction_ExportRevisionToS3Property struct {
	// encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dataexchange_event_action#encryption AwsDataexchangeEventAction#encryption}
	// Experimental.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// revision_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dataexchange_event_action#revision_destination AwsDataexchangeEventAction#revision_destination}
	// Experimental.
	RevisionDestination interface{} `field:"optional" json:"revisionDestination" yaml:"revisionDestination"`
}

