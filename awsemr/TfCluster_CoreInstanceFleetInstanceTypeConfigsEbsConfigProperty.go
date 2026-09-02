package awsemr


// Experimental.
type TfCluster_CoreInstanceFleetInstanceTypeConfigsEbsConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#size TfCluster#size}.
	// Experimental.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#type TfCluster#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#iops TfCluster#iops}.
	// Experimental.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#volumes_per_instance TfCluster#volumes_per_instance}.
	// Experimental.
	VolumesPerInstance *float64 `field:"optional" json:"volumesPerInstance" yaml:"volumesPerInstance"`
}

