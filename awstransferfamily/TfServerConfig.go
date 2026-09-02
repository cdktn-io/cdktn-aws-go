package awstransferfamily

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfServerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#certificate TfServer#certificate}.
	// Experimental.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#directory_id TfServer#directory_id}.
	// Experimental.
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#domain TfServer#domain}.
	// Experimental.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// endpoint_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#endpoint_details TfServer#endpoint_details}
	// Experimental.
	EndpointDetails *TfServer_EndpointDetailsProperty `field:"optional" json:"endpointDetails" yaml:"endpointDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#endpoint_type TfServer#endpoint_type}.
	// Experimental.
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#force_destroy TfServer#force_destroy}.
	// Experimental.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#function TfServer#function}.
	// Experimental.
	Function *string `field:"optional" json:"function" yaml:"function"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#host_key TfServer#host_key}.
	// Experimental.
	HostKey *string `field:"optional" json:"hostKey" yaml:"hostKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#id TfServer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#identity_provider_type TfServer#identity_provider_type}.
	// Experimental.
	IdentityProviderType *string `field:"optional" json:"identityProviderType" yaml:"identityProviderType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#invocation_role TfServer#invocation_role}.
	// Experimental.
	InvocationRole *string `field:"optional" json:"invocationRole" yaml:"invocationRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#ip_address_type TfServer#ip_address_type}.
	// Experimental.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#logging_role TfServer#logging_role}.
	// Experimental.
	LoggingRole *string `field:"optional" json:"loggingRole" yaml:"loggingRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#post_authentication_login_banner TfServer#post_authentication_login_banner}.
	// Experimental.
	PostAuthenticationLoginBanner *string `field:"optional" json:"postAuthenticationLoginBanner" yaml:"postAuthenticationLoginBanner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#pre_authentication_login_banner TfServer#pre_authentication_login_banner}.
	// Experimental.
	PreAuthenticationLoginBanner *string `field:"optional" json:"preAuthenticationLoginBanner" yaml:"preAuthenticationLoginBanner"`
	// protocol_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#protocol_details TfServer#protocol_details}
	// Experimental.
	ProtocolDetails *TfServer_ProtocolDetailsProperty `field:"optional" json:"protocolDetails" yaml:"protocolDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#protocols TfServer#protocols}.
	// Experimental.
	Protocols *[]*string `field:"optional" json:"protocols" yaml:"protocols"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#region TfServer#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// s3_storage_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#s3_storage_options TfServer#s3_storage_options}
	// Experimental.
	S3StorageOptions *TfServer_S3StorageOptionsProperty `field:"optional" json:"s3StorageOptions" yaml:"s3StorageOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#security_policy_name TfServer#security_policy_name}.
	// Experimental.
	SecurityPolicyName *string `field:"optional" json:"securityPolicyName" yaml:"securityPolicyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#sftp_authentication_methods TfServer#sftp_authentication_methods}.
	// Experimental.
	SftpAuthenticationMethods *string `field:"optional" json:"sftpAuthenticationMethods" yaml:"sftpAuthenticationMethods"`
	// This is a set of arns of destinations that will receive structured logs from the transfer server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#structured_log_destinations TfServer#structured_log_destinations}
	// Experimental.
	StructuredLogDestinations *[]*string `field:"optional" json:"structuredLogDestinations" yaml:"structuredLogDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#tags TfServer#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#tags_all TfServer#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#url TfServer#url}.
	// Experimental.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// workflow_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#workflow_details TfServer#workflow_details}
	// Experimental.
	WorkflowDetails *TfServer_WorkflowDetailsProperty `field:"optional" json:"workflowDetails" yaml:"workflowDetails"`
}

