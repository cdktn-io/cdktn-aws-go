package awsec2


// Experimental.
type AwsSpotFleetRequest_NetworkInterfaceCountProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#max AwsSpotFleetRequest#max}.
	// Experimental.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#min AwsSpotFleetRequest#min}.
	// Experimental.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

