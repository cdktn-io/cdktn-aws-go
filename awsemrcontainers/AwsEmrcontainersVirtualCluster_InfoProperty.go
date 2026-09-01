package awsemrcontainers


// Experimental.
type AwsEmrcontainersVirtualCluster_InfoProperty struct {
	// eks_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_virtual_cluster#eks_info AwsEmrcontainersVirtualCluster#eks_info}
	// Experimental.
	EksInfo *AwsEmrcontainersVirtualCluster_EksInfoProperty `field:"required" json:"eksInfo" yaml:"eksInfo"`
}

