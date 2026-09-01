package awsdataexchange


// Experimental.
type AwsDataexchangeEventAction_EncryptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dataexchange_event_action#kms_key_arn AwsDataexchangeEventAction#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dataexchange_event_action#type AwsDataexchangeEventAction#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

