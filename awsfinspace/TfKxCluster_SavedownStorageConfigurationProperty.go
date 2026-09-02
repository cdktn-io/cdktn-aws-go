package awsfinspace


// Experimental.
type TfKxCluster_SavedownStorageConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#size TfKxCluster#size}.
	// Experimental.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#type TfKxCluster#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#volume_name TfKxCluster#volume_name}.
	// Experimental.
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

