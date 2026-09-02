package awslakeformation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfOptInConfig struct {
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
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#condition TfOptIn#condition}
	// Experimental.
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#principal TfOptIn#principal}
	// Experimental.
	Principal interface{} `field:"optional" json:"principal" yaml:"principal"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#region TfOptIn#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// resource_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#resource_data TfOptIn#resource_data}
	// Experimental.
	ResourceData interface{} `field:"optional" json:"resourceData" yaml:"resourceData"`
}

