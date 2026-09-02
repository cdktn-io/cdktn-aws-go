package awscodebuild

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfProjectConfig struct {
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
	// artifacts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#artifacts TfProject#artifacts}
	// Experimental.
	Artifacts *TfProject_ArtifactsProperty `field:"required" json:"artifacts" yaml:"artifacts"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#environment TfProject#environment}
	// Experimental.
	Environment *TfProject_EnvironmentProperty `field:"required" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#name TfProject#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#service_role TfProject#service_role}.
	// Experimental.
	ServiceRole *string `field:"required" json:"serviceRole" yaml:"serviceRole"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#source TfProject#source}
	// Experimental.
	Source *TfProject_SourceProperty `field:"required" json:"source" yaml:"source"`
	// Maximum number of additional automatic retries after a failed build. The default value is 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#auto_retry_limit TfProject#auto_retry_limit}
	// Experimental.
	AutoRetryLimit *float64 `field:"optional" json:"autoRetryLimit" yaml:"autoRetryLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#badge_enabled TfProject#badge_enabled}.
	// Experimental.
	BadgeEnabled interface{} `field:"optional" json:"badgeEnabled" yaml:"badgeEnabled"`
	// build_batch_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#build_batch_config TfProject#build_batch_config}
	// Experimental.
	BuildBatchConfig *TfProject_BuildBatchConfigProperty `field:"optional" json:"buildBatchConfig" yaml:"buildBatchConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#build_timeout TfProject#build_timeout}.
	// Experimental.
	BuildTimeout *float64 `field:"optional" json:"buildTimeout" yaml:"buildTimeout"`
	// cache block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#cache TfProject#cache}
	// Experimental.
	Cache *TfProject_CacheProperty `field:"optional" json:"cache" yaml:"cache"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#concurrent_build_limit TfProject#concurrent_build_limit}.
	// Experimental.
	ConcurrentBuildLimit *float64 `field:"optional" json:"concurrentBuildLimit" yaml:"concurrentBuildLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#description TfProject#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#encryption_key TfProject#encryption_key}.
	// Experimental.
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// file_system_locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#file_system_locations TfProject#file_system_locations}
	// Experimental.
	FileSystemLocations interface{} `field:"optional" json:"fileSystemLocations" yaml:"fileSystemLocations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#id TfProject#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// logs_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#logs_config TfProject#logs_config}
	// Experimental.
	LogsConfig *TfProject_LogsConfigProperty `field:"optional" json:"logsConfig" yaml:"logsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#project_visibility TfProject#project_visibility}.
	// Experimental.
	ProjectVisibility *string `field:"optional" json:"projectVisibility" yaml:"projectVisibility"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#queued_timeout TfProject#queued_timeout}.
	// Experimental.
	QueuedTimeout *float64 `field:"optional" json:"queuedTimeout" yaml:"queuedTimeout"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#region TfProject#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#resource_access_role TfProject#resource_access_role}.
	// Experimental.
	ResourceAccessRole *string `field:"optional" json:"resourceAccessRole" yaml:"resourceAccessRole"`
	// secondary_artifacts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#secondary_artifacts TfProject#secondary_artifacts}
	// Experimental.
	SecondaryArtifacts interface{} `field:"optional" json:"secondaryArtifacts" yaml:"secondaryArtifacts"`
	// secondary_sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#secondary_sources TfProject#secondary_sources}
	// Experimental.
	SecondarySources interface{} `field:"optional" json:"secondarySources" yaml:"secondarySources"`
	// secondary_source_version block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#secondary_source_version TfProject#secondary_source_version}
	// Experimental.
	SecondarySourceVersion interface{} `field:"optional" json:"secondarySourceVersion" yaml:"secondarySourceVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#source_version TfProject#source_version}.
	// Experimental.
	SourceVersion *string `field:"optional" json:"sourceVersion" yaml:"sourceVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#tags TfProject#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#tags_all TfProject#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#vpc_config TfProject#vpc_config}
	// Experimental.
	VpcConfig *TfProject_VpcConfigProperty `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

