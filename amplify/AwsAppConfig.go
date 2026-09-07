package amplify

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#name AwsApp#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#access_token AwsApp#access_token}.
	// Experimental.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// auto_branch_creation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#auto_branch_creation_config AwsApp#auto_branch_creation_config}
	// Experimental.
	AutoBranchCreationConfig *AwsApp_AutoBranchCreationConfigProperty `field:"optional" json:"autoBranchCreationConfig" yaml:"autoBranchCreationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#auto_branch_creation_patterns AwsApp#auto_branch_creation_patterns}.
	// Experimental.
	AutoBranchCreationPatterns *[]*string `field:"optional" json:"autoBranchCreationPatterns" yaml:"autoBranchCreationPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#basic_auth_credentials AwsApp#basic_auth_credentials}.
	// Experimental.
	BasicAuthCredentials *string `field:"optional" json:"basicAuthCredentials" yaml:"basicAuthCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#build_spec AwsApp#build_spec}.
	// Experimental.
	BuildSpec *string `field:"optional" json:"buildSpec" yaml:"buildSpec"`
	// cache_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#cache_config AwsApp#cache_config}
	// Experimental.
	CacheConfig *AwsApp_CacheConfigProperty `field:"optional" json:"cacheConfig" yaml:"cacheConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#compute_role_arn AwsApp#compute_role_arn}.
	// Experimental.
	ComputeRoleArn *string `field:"optional" json:"computeRoleArn" yaml:"computeRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#custom_headers AwsApp#custom_headers}.
	// Experimental.
	CustomHeaders *string `field:"optional" json:"customHeaders" yaml:"customHeaders"`
	// custom_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#custom_rule AwsApp#custom_rule}
	// Experimental.
	CustomRule interface{} `field:"optional" json:"customRule" yaml:"customRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#description AwsApp#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#enable_auto_branch_creation AwsApp#enable_auto_branch_creation}.
	// Experimental.
	EnableAutoBranchCreation interface{} `field:"optional" json:"enableAutoBranchCreation" yaml:"enableAutoBranchCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#enable_basic_auth AwsApp#enable_basic_auth}.
	// Experimental.
	EnableBasicAuth interface{} `field:"optional" json:"enableBasicAuth" yaml:"enableBasicAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#enable_branch_auto_build AwsApp#enable_branch_auto_build}.
	// Experimental.
	EnableBranchAutoBuild interface{} `field:"optional" json:"enableBranchAutoBuild" yaml:"enableBranchAutoBuild"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#enable_branch_auto_deletion AwsApp#enable_branch_auto_deletion}.
	// Experimental.
	EnableBranchAutoDeletion interface{} `field:"optional" json:"enableBranchAutoDeletion" yaml:"enableBranchAutoDeletion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#environment_variables AwsApp#environment_variables}.
	// Experimental.
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#iam_service_role_arn AwsApp#iam_service_role_arn}.
	// Experimental.
	IamServiceRoleArn *string `field:"optional" json:"iamServiceRoleArn" yaml:"iamServiceRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#id AwsApp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// job_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#job_config AwsApp#job_config}
	// Experimental.
	JobConfig *AwsApp_JobConfigProperty `field:"optional" json:"jobConfig" yaml:"jobConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#oauth_token AwsApp#oauth_token}.
	// Experimental.
	OauthToken *string `field:"optional" json:"oauthToken" yaml:"oauthToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#platform AwsApp#platform}.
	// Experimental.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#region AwsApp#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#repository AwsApp#repository}.
	// Experimental.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#tags AwsApp#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app#tags_all AwsApp#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

