package awsemr


// Experimental.
type TfInstanceFleet_ConfigurationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_instance_fleet#classification TfInstanceFleet#classification}.
	// Experimental.
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_instance_fleet#properties TfInstanceFleet#properties}.
	// Experimental.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

