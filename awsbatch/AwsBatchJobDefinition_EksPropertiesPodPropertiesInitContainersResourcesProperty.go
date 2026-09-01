package awsbatch


// Experimental.
type AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersResourcesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#limits AwsBatchJobDefinition#limits}.
	// Experimental.
	Limits *map[string]*string `field:"optional" json:"limits" yaml:"limits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#requests AwsBatchJobDefinition#requests}.
	// Experimental.
	Requests *map[string]*string `field:"optional" json:"requests" yaml:"requests"`
}

