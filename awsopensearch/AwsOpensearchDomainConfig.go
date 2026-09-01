package awsopensearch

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsOpensearchDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#domain_name AwsOpensearchDomain#domain_name}.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#access_policies AwsOpensearchDomain#access_policies}.
	// Experimental.
	AccessPolicies *string `field:"optional" json:"accessPolicies" yaml:"accessPolicies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#advanced_options AwsOpensearchDomain#advanced_options}.
	// Experimental.
	AdvancedOptions *map[string]*string `field:"optional" json:"advancedOptions" yaml:"advancedOptions"`
	// advanced_security_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#advanced_security_options AwsOpensearchDomain#advanced_security_options}
	// Experimental.
	AdvancedSecurityOptions *AwsOpensearchDomain_AdvancedSecurityOptionsProperty `field:"optional" json:"advancedSecurityOptions" yaml:"advancedSecurityOptions"`
	// aiml_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#aiml_options AwsOpensearchDomain#aiml_options}
	// Experimental.
	AimlOptions *AwsOpensearchDomain_AimlOptionsProperty `field:"optional" json:"aimlOptions" yaml:"aimlOptions"`
	// auto_tune_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#auto_tune_options AwsOpensearchDomain#auto_tune_options}
	// Experimental.
	AutoTuneOptions *AwsOpensearchDomain_AutoTuneOptionsProperty `field:"optional" json:"autoTuneOptions" yaml:"autoTuneOptions"`
	// cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#cluster_config AwsOpensearchDomain#cluster_config}
	// Experimental.
	ClusterConfig *AwsOpensearchDomain_ClusterConfigProperty `field:"optional" json:"clusterConfig" yaml:"clusterConfig"`
	// cognito_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#cognito_options AwsOpensearchDomain#cognito_options}
	// Experimental.
	CognitoOptions *AwsOpensearchDomain_CognitoOptionsProperty `field:"optional" json:"cognitoOptions" yaml:"cognitoOptions"`
	// deployment_strategy_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#deployment_strategy_options AwsOpensearchDomain#deployment_strategy_options}
	// Experimental.
	DeploymentStrategyOptions *AwsOpensearchDomain_DeploymentStrategyOptionsProperty `field:"optional" json:"deploymentStrategyOptions" yaml:"deploymentStrategyOptions"`
	// domain_endpoint_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#domain_endpoint_options AwsOpensearchDomain#domain_endpoint_options}
	// Experimental.
	DomainEndpointOptions *AwsOpensearchDomain_DomainEndpointOptionsProperty `field:"optional" json:"domainEndpointOptions" yaml:"domainEndpointOptions"`
	// ebs_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#ebs_options AwsOpensearchDomain#ebs_options}
	// Experimental.
	EbsOptions *AwsOpensearchDomain_EbsOptionsProperty `field:"optional" json:"ebsOptions" yaml:"ebsOptions"`
	// encrypt_at_rest block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#encrypt_at_rest AwsOpensearchDomain#encrypt_at_rest}
	// Experimental.
	EncryptAtRest *AwsOpensearchDomain_EncryptAtRestProperty `field:"optional" json:"encryptAtRest" yaml:"encryptAtRest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#engine_version AwsOpensearchDomain#engine_version}.
	// Experimental.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#id AwsOpensearchDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity_center_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#identity_center_options AwsOpensearchDomain#identity_center_options}
	// Experimental.
	IdentityCenterOptions *AwsOpensearchDomain_IdentityCenterOptionsProperty `field:"optional" json:"identityCenterOptions" yaml:"identityCenterOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#ip_address_type AwsOpensearchDomain#ip_address_type}.
	// Experimental.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// log_publishing_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#log_publishing_options AwsOpensearchDomain#log_publishing_options}
	// Experimental.
	LogPublishingOptions interface{} `field:"optional" json:"logPublishingOptions" yaml:"logPublishingOptions"`
	// node_to_node_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#node_to_node_encryption AwsOpensearchDomain#node_to_node_encryption}
	// Experimental.
	NodeToNodeEncryption *AwsOpensearchDomain_NodeToNodeEncryptionProperty `field:"optional" json:"nodeToNodeEncryption" yaml:"nodeToNodeEncryption"`
	// off_peak_window_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#off_peak_window_options AwsOpensearchDomain#off_peak_window_options}
	// Experimental.
	OffPeakWindowOptions *AwsOpensearchDomain_OffPeakWindowOptionsProperty `field:"optional" json:"offPeakWindowOptions" yaml:"offPeakWindowOptions"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#region AwsOpensearchDomain#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// snapshot_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#snapshot_options AwsOpensearchDomain#snapshot_options}
	// Experimental.
	SnapshotOptions *AwsOpensearchDomain_SnapshotOptionsProperty `field:"optional" json:"snapshotOptions" yaml:"snapshotOptions"`
	// software_update_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#software_update_options AwsOpensearchDomain#software_update_options}
	// Experimental.
	SoftwareUpdateOptions *AwsOpensearchDomain_SoftwareUpdateOptionsProperty `field:"optional" json:"softwareUpdateOptions" yaml:"softwareUpdateOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#tags AwsOpensearchDomain#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#tags_all AwsOpensearchDomain#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#timeouts AwsOpensearchDomain#timeouts}
	// Experimental.
	Timeouts *AwsOpensearchDomain_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vpc_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#vpc_options AwsOpensearchDomain#vpc_options}
	// Experimental.
	VpcOptions *AwsOpensearchDomain_VpcOptionsProperty `field:"optional" json:"vpcOptions" yaml:"vpcOptions"`
}

