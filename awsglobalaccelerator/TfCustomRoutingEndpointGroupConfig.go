package awsglobalaccelerator

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCustomRoutingEndpointGroupConfig struct {
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
	// destination_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_custom_routing_endpoint_group#destination_configuration TfCustomRoutingEndpointGroup#destination_configuration}
	// Experimental.
	DestinationConfiguration interface{} `field:"required" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_custom_routing_endpoint_group#listener_arn TfCustomRoutingEndpointGroup#listener_arn}.
	// Experimental.
	ListenerArn *string `field:"required" json:"listenerArn" yaml:"listenerArn"`
	// endpoint_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_custom_routing_endpoint_group#endpoint_configuration TfCustomRoutingEndpointGroup#endpoint_configuration}
	// Experimental.
	EndpointConfiguration interface{} `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_custom_routing_endpoint_group#endpoint_group_region TfCustomRoutingEndpointGroup#endpoint_group_region}.
	// Experimental.
	EndpointGroupRegion *string `field:"optional" json:"endpointGroupRegion" yaml:"endpointGroupRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_custom_routing_endpoint_group#id TfCustomRoutingEndpointGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_custom_routing_endpoint_group#timeouts TfCustomRoutingEndpointGroup#timeouts}
	// Experimental.
	Timeouts *TfCustomRoutingEndpointGroup_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

