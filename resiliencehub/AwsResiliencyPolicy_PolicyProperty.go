package resiliencehub


// Experimental.
type AwsResiliencyPolicy_PolicyProperty struct {
	// az block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehub_resiliency_policy#az AwsResiliencyPolicy#az}
	// Experimental.
	Az interface{} `field:"optional" json:"az" yaml:"az"`
	// hardware block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehub_resiliency_policy#hardware AwsResiliencyPolicy#hardware}
	// Experimental.
	Hardware interface{} `field:"optional" json:"hardware" yaml:"hardware"`
	// region block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehub_resiliency_policy#region AwsResiliencyPolicy#region}
	// Experimental.
	Region interface{} `field:"optional" json:"region" yaml:"region"`
	// software block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehub_resiliency_policy#software AwsResiliencyPolicy#software}
	// Experimental.
	SoftwareAttribute interface{} `field:"optional" json:"softwareAttribute" yaml:"softwareAttribute"`
}

