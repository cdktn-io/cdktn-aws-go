package awscleanrooms

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCollaborationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_collaboration#creator_display_name TfCollaboration#creator_display_name}.
	// Experimental.
	CreatorDisplayName *string `field:"required" json:"creatorDisplayName" yaml:"creatorDisplayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_collaboration#creator_member_abilities TfCollaboration#creator_member_abilities}.
	// Experimental.
	CreatorMemberAbilities *[]*string `field:"required" json:"creatorMemberAbilities" yaml:"creatorMemberAbilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_collaboration#description TfCollaboration#description}.
	// Experimental.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_collaboration#name TfCollaboration#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_collaboration#query_log_status TfCollaboration#query_log_status}.
	// Experimental.
	QueryLogStatus *string `field:"required" json:"queryLogStatus" yaml:"queryLogStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_collaboration#analytics_engine TfCollaboration#analytics_engine}.
	// Experimental.
	AnalyticsEngine *string `field:"optional" json:"analyticsEngine" yaml:"analyticsEngine"`
	// data_encryption_metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_collaboration#data_encryption_metadata TfCollaboration#data_encryption_metadata}
	// Experimental.
	DataEncryptionMetadata *TfCollaboration_DataEncryptionMetadataProperty `field:"optional" json:"dataEncryptionMetadata" yaml:"dataEncryptionMetadata"`
	// member block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_collaboration#member TfCollaboration#member}
	// Experimental.
	Member interface{} `field:"optional" json:"member" yaml:"member"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_collaboration#region TfCollaboration#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_collaboration#tags TfCollaboration#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_collaboration#tags_all TfCollaboration#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_collaboration#timeouts TfCollaboration#timeouts}
	// Experimental.
	Timeouts *TfCollaboration_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

