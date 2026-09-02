package awsglue


// Experimental.
type DataTfScript_DagEdgeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/glue_script#source DataTfScript#source}.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/glue_script#target DataTfScript#target}.
	// Experimental.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/glue_script#target_parameter DataTfScript#target_parameter}.
	// Experimental.
	TargetParameter *string `field:"optional" json:"targetParameter" yaml:"targetParameter"`
}

