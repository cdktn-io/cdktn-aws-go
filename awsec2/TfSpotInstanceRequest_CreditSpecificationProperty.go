package awsec2


// Experimental.
type TfSpotInstanceRequest_CreditSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#cpu_credits TfSpotInstanceRequest#cpu_credits}.
	// Experimental.
	CpuCredits *string `field:"optional" json:"cpuCredits" yaml:"cpuCredits"`
}

