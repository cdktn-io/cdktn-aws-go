package keyspaces


// Experimental.
type AwsTable_CapacitySpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#read_capacity_units AwsTable#read_capacity_units}.
	// Experimental.
	ReadCapacityUnits *float64 `field:"optional" json:"readCapacityUnits" yaml:"readCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#throughput_mode AwsTable#throughput_mode}.
	// Experimental.
	ThroughputMode *string `field:"optional" json:"throughputMode" yaml:"throughputMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#write_capacity_units AwsTable#write_capacity_units}.
	// Experimental.
	WriteCapacityUnits *float64 `field:"optional" json:"writeCapacityUnits" yaml:"writeCapacityUnits"`
}

