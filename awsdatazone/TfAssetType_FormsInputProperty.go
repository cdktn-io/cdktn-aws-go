package awsdatazone


// Experimental.
type TfAssetType_FormsInputProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_asset_type#map_block_key TfAssetType#map_block_key}.
	// Experimental.
	MapBlockKey *string `field:"required" json:"mapBlockKey" yaml:"mapBlockKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_asset_type#type_identifier TfAssetType#type_identifier}.
	// Experimental.
	TypeIdentifier *string `field:"required" json:"typeIdentifier" yaml:"typeIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_asset_type#type_revision TfAssetType#type_revision}.
	// Experimental.
	TypeRevision *string `field:"required" json:"typeRevision" yaml:"typeRevision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_asset_type#required TfAssetType#required}.
	// Experimental.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}

