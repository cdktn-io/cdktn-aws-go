package awsresiliencehub


// Experimental.
type TfResiliencyPolicy_RegionProperty struct {
	// Recovery Point Objective (RPO) as a Go duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehub_resiliency_policy#rpo TfResiliencyPolicy#rpo}
	// Experimental.
	Rpo *string `field:"optional" json:"rpo" yaml:"rpo"`
	// Recovery Time Objective (RTO) as a Go duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehub_resiliency_policy#rto TfResiliencyPolicy#rto}
	// Experimental.
	Rto *string `field:"optional" json:"rto" yaml:"rto"`
}

