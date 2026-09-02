package awsbatch


// Experimental.
type TfJobDefinition_EksPropertiesPodPropertiesContainersResourcesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#limits TfJobDefinition#limits}.
	// Experimental.
	Limits *map[string]*string `field:"optional" json:"limits" yaml:"limits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#requests TfJobDefinition#requests}.
	// Experimental.
	Requests *map[string]*string `field:"optional" json:"requests" yaml:"requests"`
}

