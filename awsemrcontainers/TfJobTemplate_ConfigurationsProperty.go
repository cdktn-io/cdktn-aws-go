package awsemrcontainers


// Experimental.
type TfJobTemplate_ConfigurationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#classification TfJobTemplate#classification}.
	// Experimental.
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#properties TfJobTemplate#properties}.
	// Experimental.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

