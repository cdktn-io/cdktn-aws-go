package awsapigateway

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfStageConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#deployment_id TfStage#deployment_id}.
	// Experimental.
	DeploymentId *string `field:"required" json:"deploymentId" yaml:"deploymentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#rest_api_id TfStage#rest_api_id}.
	// Experimental.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#stage_name TfStage#stage_name}.
	// Experimental.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// access_log_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#access_log_settings TfStage#access_log_settings}
	// Experimental.
	AccessLogSettings *TfStage_AccessLogSettingsProperty `field:"optional" json:"accessLogSettings" yaml:"accessLogSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#cache_cluster_enabled TfStage#cache_cluster_enabled}.
	// Experimental.
	CacheClusterEnabled interface{} `field:"optional" json:"cacheClusterEnabled" yaml:"cacheClusterEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#cache_cluster_size TfStage#cache_cluster_size}.
	// Experimental.
	CacheClusterSize *string `field:"optional" json:"cacheClusterSize" yaml:"cacheClusterSize"`
	// canary_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#canary_settings TfStage#canary_settings}
	// Experimental.
	CanarySettings *TfStage_CanarySettingsProperty `field:"optional" json:"canarySettings" yaml:"canarySettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#client_certificate_id TfStage#client_certificate_id}.
	// Experimental.
	ClientCertificateId *string `field:"optional" json:"clientCertificateId" yaml:"clientCertificateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#description TfStage#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#documentation_version TfStage#documentation_version}.
	// Experimental.
	DocumentationVersion *string `field:"optional" json:"documentationVersion" yaml:"documentationVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#id TfStage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#region TfStage#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#tags TfStage#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#tags_all TfStage#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#variables TfStage#variables}.
	// Experimental.
	Variables *map[string]*string `field:"optional" json:"variables" yaml:"variables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#xray_tracing_enabled TfStage#xray_tracing_enabled}.
	// Experimental.
	XrayTracingEnabled interface{} `field:"optional" json:"xrayTracingEnabled" yaml:"xrayTracingEnabled"`
}

