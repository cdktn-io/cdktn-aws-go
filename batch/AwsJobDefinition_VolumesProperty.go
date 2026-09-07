package batch


// Experimental.
type AwsJobDefinition_VolumesProperty struct {
	// empty_dir block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#empty_dir AwsJobDefinition#empty_dir}
	// Experimental.
	EmptyDir *AwsJobDefinition_EmptyDirProperty `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// host_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#host_path AwsJobDefinition#host_path}
	// Experimental.
	HostPath *AwsJobDefinition_HostPathProperty `field:"optional" json:"hostPath" yaml:"hostPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#name AwsJobDefinition#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#secret AwsJobDefinition#secret}
	// Experimental.
	Secret *AwsJobDefinition_SecretProperty `field:"optional" json:"secret" yaml:"secret"`
}

