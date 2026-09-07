package dataexchange


// Experimental.
type AwsEventAction_EncryptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dataexchange_event_action#kms_key_arn AwsEventAction#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dataexchange_event_action#type AwsEventAction#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

