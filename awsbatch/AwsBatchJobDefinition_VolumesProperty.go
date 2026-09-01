package awsbatch


// Experimental.
type AwsBatchJobDefinition_VolumesProperty struct {
	// empty_dir block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#empty_dir AwsBatchJobDefinition#empty_dir}
	// Experimental.
	EmptyDir *AwsBatchJobDefinition_EmptyDirProperty `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// host_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#host_path AwsBatchJobDefinition#host_path}
	// Experimental.
	HostPath *AwsBatchJobDefinition_HostPathProperty `field:"optional" json:"hostPath" yaml:"hostPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#name AwsBatchJobDefinition#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#secret AwsBatchJobDefinition#secret}
	// Experimental.
	Secret *AwsBatchJobDefinition_SecretProperty `field:"optional" json:"secret" yaml:"secret"`
}

