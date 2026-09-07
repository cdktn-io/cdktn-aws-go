package ssm


// Experimental.
type AwsPatchBaseline_GlobalFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#key AwsPatchBaseline#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#values AwsPatchBaseline#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

