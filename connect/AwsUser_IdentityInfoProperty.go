package connect


// Experimental.
type AwsUser_IdentityInfoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user#email AwsUser#email}.
	// Experimental.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user#first_name AwsUser#first_name}.
	// Experimental.
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user#last_name AwsUser#last_name}.
	// Experimental.
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_user#secondary_email AwsUser#secondary_email}.
	// Experimental.
	SecondaryEmail *string `field:"optional" json:"secondaryEmail" yaml:"secondaryEmail"`
}

