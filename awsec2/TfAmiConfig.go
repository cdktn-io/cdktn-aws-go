package awsec2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAmiConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#name TfAmi#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#architecture TfAmi#architecture}.
	// Experimental.
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#boot_mode TfAmi#boot_mode}.
	// Experimental.
	BootMode *string `field:"optional" json:"bootMode" yaml:"bootMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#deprecation_time TfAmi#deprecation_time}.
	// Experimental.
	DeprecationTime *string `field:"optional" json:"deprecationTime" yaml:"deprecationTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#description TfAmi#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// ebs_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#ebs_block_device TfAmi#ebs_block_device}
	// Experimental.
	EbsBlockDevice interface{} `field:"optional" json:"ebsBlockDevice" yaml:"ebsBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#ena_support TfAmi#ena_support}.
	// Experimental.
	EnaSupport interface{} `field:"optional" json:"enaSupport" yaml:"enaSupport"`
	// ephemeral_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#ephemeral_block_device TfAmi#ephemeral_block_device}
	// Experimental.
	EphemeralBlockDevice interface{} `field:"optional" json:"ephemeralBlockDevice" yaml:"ephemeralBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#id TfAmi#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#image_location TfAmi#image_location}.
	// Experimental.
	ImageLocation *string `field:"optional" json:"imageLocation" yaml:"imageLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#imds_support TfAmi#imds_support}.
	// Experimental.
	ImdsSupport *string `field:"optional" json:"imdsSupport" yaml:"imdsSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#kernel_id TfAmi#kernel_id}.
	// Experimental.
	KernelId *string `field:"optional" json:"kernelId" yaml:"kernelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#ramdisk_id TfAmi#ramdisk_id}.
	// Experimental.
	RamdiskId *string `field:"optional" json:"ramdiskId" yaml:"ramdiskId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#region TfAmi#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#root_device_name TfAmi#root_device_name}.
	// Experimental.
	RootDeviceName *string `field:"optional" json:"rootDeviceName" yaml:"rootDeviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#sriov_net_support TfAmi#sriov_net_support}.
	// Experimental.
	SriovNetSupport *string `field:"optional" json:"sriovNetSupport" yaml:"sriovNetSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#tags TfAmi#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#tags_all TfAmi#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#timeouts TfAmi#timeouts}
	// Experimental.
	Timeouts *TfAmi_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#tpm_support TfAmi#tpm_support}.
	// Experimental.
	TpmSupport *string `field:"optional" json:"tpmSupport" yaml:"tpmSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#uefi_data TfAmi#uefi_data}.
	// Experimental.
	UefiData *string `field:"optional" json:"uefiData" yaml:"uefiData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami#virtualization_type TfAmi#virtualization_type}.
	// Experimental.
	VirtualizationType *string `field:"optional" json:"virtualizationType" yaml:"virtualizationType"`
}

