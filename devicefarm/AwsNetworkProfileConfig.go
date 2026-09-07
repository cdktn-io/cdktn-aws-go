package devicefarm

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsNetworkProfileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devicefarm_network_profile#name AwsNetworkProfile#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devicefarm_network_profile#project_arn AwsNetworkProfile#project_arn}.
	// Experimental.
	ProjectArn *string `field:"required" json:"projectArn" yaml:"projectArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devicefarm_network_profile#description AwsNetworkProfile#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devicefarm_network_profile#downlink_bandwidth_bits AwsNetworkProfile#downlink_bandwidth_bits}.
	// Experimental.
	DownlinkBandwidthBits *float64 `field:"optional" json:"downlinkBandwidthBits" yaml:"downlinkBandwidthBits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devicefarm_network_profile#downlink_delay_ms AwsNetworkProfile#downlink_delay_ms}.
	// Experimental.
	DownlinkDelayMs *float64 `field:"optional" json:"downlinkDelayMs" yaml:"downlinkDelayMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devicefarm_network_profile#downlink_jitter_ms AwsNetworkProfile#downlink_jitter_ms}.
	// Experimental.
	DownlinkJitterMs *float64 `field:"optional" json:"downlinkJitterMs" yaml:"downlinkJitterMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devicefarm_network_profile#downlink_loss_percent AwsNetworkProfile#downlink_loss_percent}.
	// Experimental.
	DownlinkLossPercent *float64 `field:"optional" json:"downlinkLossPercent" yaml:"downlinkLossPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devicefarm_network_profile#id AwsNetworkProfile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devicefarm_network_profile#region AwsNetworkProfile#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devicefarm_network_profile#tags AwsNetworkProfile#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devicefarm_network_profile#tags_all AwsNetworkProfile#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devicefarm_network_profile#type AwsNetworkProfile#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devicefarm_network_profile#uplink_bandwidth_bits AwsNetworkProfile#uplink_bandwidth_bits}.
	// Experimental.
	UplinkBandwidthBits *float64 `field:"optional" json:"uplinkBandwidthBits" yaml:"uplinkBandwidthBits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devicefarm_network_profile#uplink_delay_ms AwsNetworkProfile#uplink_delay_ms}.
	// Experimental.
	UplinkDelayMs *float64 `field:"optional" json:"uplinkDelayMs" yaml:"uplinkDelayMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devicefarm_network_profile#uplink_jitter_ms AwsNetworkProfile#uplink_jitter_ms}.
	// Experimental.
	UplinkJitterMs *float64 `field:"optional" json:"uplinkJitterMs" yaml:"uplinkJitterMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devicefarm_network_profile#uplink_loss_percent AwsNetworkProfile#uplink_loss_percent}.
	// Experimental.
	UplinkLossPercent *float64 `field:"optional" json:"uplinkLossPercent" yaml:"uplinkLossPercent"`
}

