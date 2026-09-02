package awsglue


// Experimental.
type DataTfScript_ArgsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/glue_script#name DataTfScript#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/glue_script#value DataTfScript#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/glue_script#param DataTfScript#param}.
	// Experimental.
	Param interface{} `field:"optional" json:"param" yaml:"param"`
}

