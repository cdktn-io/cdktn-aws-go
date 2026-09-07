package batch


// Experimental.
type AwsJobDefinition_EksPropertiesPodPropertiesInitContainersResourcesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#limits AwsJobDefinition#limits}.
	// Experimental.
	Limits *map[string]*string `field:"optional" json:"limits" yaml:"limits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#requests AwsJobDefinition#requests}.
	// Experimental.
	Requests *map[string]*string `field:"optional" json:"requests" yaml:"requests"`
}

