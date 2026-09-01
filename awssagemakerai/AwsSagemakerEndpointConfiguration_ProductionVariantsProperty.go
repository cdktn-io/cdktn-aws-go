package awssagemakerai


// Experimental.
type AwsSagemakerEndpointConfiguration_ProductionVariantsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#accelerator_type AwsSagemakerEndpointConfiguration#accelerator_type}.
	// Experimental.
	AcceleratorType *string `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
	// capacity_reservation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#capacity_reservation_config AwsSagemakerEndpointConfiguration#capacity_reservation_config}
	// Experimental.
	CapacityReservationConfig *AwsSagemakerEndpointConfiguration_ProductionVariantsCapacityReservationConfigProperty `field:"optional" json:"capacityReservationConfig" yaml:"capacityReservationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#container_startup_health_check_timeout_in_seconds AwsSagemakerEndpointConfiguration#container_startup_health_check_timeout_in_seconds}.
	// Experimental.
	ContainerStartupHealthCheckTimeoutInSeconds *float64 `field:"optional" json:"containerStartupHealthCheckTimeoutInSeconds" yaml:"containerStartupHealthCheckTimeoutInSeconds"`
	// core_dump_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#core_dump_config AwsSagemakerEndpointConfiguration#core_dump_config}
	// Experimental.
	CoreDumpConfig *AwsSagemakerEndpointConfiguration_ProductionVariantsCoreDumpConfigProperty `field:"optional" json:"coreDumpConfig" yaml:"coreDumpConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#enable_ssm_access AwsSagemakerEndpointConfiguration#enable_ssm_access}.
	// Experimental.
	EnableSsmAccess interface{} `field:"optional" json:"enableSsmAccess" yaml:"enableSsmAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#inference_ami_version AwsSagemakerEndpointConfiguration#inference_ami_version}.
	// Experimental.
	InferenceAmiVersion *string `field:"optional" json:"inferenceAmiVersion" yaml:"inferenceAmiVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#initial_instance_count AwsSagemakerEndpointConfiguration#initial_instance_count}.
	// Experimental.
	InitialInstanceCount *float64 `field:"optional" json:"initialInstanceCount" yaml:"initialInstanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#initial_variant_weight AwsSagemakerEndpointConfiguration#initial_variant_weight}.
	// Experimental.
	InitialVariantWeight *float64 `field:"optional" json:"initialVariantWeight" yaml:"initialVariantWeight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#instance_type AwsSagemakerEndpointConfiguration#instance_type}.
	// Experimental.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// managed_instance_scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#managed_instance_scaling AwsSagemakerEndpointConfiguration#managed_instance_scaling}
	// Experimental.
	ManagedInstanceScaling *AwsSagemakerEndpointConfiguration_ProductionVariantsManagedInstanceScalingProperty `field:"optional" json:"managedInstanceScaling" yaml:"managedInstanceScaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#model_data_download_timeout_in_seconds AwsSagemakerEndpointConfiguration#model_data_download_timeout_in_seconds}.
	// Experimental.
	ModelDataDownloadTimeoutInSeconds *float64 `field:"optional" json:"modelDataDownloadTimeoutInSeconds" yaml:"modelDataDownloadTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#model_name AwsSagemakerEndpointConfiguration#model_name}.
	// Experimental.
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// routing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#routing_config AwsSagemakerEndpointConfiguration#routing_config}
	// Experimental.
	RoutingConfig interface{} `field:"optional" json:"routingConfig" yaml:"routingConfig"`
	// serverless_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#serverless_config AwsSagemakerEndpointConfiguration#serverless_config}
	// Experimental.
	ServerlessConfig *AwsSagemakerEndpointConfiguration_ProductionVariantsServerlessConfigProperty `field:"optional" json:"serverlessConfig" yaml:"serverlessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#variant_name AwsSagemakerEndpointConfiguration#variant_name}.
	// Experimental.
	VariantName *string `field:"optional" json:"variantName" yaml:"variantName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#volume_size_in_gb AwsSagemakerEndpointConfiguration#volume_size_in_gb}.
	// Experimental.
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

