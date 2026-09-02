package awsbatch


// Experimental.
type TfSchedulingPolicy_ShareDistributionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_scheduling_policy#share_identifier TfSchedulingPolicy#share_identifier}.
	// Experimental.
	ShareIdentifier *string `field:"required" json:"shareIdentifier" yaml:"shareIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_scheduling_policy#weight_factor TfSchedulingPolicy#weight_factor}.
	// Experimental.
	WeightFactor *float64 `field:"optional" json:"weightFactor" yaml:"weightFactor"`
}

