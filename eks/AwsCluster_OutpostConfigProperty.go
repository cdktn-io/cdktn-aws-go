package eks


// Experimental.
type AwsCluster_OutpostConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#control_plane_instance_type AwsCluster#control_plane_instance_type}.
	// Experimental.
	ControlPlaneInstanceType *string `field:"required" json:"controlPlaneInstanceType" yaml:"controlPlaneInstanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#outpost_arns AwsCluster#outpost_arns}.
	// Experimental.
	OutpostArns *[]*string `field:"required" json:"outpostArns" yaml:"outpostArns"`
	// control_plane_placement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#control_plane_placement AwsCluster#control_plane_placement}
	// Experimental.
	ControlPlanePlacement *AwsCluster_ControlPlanePlacementProperty `field:"optional" json:"controlPlanePlacement" yaml:"controlPlanePlacement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#etcd_instance_type AwsCluster#etcd_instance_type}.
	// Experimental.
	EtcdInstanceType *string `field:"optional" json:"etcdInstanceType" yaml:"etcdInstanceType"`
	// etcd_placement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#etcd_placement AwsCluster#etcd_placement}
	// Experimental.
	EtcdPlacement *AwsCluster_EtcdPlacementProperty `field:"optional" json:"etcdPlacement" yaml:"etcdPlacement"`
}

