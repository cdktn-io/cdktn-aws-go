package awsverifiedaccess

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#attachment_type TfEndpoint#attachment_type}.
	// Experimental.
	AttachmentType *string `field:"required" json:"attachmentType" yaml:"attachmentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#endpoint_type TfEndpoint#endpoint_type}.
	// Experimental.
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#verified_access_group_id TfEndpoint#verified_access_group_id}.
	// Experimental.
	VerifiedAccessGroupId *string `field:"required" json:"verifiedAccessGroupId" yaml:"verifiedAccessGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#application_domain TfEndpoint#application_domain}.
	// Experimental.
	ApplicationDomain *string `field:"optional" json:"applicationDomain" yaml:"applicationDomain"`
	// cidr_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#cidr_options TfEndpoint#cidr_options}
	// Experimental.
	CidrOptions *TfEndpoint_CidrOptionsProperty `field:"optional" json:"cidrOptions" yaml:"cidrOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#description TfEndpoint#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#domain_certificate_arn TfEndpoint#domain_certificate_arn}.
	// Experimental.
	DomainCertificateArn *string `field:"optional" json:"domainCertificateArn" yaml:"domainCertificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#endpoint_domain_prefix TfEndpoint#endpoint_domain_prefix}.
	// Experimental.
	EndpointDomainPrefix *string `field:"optional" json:"endpointDomainPrefix" yaml:"endpointDomainPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#id TfEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// load_balancer_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#load_balancer_options TfEndpoint#load_balancer_options}
	// Experimental.
	LoadBalancerOptions *TfEndpoint_LoadBalancerOptionsProperty `field:"optional" json:"loadBalancerOptions" yaml:"loadBalancerOptions"`
	// network_interface_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#network_interface_options TfEndpoint#network_interface_options}
	// Experimental.
	NetworkInterfaceOptions *TfEndpoint_NetworkInterfaceOptionsProperty `field:"optional" json:"networkInterfaceOptions" yaml:"networkInterfaceOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#policy_document TfEndpoint#policy_document}.
	// Experimental.
	PolicyDocument *string `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// rds_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#rds_options TfEndpoint#rds_options}
	// Experimental.
	RdsOptions *TfEndpoint_RdsOptionsProperty `field:"optional" json:"rdsOptions" yaml:"rdsOptions"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#region TfEndpoint#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#security_group_ids TfEndpoint#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// sse_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#sse_specification TfEndpoint#sse_specification}
	// Experimental.
	SseSpecification *TfEndpoint_SseSpecificationProperty `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#tags TfEndpoint#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#tags_all TfEndpoint#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#timeouts TfEndpoint#timeouts}
	// Experimental.
	Timeouts *TfEndpoint_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

