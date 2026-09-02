package awsmsk


// Experimental.
type TfCluster_EncryptionInfoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#encryption_at_rest_kms_key_arn TfCluster#encryption_at_rest_kms_key_arn}.
	// Experimental.
	EncryptionAtRestKmsKeyArn *string `field:"optional" json:"encryptionAtRestKmsKeyArn" yaml:"encryptionAtRestKmsKeyArn"`
	// encryption_in_transit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#encryption_in_transit TfCluster#encryption_in_transit}
	// Experimental.
	EncryptionInTransit *TfCluster_EncryptionInTransitProperty `field:"optional" json:"encryptionInTransit" yaml:"encryptionInTransit"`
}

