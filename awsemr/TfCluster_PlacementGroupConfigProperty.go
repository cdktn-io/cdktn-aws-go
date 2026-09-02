package awsemr


// Experimental.
type TfCluster_PlacementGroupConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#instance_role TfCluster#instance_role}.
	// Experimental.
	InstanceRole *string `field:"optional" json:"instanceRole" yaml:"instanceRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#placement_strategy TfCluster#placement_strategy}.
	// Experimental.
	PlacementStrategy *string `field:"optional" json:"placementStrategy" yaml:"placementStrategy"`
}

