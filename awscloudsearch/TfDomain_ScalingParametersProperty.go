package awscloudsearch


// Experimental.
type TfDomain_ScalingParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudsearch_domain#desired_instance_type TfDomain#desired_instance_type}.
	// Experimental.
	DesiredInstanceType *string `field:"optional" json:"desiredInstanceType" yaml:"desiredInstanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudsearch_domain#desired_partition_count TfDomain#desired_partition_count}.
	// Experimental.
	DesiredPartitionCount *float64 `field:"optional" json:"desiredPartitionCount" yaml:"desiredPartitionCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudsearch_domain#desired_replication_count TfDomain#desired_replication_count}.
	// Experimental.
	DesiredReplicationCount *float64 `field:"optional" json:"desiredReplicationCount" yaml:"desiredReplicationCount"`
}

