package awssagemakerai


// Experimental.
type AwsSagemakerFeatureGroup_ThroughputConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#provisioned_read_capacity_units AwsSagemakerFeatureGroup#provisioned_read_capacity_units}.
	// Experimental.
	ProvisionedReadCapacityUnits *float64 `field:"optional" json:"provisionedReadCapacityUnits" yaml:"provisionedReadCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#provisioned_write_capacity_units AwsSagemakerFeatureGroup#provisioned_write_capacity_units}.
	// Experimental.
	ProvisionedWriteCapacityUnits *float64 `field:"optional" json:"provisionedWriteCapacityUnits" yaml:"provisionedWriteCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#throughput_mode AwsSagemakerFeatureGroup#throughput_mode}.
	// Experimental.
	ThroughputMode *string `field:"optional" json:"throughputMode" yaml:"throughputMode"`
}

