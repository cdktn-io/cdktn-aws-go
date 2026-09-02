package awselb


// Experimental.
type TfLb_MinimumLoadBalancerCapacityProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#capacity_units TfLb#capacity_units}.
	// Experimental.
	CapacityUnits *float64 `field:"required" json:"capacityUnits" yaml:"capacityUnits"`
}

