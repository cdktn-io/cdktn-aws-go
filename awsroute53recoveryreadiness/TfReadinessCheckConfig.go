package awsroute53recoveryreadiness

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfReadinessCheckConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_readiness_check#readiness_check_name TfReadinessCheck#readiness_check_name}.
	// Experimental.
	ReadinessCheckName *string `field:"required" json:"readinessCheckName" yaml:"readinessCheckName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_readiness_check#resource_set_name TfReadinessCheck#resource_set_name}.
	// Experimental.
	ResourceSetName *string `field:"required" json:"resourceSetName" yaml:"resourceSetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_readiness_check#id TfReadinessCheck#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_readiness_check#tags TfReadinessCheck#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_readiness_check#tags_all TfReadinessCheck#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_readiness_check#timeouts TfReadinessCheck#timeouts}
	// Experimental.
	Timeouts *TfReadinessCheck_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

