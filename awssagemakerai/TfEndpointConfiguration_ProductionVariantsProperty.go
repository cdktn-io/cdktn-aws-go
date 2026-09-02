package awssagemakerai


// Experimental.
type TfEndpointConfiguration_ProductionVariantsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#accelerator_type TfEndpointConfiguration#accelerator_type}.
	// Experimental.
	AcceleratorType *string `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
	// capacity_reservation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#capacity_reservation_config TfEndpointConfiguration#capacity_reservation_config}
	// Experimental.
	CapacityReservationConfig *TfEndpointConfiguration_ProductionVariantsCapacityReservationConfigProperty `field:"optional" json:"capacityReservationConfig" yaml:"capacityReservationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#container_startup_health_check_timeout_in_seconds TfEndpointConfiguration#container_startup_health_check_timeout_in_seconds}.
	// Experimental.
	ContainerStartupHealthCheckTimeoutInSeconds *float64 `field:"optional" json:"containerStartupHealthCheckTimeoutInSeconds" yaml:"containerStartupHealthCheckTimeoutInSeconds"`
	// core_dump_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#core_dump_config TfEndpointConfiguration#core_dump_config}
	// Experimental.
	CoreDumpConfig *TfEndpointConfiguration_ProductionVariantsCoreDumpConfigProperty `field:"optional" json:"coreDumpConfig" yaml:"coreDumpConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#enable_ssm_access TfEndpointConfiguration#enable_ssm_access}.
	// Experimental.
	EnableSsmAccess interface{} `field:"optional" json:"enableSsmAccess" yaml:"enableSsmAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#inference_ami_version TfEndpointConfiguration#inference_ami_version}.
	// Experimental.
	InferenceAmiVersion *string `field:"optional" json:"inferenceAmiVersion" yaml:"inferenceAmiVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#initial_instance_count TfEndpointConfiguration#initial_instance_count}.
	// Experimental.
	InitialInstanceCount *float64 `field:"optional" json:"initialInstanceCount" yaml:"initialInstanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#initial_variant_weight TfEndpointConfiguration#initial_variant_weight}.
	// Experimental.
	InitialVariantWeight *float64 `field:"optional" json:"initialVariantWeight" yaml:"initialVariantWeight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#instance_type TfEndpointConfiguration#instance_type}.
	// Experimental.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// managed_instance_scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#managed_instance_scaling TfEndpointConfiguration#managed_instance_scaling}
	// Experimental.
	ManagedInstanceScaling *TfEndpointConfiguration_ProductionVariantsManagedInstanceScalingProperty `field:"optional" json:"managedInstanceScaling" yaml:"managedInstanceScaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#model_data_download_timeout_in_seconds TfEndpointConfiguration#model_data_download_timeout_in_seconds}.
	// Experimental.
	ModelDataDownloadTimeoutInSeconds *float64 `field:"optional" json:"modelDataDownloadTimeoutInSeconds" yaml:"modelDataDownloadTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#model_name TfEndpointConfiguration#model_name}.
	// Experimental.
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// routing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#routing_config TfEndpointConfiguration#routing_config}
	// Experimental.
	RoutingConfig interface{} `field:"optional" json:"routingConfig" yaml:"routingConfig"`
	// serverless_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#serverless_config TfEndpointConfiguration#serverless_config}
	// Experimental.
	ServerlessConfig *TfEndpointConfiguration_ProductionVariantsServerlessConfigProperty `field:"optional" json:"serverlessConfig" yaml:"serverlessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#variant_name TfEndpointConfiguration#variant_name}.
	// Experimental.
	VariantName *string `field:"optional" json:"variantName" yaml:"variantName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#volume_size_in_gb TfEndpointConfiguration#volume_size_in_gb}.
	// Experimental.
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

