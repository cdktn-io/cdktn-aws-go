package awsbatch


// Experimental.
type AwsBatchSchedulingPolicy_ShareDistributionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_scheduling_policy#share_identifier AwsBatchSchedulingPolicy#share_identifier}.
	// Experimental.
	ShareIdentifier *string `field:"required" json:"shareIdentifier" yaml:"shareIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_scheduling_policy#weight_factor AwsBatchSchedulingPolicy#weight_factor}.
	// Experimental.
	WeightFactor *float64 `field:"optional" json:"weightFactor" yaml:"weightFactor"`
}

