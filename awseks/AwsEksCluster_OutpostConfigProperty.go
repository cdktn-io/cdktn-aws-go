package awseks


// Experimental.
type AwsEksCluster_OutpostConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#control_plane_instance_type AwsEksCluster#control_plane_instance_type}.
	// Experimental.
	ControlPlaneInstanceType *string `field:"required" json:"controlPlaneInstanceType" yaml:"controlPlaneInstanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#outpost_arns AwsEksCluster#outpost_arns}.
	// Experimental.
	OutpostArns *[]*string `field:"required" json:"outpostArns" yaml:"outpostArns"`
	// control_plane_placement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#control_plane_placement AwsEksCluster#control_plane_placement}
	// Experimental.
	ControlPlanePlacement *AwsEksCluster_ControlPlanePlacementProperty `field:"optional" json:"controlPlanePlacement" yaml:"controlPlanePlacement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#etcd_instance_type AwsEksCluster#etcd_instance_type}.
	// Experimental.
	EtcdInstanceType *string `field:"optional" json:"etcdInstanceType" yaml:"etcdInstanceType"`
	// etcd_placement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#etcd_placement AwsEksCluster#etcd_placement}
	// Experimental.
	EtcdPlacement *AwsEksCluster_EtcdPlacementProperty `field:"optional" json:"etcdPlacement" yaml:"etcdPlacement"`
}

