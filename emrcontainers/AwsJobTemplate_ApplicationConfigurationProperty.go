package emrcontainers


// Experimental.
type AwsJobTemplate_ApplicationConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#classification AwsJobTemplate#classification}.
	// Experimental.
	Classification *string `field:"required" json:"classification" yaml:"classification"`
	// configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#configurations AwsJobTemplate#configurations}
	// Experimental.
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#properties AwsJobTemplate#properties}.
	// Experimental.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

