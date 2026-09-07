package kendra


// Experimental.
type AwsIndex_CapacityUnitsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#query_capacity_units AwsIndex#query_capacity_units}.
	// Experimental.
	QueryCapacityUnits *float64 `field:"optional" json:"queryCapacityUnits" yaml:"queryCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#storage_capacity_units AwsIndex#storage_capacity_units}.
	// Experimental.
	StorageCapacityUnits *float64 `field:"optional" json:"storageCapacityUnits" yaml:"storageCapacityUnits"`
}

