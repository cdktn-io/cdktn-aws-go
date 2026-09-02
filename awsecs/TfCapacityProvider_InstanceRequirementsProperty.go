package awsecs


// Experimental.
type TfCapacityProvider_InstanceRequirementsProperty struct {
	// memory_mib block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#memory_mib TfCapacityProvider#memory_mib}
	// Experimental.
	MemoryMib *TfCapacityProvider_MemoryMibProperty `field:"required" json:"memoryMib" yaml:"memoryMib"`
	// vcpu_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#vcpu_count TfCapacityProvider#vcpu_count}
	// Experimental.
	VcpuCount *TfCapacityProvider_VcpuCountProperty `field:"required" json:"vcpuCount" yaml:"vcpuCount"`
	// accelerator_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#accelerator_count TfCapacityProvider#accelerator_count}
	// Experimental.
	AcceleratorCount *TfCapacityProvider_AcceleratorCountProperty `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#accelerator_manufacturers TfCapacityProvider#accelerator_manufacturers}.
	// Experimental.
	AcceleratorManufacturers *[]*string `field:"optional" json:"acceleratorManufacturers" yaml:"acceleratorManufacturers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#accelerator_names TfCapacityProvider#accelerator_names}.
	// Experimental.
	AcceleratorNames *[]*string `field:"optional" json:"acceleratorNames" yaml:"acceleratorNames"`
	// accelerator_total_memory_mib block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#accelerator_total_memory_mib TfCapacityProvider#accelerator_total_memory_mib}
	// Experimental.
	AcceleratorTotalMemoryMib *TfCapacityProvider_AcceleratorTotalMemoryMibProperty `field:"optional" json:"acceleratorTotalMemoryMib" yaml:"acceleratorTotalMemoryMib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#accelerator_types TfCapacityProvider#accelerator_types}.
	// Experimental.
	AcceleratorTypes *[]*string `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#allowed_instance_types TfCapacityProvider#allowed_instance_types}.
	// Experimental.
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#bare_metal TfCapacityProvider#bare_metal}.
	// Experimental.
	BareMetal *string `field:"optional" json:"bareMetal" yaml:"bareMetal"`
	// baseline_ebs_bandwidth_mbps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#baseline_ebs_bandwidth_mbps TfCapacityProvider#baseline_ebs_bandwidth_mbps}
	// Experimental.
	BaselineEbsBandwidthMbps *TfCapacityProvider_BaselineEbsBandwidthMbpsProperty `field:"optional" json:"baselineEbsBandwidthMbps" yaml:"baselineEbsBandwidthMbps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#burstable_performance TfCapacityProvider#burstable_performance}.
	// Experimental.
	BurstablePerformance *string `field:"optional" json:"burstablePerformance" yaml:"burstablePerformance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#cpu_manufacturers TfCapacityProvider#cpu_manufacturers}.
	// Experimental.
	CpuManufacturers *[]*string `field:"optional" json:"cpuManufacturers" yaml:"cpuManufacturers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#excluded_instance_types TfCapacityProvider#excluded_instance_types}.
	// Experimental.
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#instance_generations TfCapacityProvider#instance_generations}.
	// Experimental.
	InstanceGenerations *[]*string `field:"optional" json:"instanceGenerations" yaml:"instanceGenerations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#local_storage TfCapacityProvider#local_storage}.
	// Experimental.
	LocalStorage *string `field:"optional" json:"localStorage" yaml:"localStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#local_storage_types TfCapacityProvider#local_storage_types}.
	// Experimental.
	LocalStorageTypes *[]*string `field:"optional" json:"localStorageTypes" yaml:"localStorageTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#max_spot_price_as_percentage_of_optimal_on_demand_price TfCapacityProvider#max_spot_price_as_percentage_of_optimal_on_demand_price}.
	// Experimental.
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice *float64 `field:"optional" json:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice" yaml:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice"`
	// memory_gib_per_vcpu block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#memory_gib_per_vcpu TfCapacityProvider#memory_gib_per_vcpu}
	// Experimental.
	MemoryGibPerVcpu *TfCapacityProvider_MemoryGibPerVcpuProperty `field:"optional" json:"memoryGibPerVcpu" yaml:"memoryGibPerVcpu"`
	// network_bandwidth_gbps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#network_bandwidth_gbps TfCapacityProvider#network_bandwidth_gbps}
	// Experimental.
	NetworkBandwidthGbps *TfCapacityProvider_NetworkBandwidthGbpsProperty `field:"optional" json:"networkBandwidthGbps" yaml:"networkBandwidthGbps"`
	// network_interface_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#network_interface_count TfCapacityProvider#network_interface_count}
	// Experimental.
	NetworkInterfaceCount *TfCapacityProvider_NetworkInterfaceCountProperty `field:"optional" json:"networkInterfaceCount" yaml:"networkInterfaceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#on_demand_max_price_percentage_over_lowest_price TfCapacityProvider#on_demand_max_price_percentage_over_lowest_price}.
	// Experimental.
	OnDemandMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"onDemandMaxPricePercentageOverLowestPrice" yaml:"onDemandMaxPricePercentageOverLowestPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#require_hibernate_support TfCapacityProvider#require_hibernate_support}.
	// Experimental.
	RequireHibernateSupport interface{} `field:"optional" json:"requireHibernateSupport" yaml:"requireHibernateSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#spot_max_price_percentage_over_lowest_price TfCapacityProvider#spot_max_price_percentage_over_lowest_price}.
	// Experimental.
	SpotMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"spotMaxPricePercentageOverLowestPrice" yaml:"spotMaxPricePercentageOverLowestPrice"`
	// total_local_storage_gb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#total_local_storage_gb TfCapacityProvider#total_local_storage_gb}
	// Experimental.
	TotalLocalStorageGb *TfCapacityProvider_TotalLocalStorageGbProperty `field:"optional" json:"totalLocalStorageGb" yaml:"totalLocalStorageGb"`
}

