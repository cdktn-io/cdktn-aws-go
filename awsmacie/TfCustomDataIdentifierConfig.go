package awsmacie

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCustomDataIdentifierConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_custom_data_identifier#description TfCustomDataIdentifier#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_custom_data_identifier#id TfCustomDataIdentifier#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_custom_data_identifier#ignore_words TfCustomDataIdentifier#ignore_words}.
	// Experimental.
	IgnoreWords *[]*string `field:"optional" json:"ignoreWords" yaml:"ignoreWords"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_custom_data_identifier#keywords TfCustomDataIdentifier#keywords}.
	// Experimental.
	Keywords *[]*string `field:"optional" json:"keywords" yaml:"keywords"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_custom_data_identifier#maximum_match_distance TfCustomDataIdentifier#maximum_match_distance}.
	// Experimental.
	MaximumMatchDistance *float64 `field:"optional" json:"maximumMatchDistance" yaml:"maximumMatchDistance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_custom_data_identifier#name TfCustomDataIdentifier#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_custom_data_identifier#name_prefix TfCustomDataIdentifier#name_prefix}.
	// Experimental.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_custom_data_identifier#regex TfCustomDataIdentifier#regex}.
	// Experimental.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_custom_data_identifier#region TfCustomDataIdentifier#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_custom_data_identifier#tags TfCustomDataIdentifier#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_custom_data_identifier#tags_all TfCustomDataIdentifier#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_custom_data_identifier#timeouts TfCustomDataIdentifier#timeouts}
	// Experimental.
	Timeouts *TfCustomDataIdentifier_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

