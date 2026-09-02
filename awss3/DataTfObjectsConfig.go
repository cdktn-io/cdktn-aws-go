package awss3

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfObjectsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/s3_objects#bucket DataTfObjects#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/s3_objects#delimiter DataTfObjects#delimiter}.
	// Experimental.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/s3_objects#encoding_type DataTfObjects#encoding_type}.
	// Experimental.
	EncodingType *string `field:"optional" json:"encodingType" yaml:"encodingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/s3_objects#fetch_owner DataTfObjects#fetch_owner}.
	// Experimental.
	FetchOwner interface{} `field:"optional" json:"fetchOwner" yaml:"fetchOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/s3_objects#id DataTfObjects#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/s3_objects#max_keys DataTfObjects#max_keys}.
	// Experimental.
	MaxKeys *float64 `field:"optional" json:"maxKeys" yaml:"maxKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/s3_objects#prefix DataTfObjects#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/s3_objects#region DataTfObjects#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/s3_objects#request_payer DataTfObjects#request_payer}.
	// Experimental.
	RequestPayer *string `field:"optional" json:"requestPayer" yaml:"requestPayer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/s3_objects#start_after DataTfObjects#start_after}.
	// Experimental.
	StartAfter *string `field:"optional" json:"startAfter" yaml:"startAfter"`
}

