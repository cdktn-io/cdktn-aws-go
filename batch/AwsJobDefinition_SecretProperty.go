package batch


// Experimental.
type AwsJobDefinition_SecretProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#secret_name AwsJobDefinition#secret_name}.
	// Experimental.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#optional AwsJobDefinition#optional}.
	// Experimental.
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

