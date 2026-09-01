package awsglue


// Experimental.
type DataAwsGlueScript_DagEdgeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/glue_script#source DataAwsGlueScript#source}.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/glue_script#target DataAwsGlueScript#target}.
	// Experimental.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/glue_script#target_parameter DataAwsGlueScript#target_parameter}.
	// Experimental.
	TargetParameter *string `field:"optional" json:"targetParameter" yaml:"targetParameter"`
}

