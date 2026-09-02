package awsec2


// Experimental.
type TfSpotFleetRequest_VcpuCountProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#max TfSpotFleetRequest#max}.
	// Experimental.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#min TfSpotFleetRequest#min}.
	// Experimental.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

