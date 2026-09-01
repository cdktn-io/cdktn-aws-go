package awsdataexchange


// Experimental.
type AwsDataexchangeRevisionAssets_AssetImportAssetsFromS3AssetSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dataexchange_revision_assets#bucket AwsDataexchangeRevisionAssets#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dataexchange_revision_assets#key AwsDataexchangeRevisionAssets#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
}

