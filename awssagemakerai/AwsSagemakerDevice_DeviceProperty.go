package awssagemakerai


// Experimental.
type AwsSagemakerDevice_DeviceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_device#device_name AwsSagemakerDevice#device_name}.
	// Experimental.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_device#description AwsSagemakerDevice#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_device#iot_thing_name AwsSagemakerDevice#iot_thing_name}.
	// Experimental.
	IotThingName *string `field:"optional" json:"iotThingName" yaml:"iotThingName"`
}

