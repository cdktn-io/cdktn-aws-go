package awsfinspace


// Experimental.
type TfKxDataview_SegmentConfigurationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_dataview#db_paths TfKxDataview#db_paths}.
	// Experimental.
	DbPaths *[]*string `field:"required" json:"dbPaths" yaml:"dbPaths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_dataview#volume_name TfKxDataview#volume_name}.
	// Experimental.
	VolumeName *string `field:"required" json:"volumeName" yaml:"volumeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_dataview#on_demand TfKxDataview#on_demand}.
	// Experimental.
	OnDemand interface{} `field:"optional" json:"onDemand" yaml:"onDemand"`
}

