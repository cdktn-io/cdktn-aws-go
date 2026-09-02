package awsbatch


// Experimental.
type TfJobDefinition_VolumesProperty struct {
	// empty_dir block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#empty_dir TfJobDefinition#empty_dir}
	// Experimental.
	EmptyDir *TfJobDefinition_EmptyDirProperty `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// host_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#host_path TfJobDefinition#host_path}
	// Experimental.
	HostPath *TfJobDefinition_HostPathProperty `field:"optional" json:"hostPath" yaml:"hostPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#name TfJobDefinition#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#secret TfJobDefinition#secret}
	// Experimental.
	Secret *TfJobDefinition_SecretProperty `field:"optional" json:"secret" yaml:"secret"`
}

