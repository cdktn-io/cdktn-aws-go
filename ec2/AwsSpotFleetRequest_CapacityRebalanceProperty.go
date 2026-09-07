package ec2


// Experimental.
type AwsSpotFleetRequest_CapacityRebalanceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#replacement_strategy AwsSpotFleetRequest#replacement_strategy}.
	// Experimental.
	ReplacementStrategy *string `field:"optional" json:"replacementStrategy" yaml:"replacementStrategy"`
}

