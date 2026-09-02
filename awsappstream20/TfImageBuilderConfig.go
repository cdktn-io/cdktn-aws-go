package awsappstream20

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfImageBuilderConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_image_builder#instance_type TfImageBuilder#instance_type}.
	// Experimental.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_image_builder#name TfImageBuilder#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// access_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_image_builder#access_endpoint TfImageBuilder#access_endpoint}
	// Experimental.
	AccessEndpoint interface{} `field:"optional" json:"accessEndpoint" yaml:"accessEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_image_builder#appstream_agent_version TfImageBuilder#appstream_agent_version}.
	// Experimental.
	AppstreamAgentVersion *string `field:"optional" json:"appstreamAgentVersion" yaml:"appstreamAgentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_image_builder#description TfImageBuilder#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_image_builder#display_name TfImageBuilder#display_name}.
	// Experimental.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// domain_join_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_image_builder#domain_join_info TfImageBuilder#domain_join_info}
	// Experimental.
	DomainJoinInfo *TfImageBuilder_DomainJoinInfoProperty `field:"optional" json:"domainJoinInfo" yaml:"domainJoinInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_image_builder#enable_default_internet_access TfImageBuilder#enable_default_internet_access}.
	// Experimental.
	EnableDefaultInternetAccess interface{} `field:"optional" json:"enableDefaultInternetAccess" yaml:"enableDefaultInternetAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_image_builder#iam_role_arn TfImageBuilder#iam_role_arn}.
	// Experimental.
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_image_builder#id TfImageBuilder#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_image_builder#image_arn TfImageBuilder#image_arn}.
	// Experimental.
	ImageArn *string `field:"optional" json:"imageArn" yaml:"imageArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_image_builder#image_name TfImageBuilder#image_name}.
	// Experimental.
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_image_builder#region TfImageBuilder#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_image_builder#tags TfImageBuilder#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_image_builder#tags_all TfImageBuilder#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_image_builder#vpc_config TfImageBuilder#vpc_config}
	// Experimental.
	VpcConfig *TfImageBuilder_VpcConfigProperty `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

