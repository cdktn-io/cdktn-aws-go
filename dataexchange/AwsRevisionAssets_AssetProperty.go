package dataexchange


// Experimental.
type AwsRevisionAssets_AssetProperty struct {
	// create_s3_data_access_from_s3_bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dataexchange_revision_assets#create_s3_data_access_from_s3_bucket AwsRevisionAssets#create_s3_data_access_from_s3_bucket}
	// Experimental.
	CreateS3DataAccessFromS3Bucket interface{} `field:"optional" json:"createS3DataAccessFromS3Bucket" yaml:"createS3DataAccessFromS3Bucket"`
	// import_assets_from_s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dataexchange_revision_assets#import_assets_from_s3 AwsRevisionAssets#import_assets_from_s3}
	// Experimental.
	ImportAssetsFromS3 interface{} `field:"optional" json:"importAssetsFromS3" yaml:"importAssetsFromS3"`
	// import_assets_from_signed_url block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dataexchange_revision_assets#import_assets_from_signed_url AwsRevisionAssets#import_assets_from_signed_url}
	// Experimental.
	ImportAssetsFromSignedUrl interface{} `field:"optional" json:"importAssetsFromSignedUrl" yaml:"importAssetsFromSignedUrl"`
}

