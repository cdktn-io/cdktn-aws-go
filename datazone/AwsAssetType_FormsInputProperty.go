package datazone


// Experimental.
type AwsAssetType_FormsInputProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_asset_type#map_block_key AwsAssetType#map_block_key}.
	// Experimental.
	MapBlockKey *string `field:"required" json:"mapBlockKey" yaml:"mapBlockKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_asset_type#type_identifier AwsAssetType#type_identifier}.
	// Experimental.
	TypeIdentifier *string `field:"required" json:"typeIdentifier" yaml:"typeIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_asset_type#type_revision AwsAssetType#type_revision}.
	// Experimental.
	TypeRevision *string `field:"required" json:"typeRevision" yaml:"typeRevision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_asset_type#required AwsAssetType#required}.
	// Experimental.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}

