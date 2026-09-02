package awsrds


// Experimental.
type TfDbInstance_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#create TfDbInstance#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#delete TfDbInstance#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#update TfDbInstance#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

