package dynamodb


// Experimental.
type AwsGlobalSecondaryIndex_ProvisionedThroughputProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_global_secondary_index#read_capacity_units AwsGlobalSecondaryIndex#read_capacity_units}.
	// Experimental.
	ReadCapacityUnits *float64 `field:"optional" json:"readCapacityUnits" yaml:"readCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_global_secondary_index#write_capacity_units AwsGlobalSecondaryIndex#write_capacity_units}.
	// Experimental.
	WriteCapacityUnits *float64 `field:"optional" json:"writeCapacityUnits" yaml:"writeCapacityUnits"`
}

