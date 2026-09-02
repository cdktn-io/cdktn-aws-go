package awsappstream20

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfStackConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#name TfStack#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// access_endpoints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#access_endpoints TfStack#access_endpoints}
	// Experimental.
	AccessEndpoints interface{} `field:"optional" json:"accessEndpoints" yaml:"accessEndpoints"`
	// application_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#application_settings TfStack#application_settings}
	// Experimental.
	ApplicationSettings *TfStack_ApplicationSettingsProperty `field:"optional" json:"applicationSettings" yaml:"applicationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#description TfStack#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#display_name TfStack#display_name}.
	// Experimental.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#embed_host_domains TfStack#embed_host_domains}.
	// Experimental.
	EmbedHostDomains *[]*string `field:"optional" json:"embedHostDomains" yaml:"embedHostDomains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#feedback_url TfStack#feedback_url}.
	// Experimental.
	FeedbackUrl *string `field:"optional" json:"feedbackUrl" yaml:"feedbackUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#id TfStack#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#redirect_url TfStack#redirect_url}.
	// Experimental.
	RedirectUrl *string `field:"optional" json:"redirectUrl" yaml:"redirectUrl"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#region TfStack#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// storage_connectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#storage_connectors TfStack#storage_connectors}
	// Experimental.
	StorageConnectors interface{} `field:"optional" json:"storageConnectors" yaml:"storageConnectors"`
	// streaming_experience_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#streaming_experience_settings TfStack#streaming_experience_settings}
	// Experimental.
	StreamingExperienceSettings *TfStack_StreamingExperienceSettingsProperty `field:"optional" json:"streamingExperienceSettings" yaml:"streamingExperienceSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#tags TfStack#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#tags_all TfStack#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// user_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#user_settings TfStack#user_settings}
	// Experimental.
	UserSettings interface{} `field:"optional" json:"userSettings" yaml:"userSettings"`
}

