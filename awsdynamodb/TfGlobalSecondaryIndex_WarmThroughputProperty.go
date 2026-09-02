package awsdynamodb


// Experimental.
type TfGlobalSecondaryIndex_WarmThroughputProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_global_secondary_index#read_units_per_second TfGlobalSecondaryIndex#read_units_per_second}.
	// Experimental.
	ReadUnitsPerSecond *float64 `field:"optional" json:"readUnitsPerSecond" yaml:"readUnitsPerSecond"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_global_secondary_index#write_units_per_second TfGlobalSecondaryIndex#write_units_per_second}.
	// Experimental.
	WriteUnitsPerSecond *float64 `field:"optional" json:"writeUnitsPerSecond" yaml:"writeUnitsPerSecond"`
}

