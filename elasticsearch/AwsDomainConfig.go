package elasticsearch

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#domain_name AwsDomain#domain_name}.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#access_policies AwsDomain#access_policies}.
	// Experimental.
	AccessPolicies *string `field:"optional" json:"accessPolicies" yaml:"accessPolicies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#advanced_options AwsDomain#advanced_options}.
	// Experimental.
	AdvancedOptions *map[string]*string `field:"optional" json:"advancedOptions" yaml:"advancedOptions"`
	// advanced_security_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#advanced_security_options AwsDomain#advanced_security_options}
	// Experimental.
	AdvancedSecurityOptions *AwsDomain_AdvancedSecurityOptionsProperty `field:"optional" json:"advancedSecurityOptions" yaml:"advancedSecurityOptions"`
	// auto_tune_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#auto_tune_options AwsDomain#auto_tune_options}
	// Experimental.
	AutoTuneOptions *AwsDomain_AutoTuneOptionsProperty `field:"optional" json:"autoTuneOptions" yaml:"autoTuneOptions"`
	// cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#cluster_config AwsDomain#cluster_config}
	// Experimental.
	ClusterConfig *AwsDomain_ClusterConfigProperty `field:"optional" json:"clusterConfig" yaml:"clusterConfig"`
	// cognito_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#cognito_options AwsDomain#cognito_options}
	// Experimental.
	CognitoOptions *AwsDomain_CognitoOptionsProperty `field:"optional" json:"cognitoOptions" yaml:"cognitoOptions"`
	// domain_endpoint_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#domain_endpoint_options AwsDomain#domain_endpoint_options}
	// Experimental.
	DomainEndpointOptions *AwsDomain_DomainEndpointOptionsProperty `field:"optional" json:"domainEndpointOptions" yaml:"domainEndpointOptions"`
	// ebs_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#ebs_options AwsDomain#ebs_options}
	// Experimental.
	EbsOptions *AwsDomain_EbsOptionsProperty `field:"optional" json:"ebsOptions" yaml:"ebsOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#elasticsearch_version AwsDomain#elasticsearch_version}.
	// Experimental.
	ElasticsearchVersion *string `field:"optional" json:"elasticsearchVersion" yaml:"elasticsearchVersion"`
	// encrypt_at_rest block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#encrypt_at_rest AwsDomain#encrypt_at_rest}
	// Experimental.
	EncryptAtRest *AwsDomain_EncryptAtRestProperty `field:"optional" json:"encryptAtRest" yaml:"encryptAtRest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#id AwsDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// log_publishing_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#log_publishing_options AwsDomain#log_publishing_options}
	// Experimental.
	LogPublishingOptions interface{} `field:"optional" json:"logPublishingOptions" yaml:"logPublishingOptions"`
	// node_to_node_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#node_to_node_encryption AwsDomain#node_to_node_encryption}
	// Experimental.
	NodeToNodeEncryption *AwsDomain_NodeToNodeEncryptionProperty `field:"optional" json:"nodeToNodeEncryption" yaml:"nodeToNodeEncryption"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#region AwsDomain#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// snapshot_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#snapshot_options AwsDomain#snapshot_options}
	// Experimental.
	SnapshotOptions *AwsDomain_SnapshotOptionsProperty `field:"optional" json:"snapshotOptions" yaml:"snapshotOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#tags AwsDomain#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#tags_all AwsDomain#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#timeouts AwsDomain#timeouts}
	// Experimental.
	Timeouts *AwsDomain_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vpc_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#vpc_options AwsDomain#vpc_options}
	// Experimental.
	VpcOptions *AwsDomain_VpcOptionsProperty `field:"optional" json:"vpcOptions" yaml:"vpcOptions"`
}

