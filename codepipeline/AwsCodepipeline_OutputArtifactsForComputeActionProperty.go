package codepipeline


// Experimental.
type AwsCodepipeline_OutputArtifactsForComputeActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#name AwsCodepipeline#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#files AwsCodepipeline#files}.
	// Experimental.
	Files *[]*string `field:"optional" json:"files" yaml:"files"`
}

