package awsdatazone


// Experimental.
type AwsDatazoneGlossaryTerm_TermRelationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_glossary_term#classifies AwsDatazoneGlossaryTerm#classifies}.
	// Experimental.
	Classifies *[]*string `field:"optional" json:"classifies" yaml:"classifies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_glossary_term#is_a AwsDatazoneGlossaryTerm#is_a}.
	// Experimental.
	IsA *[]*string `field:"optional" json:"isA" yaml:"isA"`
}

