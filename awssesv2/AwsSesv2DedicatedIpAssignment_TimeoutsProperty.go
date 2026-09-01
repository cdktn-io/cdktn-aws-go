package awssesv2


// Experimental.
type AwsSesv2DedicatedIpAssignment_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_dedicated_ip_assignment#create AwsSesv2DedicatedIpAssignment#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_dedicated_ip_assignment#delete AwsSesv2DedicatedIpAssignment#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

