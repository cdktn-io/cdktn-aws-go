package awsbatch


// Experimental.
type TfJobDefinition_EksPropertiesPodPropertiesContainersVolumeMountsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#mount_path TfJobDefinition#mount_path}.
	// Experimental.
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#name TfJobDefinition#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#read_only TfJobDefinition#read_only}.
	// Experimental.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

