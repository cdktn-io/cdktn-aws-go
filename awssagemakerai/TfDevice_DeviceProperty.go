package awssagemakerai


// Experimental.
type TfDevice_DeviceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_device#device_name TfDevice#device_name}.
	// Experimental.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_device#description TfDevice#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_device#iot_thing_name TfDevice#iot_thing_name}.
	// Experimental.
	IotThingName *string `field:"optional" json:"iotThingName" yaml:"iotThingName"`
}

