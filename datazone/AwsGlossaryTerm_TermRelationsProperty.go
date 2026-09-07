package datazone


// Experimental.
type AwsGlossaryTerm_TermRelationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_glossary_term#classifies AwsGlossaryTerm#classifies}.
	// Experimental.
	Classifies *[]*string `field:"optional" json:"classifies" yaml:"classifies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_glossary_term#is_a AwsGlossaryTerm#is_a}.
	// Experimental.
	IsA *[]*string `field:"optional" json:"isA" yaml:"isA"`
}

