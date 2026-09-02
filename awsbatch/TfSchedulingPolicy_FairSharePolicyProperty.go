package awsbatch


// Experimental.
type TfSchedulingPolicy_FairSharePolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_scheduling_policy#compute_reservation TfSchedulingPolicy#compute_reservation}.
	// Experimental.
	ComputeReservation *float64 `field:"optional" json:"computeReservation" yaml:"computeReservation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_scheduling_policy#share_decay_seconds TfSchedulingPolicy#share_decay_seconds}.
	// Experimental.
	ShareDecaySeconds *float64 `field:"optional" json:"shareDecaySeconds" yaml:"shareDecaySeconds"`
	// share_distribution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_scheduling_policy#share_distribution TfSchedulingPolicy#share_distribution}
	// Experimental.
	ShareDistribution interface{} `field:"optional" json:"shareDistribution" yaml:"shareDistribution"`
}

