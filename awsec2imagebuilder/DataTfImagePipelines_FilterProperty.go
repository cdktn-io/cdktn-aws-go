package awsec2imagebuilder


// Experimental.
type DataTfImagePipelines_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/imagebuilder_image_pipelines#name DataTfImagePipelines#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/imagebuilder_image_pipelines#values DataTfImagePipelines#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

