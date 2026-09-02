package awseks


// Experimental.
type TfCluster_OutpostConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#control_plane_instance_type TfCluster#control_plane_instance_type}.
	// Experimental.
	ControlPlaneInstanceType *string `field:"required" json:"controlPlaneInstanceType" yaml:"controlPlaneInstanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#outpost_arns TfCluster#outpost_arns}.
	// Experimental.
	OutpostArns *[]*string `field:"required" json:"outpostArns" yaml:"outpostArns"`
	// control_plane_placement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#control_plane_placement TfCluster#control_plane_placement}
	// Experimental.
	ControlPlanePlacement *TfCluster_ControlPlanePlacementProperty `field:"optional" json:"controlPlanePlacement" yaml:"controlPlanePlacement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#etcd_instance_type TfCluster#etcd_instance_type}.
	// Experimental.
	EtcdInstanceType *string `field:"optional" json:"etcdInstanceType" yaml:"etcdInstanceType"`
	// etcd_placement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#etcd_placement TfCluster#etcd_placement}
	// Experimental.
	EtcdPlacement *TfCluster_EtcdPlacementProperty `field:"optional" json:"etcdPlacement" yaml:"etcdPlacement"`
}

