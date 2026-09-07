package connect


// Experimental.
type DataAwsBotAssociation_LexBotProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/connect_bot_association#name DataAwsBotAssociation#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/connect_bot_association#lex_region DataAwsBotAssociation#lex_region}.
	// Experimental.
	LexRegion *string `field:"optional" json:"lexRegion" yaml:"lexRegion"`
}

