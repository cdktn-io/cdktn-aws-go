package dataexchange


// Experimental.
type AwsRevisionAssets_AssetCreateS3DataAccessFromS3BucketAssetSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dataexchange_revision_assets#bucket AwsRevisionAssets#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dataexchange_revision_assets#key_prefixes AwsRevisionAssets#key_prefixes}.
	// Experimental.
	KeyPrefixes *[]*string `field:"optional" json:"keyPrefixes" yaml:"keyPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dataexchange_revision_assets#keys AwsRevisionAssets#keys}.
	// Experimental.
	Keys *[]*string `field:"optional" json:"keys" yaml:"keys"`
	// kms_keys_to_grant block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dataexchange_revision_assets#kms_keys_to_grant AwsRevisionAssets#kms_keys_to_grant}
	// Experimental.
	KmsKeysToGrant interface{} `field:"optional" json:"kmsKeysToGrant" yaml:"kmsKeysToGrant"`
}

