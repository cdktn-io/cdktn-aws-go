package glue


// Experimental.
type AwsJob_CommandProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#script_location AwsJob#script_location}.
	// Experimental.
	ScriptLocation *string `field:"required" json:"scriptLocation" yaml:"scriptLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#name AwsJob#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#python_version AwsJob#python_version}.
	// Experimental.
	PythonVersion *string `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#runtime AwsJob#runtime}.
	// Experimental.
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
}

