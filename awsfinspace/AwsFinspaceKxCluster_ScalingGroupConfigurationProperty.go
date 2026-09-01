package awsfinspace


// Experimental.
type AwsFinspaceKxCluster_ScalingGroupConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#memory_reservation AwsFinspaceKxCluster#memory_reservation}.
	// Experimental.
	MemoryReservation *float64 `field:"required" json:"memoryReservation" yaml:"memoryReservation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#node_count AwsFinspaceKxCluster#node_count}.
	// Experimental.
	NodeCount *float64 `field:"required" json:"nodeCount" yaml:"nodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#scaling_group_name AwsFinspaceKxCluster#scaling_group_name}.
	// Experimental.
	ScalingGroupName *string `field:"required" json:"scalingGroupName" yaml:"scalingGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#cpu AwsFinspaceKxCluster#cpu}.
	// Experimental.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#memory_limit AwsFinspaceKxCluster#memory_limit}.
	// Experimental.
	MemoryLimit *float64 `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
}

