package awskeyspaces


// Experimental.
type TfTable_CapacitySpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#read_capacity_units TfTable#read_capacity_units}.
	// Experimental.
	ReadCapacityUnits *float64 `field:"optional" json:"readCapacityUnits" yaml:"readCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#throughput_mode TfTable#throughput_mode}.
	// Experimental.
	ThroughputMode *string `field:"optional" json:"throughputMode" yaml:"throughputMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#write_capacity_units TfTable#write_capacity_units}.
	// Experimental.
	WriteCapacityUnits *float64 `field:"optional" json:"writeCapacityUnits" yaml:"writeCapacityUnits"`
}

