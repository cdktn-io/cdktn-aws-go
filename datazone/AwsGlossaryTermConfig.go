package datazone

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGlossaryTermConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_glossary_term#glossary_identifier AwsGlossaryTerm#glossary_identifier}.
	// Experimental.
	GlossaryIdentifier *string `field:"required" json:"glossaryIdentifier" yaml:"glossaryIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_glossary_term#name AwsGlossaryTerm#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_glossary_term#domain_identifier AwsGlossaryTerm#domain_identifier}.
	// Experimental.
	DomainIdentifier *string `field:"optional" json:"domainIdentifier" yaml:"domainIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_glossary_term#long_description AwsGlossaryTerm#long_description}.
	// Experimental.
	LongDescription *string `field:"optional" json:"longDescription" yaml:"longDescription"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_glossary_term#region AwsGlossaryTerm#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_glossary_term#short_description AwsGlossaryTerm#short_description}.
	// Experimental.
	ShortDescription *string `field:"optional" json:"shortDescription" yaml:"shortDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_glossary_term#status AwsGlossaryTerm#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// term_relations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_glossary_term#term_relations AwsGlossaryTerm#term_relations}
	// Experimental.
	TermRelations interface{} `field:"optional" json:"termRelations" yaml:"termRelations"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_glossary_term#timeouts AwsGlossaryTerm#timeouts}
	// Experimental.
	Timeouts *AwsGlossaryTerm_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

