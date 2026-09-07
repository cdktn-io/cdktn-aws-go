package resiliencehub


// Experimental.
type AwsResiliencyPolicy_SoftwareProperty struct {
	// Recovery Point Objective (RPO) as a Go duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehub_resiliency_policy#rpo AwsResiliencyPolicy#rpo}
	// Experimental.
	Rpo *string `field:"required" json:"rpo" yaml:"rpo"`
	// Recovery Time Objective (RTO) as a Go duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehub_resiliency_policy#rto AwsResiliencyPolicy#rto}
	// Experimental.
	Rto *string `field:"required" json:"rto" yaml:"rto"`
}

