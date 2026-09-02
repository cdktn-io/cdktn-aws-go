package awsappsync

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfResolverConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#api_id TfResolver#api_id}.
	// Experimental.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#field TfResolver#field}.
	// Experimental.
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#type TfResolver#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// caching_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#caching_config TfResolver#caching_config}
	// Experimental.
	CachingConfig *TfResolver_CachingConfigProperty `field:"optional" json:"cachingConfig" yaml:"cachingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#code TfResolver#code}.
	// Experimental.
	Code *string `field:"optional" json:"code" yaml:"code"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#data_source TfResolver#data_source}.
	// Experimental.
	DataSource *string `field:"optional" json:"dataSource" yaml:"dataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#id TfResolver#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#kind TfResolver#kind}.
	// Experimental.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#max_batch_size TfResolver#max_batch_size}.
	// Experimental.
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// pipeline_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#pipeline_config TfResolver#pipeline_config}
	// Experimental.
	PipelineConfig *TfResolver_PipelineConfigProperty `field:"optional" json:"pipelineConfig" yaml:"pipelineConfig"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#region TfResolver#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#request_template TfResolver#request_template}.
	// Experimental.
	RequestTemplate *string `field:"optional" json:"requestTemplate" yaml:"requestTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#response_template TfResolver#response_template}.
	// Experimental.
	ResponseTemplate *string `field:"optional" json:"responseTemplate" yaml:"responseTemplate"`
	// runtime block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#runtime TfResolver#runtime}
	// Experimental.
	Runtime *TfResolver_RuntimeProperty `field:"optional" json:"runtime" yaml:"runtime"`
	// sync_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#sync_config TfResolver#sync_config}
	// Experimental.
	SyncConfig *TfResolver_SyncConfigProperty `field:"optional" json:"syncConfig" yaml:"syncConfig"`
}

