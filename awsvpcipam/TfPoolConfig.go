package awsvpcipam

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPoolConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#address_family TfPool#address_family}.
	// Experimental.
	AddressFamily *string `field:"required" json:"addressFamily" yaml:"addressFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#ipam_scope_id TfPool#ipam_scope_id}.
	// Experimental.
	IpamScopeId *string `field:"required" json:"ipamScopeId" yaml:"ipamScopeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#allocation_default_netmask_length TfPool#allocation_default_netmask_length}.
	// Experimental.
	AllocationDefaultNetmaskLength *float64 `field:"optional" json:"allocationDefaultNetmaskLength" yaml:"allocationDefaultNetmaskLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#allocation_max_netmask_length TfPool#allocation_max_netmask_length}.
	// Experimental.
	AllocationMaxNetmaskLength *float64 `field:"optional" json:"allocationMaxNetmaskLength" yaml:"allocationMaxNetmaskLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#allocation_min_netmask_length TfPool#allocation_min_netmask_length}.
	// Experimental.
	AllocationMinNetmaskLength *float64 `field:"optional" json:"allocationMinNetmaskLength" yaml:"allocationMinNetmaskLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#allocation_resource_tags TfPool#allocation_resource_tags}.
	// Experimental.
	AllocationResourceTags *map[string]*string `field:"optional" json:"allocationResourceTags" yaml:"allocationResourceTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#auto_import TfPool#auto_import}.
	// Experimental.
	AutoImport interface{} `field:"optional" json:"autoImport" yaml:"autoImport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#aws_service TfPool#aws_service}.
	// Experimental.
	AwsService *string `field:"optional" json:"awsService" yaml:"awsService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#cascade TfPool#cascade}.
	// Experimental.
	Cascade interface{} `field:"optional" json:"cascade" yaml:"cascade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#description TfPool#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#id TfPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#locale TfPool#locale}.
	// Experimental.
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#public_ip_source TfPool#public_ip_source}.
	// Experimental.
	PublicIpSource *string `field:"optional" json:"publicIpSource" yaml:"publicIpSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#publicly_advertisable TfPool#publicly_advertisable}.
	// Experimental.
	PubliclyAdvertisable interface{} `field:"optional" json:"publiclyAdvertisable" yaml:"publiclyAdvertisable"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#region TfPool#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#source_ipam_pool_id TfPool#source_ipam_pool_id}.
	// Experimental.
	SourceIpamPoolId *string `field:"optional" json:"sourceIpamPoolId" yaml:"sourceIpamPoolId"`
	// source_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#source_resource TfPool#source_resource}
	// Experimental.
	SourceResource *TfPool_SourceResourceProperty `field:"optional" json:"sourceResource" yaml:"sourceResource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#tags TfPool#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#tags_all TfPool#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool#timeouts TfPool#timeouts}
	// Experimental.
	Timeouts *TfPool_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

