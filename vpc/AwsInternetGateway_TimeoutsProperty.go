package vpc


// Experimental.
type AwsInternetGateway_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/internet_gateway#create AwsInternetGateway#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/internet_gateway#delete AwsInternetGateway#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/internet_gateway#update AwsInternetGateway#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

