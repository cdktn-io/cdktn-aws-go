package awsemr


// Experimental.
type TfCluster_CoreInstanceFleetInstanceTypeConfigsConfigurationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#classification TfCluster#classification}.
	// Experimental.
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#properties TfCluster#properties}.
	// Experimental.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

