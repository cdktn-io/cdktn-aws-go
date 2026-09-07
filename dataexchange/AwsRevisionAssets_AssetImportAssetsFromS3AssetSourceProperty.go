package dataexchange


// Experimental.
type AwsRevisionAssets_AssetImportAssetsFromS3AssetSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dataexchange_revision_assets#bucket AwsRevisionAssets#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dataexchange_revision_assets#key AwsRevisionAssets#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
}

