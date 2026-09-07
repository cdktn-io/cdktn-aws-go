package ecs


// Experimental.
type AwsCapacityProvider_InstanceRequirementsProperty struct {
	// memory_mib block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#memory_mib AwsCapacityProvider#memory_mib}
	// Experimental.
	MemoryMib *AwsCapacityProvider_MemoryMibProperty `field:"required" json:"memoryMib" yaml:"memoryMib"`
	// vcpu_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#vcpu_count AwsCapacityProvider#vcpu_count}
	// Experimental.
	VcpuCount *AwsCapacityProvider_VcpuCountProperty `field:"required" json:"vcpuCount" yaml:"vcpuCount"`
	// accelerator_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#accelerator_count AwsCapacityProvider#accelerator_count}
	// Experimental.
	AcceleratorCount *AwsCapacityProvider_AcceleratorCountProperty `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#accelerator_manufacturers AwsCapacityProvider#accelerator_manufacturers}.
	// Experimental.
	AcceleratorManufacturers *[]*string `field:"optional" json:"acceleratorManufacturers" yaml:"acceleratorManufacturers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#accelerator_names AwsCapacityProvider#accelerator_names}.
	// Experimental.
	AcceleratorNames *[]*string `field:"optional" json:"acceleratorNames" yaml:"acceleratorNames"`
	// accelerator_total_memory_mib block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#accelerator_total_memory_mib AwsCapacityProvider#accelerator_total_memory_mib}
	// Experimental.
	AcceleratorTotalMemoryMib *AwsCapacityProvider_AcceleratorTotalMemoryMibProperty `field:"optional" json:"acceleratorTotalMemoryMib" yaml:"acceleratorTotalMemoryMib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#accelerator_types AwsCapacityProvider#accelerator_types}.
	// Experimental.
	AcceleratorTypes *[]*string `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#allowed_instance_types AwsCapacityProvider#allowed_instance_types}.
	// Experimental.
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#bare_metal AwsCapacityProvider#bare_metal}.
	// Experimental.
	BareMetal *string `field:"optional" json:"bareMetal" yaml:"bareMetal"`
	// baseline_ebs_bandwidth_mbps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#baseline_ebs_bandwidth_mbps AwsCapacityProvider#baseline_ebs_bandwidth_mbps}
	// Experimental.
	BaselineEbsBandwidthMbps *AwsCapacityProvider_BaselineEbsBandwidthMbpsProperty `field:"optional" json:"baselineEbsBandwidthMbps" yaml:"baselineEbsBandwidthMbps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#burstable_performance AwsCapacityProvider#burstable_performance}.
	// Experimental.
	BurstablePerformance *string `field:"optional" json:"burstablePerformance" yaml:"burstablePerformance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#cpu_manufacturers AwsCapacityProvider#cpu_manufacturers}.
	// Experimental.
	CpuManufacturers *[]*string `field:"optional" json:"cpuManufacturers" yaml:"cpuManufacturers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#excluded_instance_types AwsCapacityProvider#excluded_instance_types}.
	// Experimental.
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#instance_generations AwsCapacityProvider#instance_generations}.
	// Experimental.
	InstanceGenerations *[]*string `field:"optional" json:"instanceGenerations" yaml:"instanceGenerations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#local_storage AwsCapacityProvider#local_storage}.
	// Experimental.
	LocalStorage *string `field:"optional" json:"localStorage" yaml:"localStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#local_storage_types AwsCapacityProvider#local_storage_types}.
	// Experimental.
	LocalStorageTypes *[]*string `field:"optional" json:"localStorageTypes" yaml:"localStorageTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#max_spot_price_as_percentage_of_optimal_on_demand_price AwsCapacityProvider#max_spot_price_as_percentage_of_optimal_on_demand_price}.
	// Experimental.
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice *float64 `field:"optional" json:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice" yaml:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice"`
	// memory_gib_per_vcpu block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#memory_gib_per_vcpu AwsCapacityProvider#memory_gib_per_vcpu}
	// Experimental.
	MemoryGibPerVcpu *AwsCapacityProvider_MemoryGibPerVcpuProperty `field:"optional" json:"memoryGibPerVcpu" yaml:"memoryGibPerVcpu"`
	// network_bandwidth_gbps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#network_bandwidth_gbps AwsCapacityProvider#network_bandwidth_gbps}
	// Experimental.
	NetworkBandwidthGbps *AwsCapacityProvider_NetworkBandwidthGbpsProperty `field:"optional" json:"networkBandwidthGbps" yaml:"networkBandwidthGbps"`
	// network_interface_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#network_interface_count AwsCapacityProvider#network_interface_count}
	// Experimental.
	NetworkInterfaceCount *AwsCapacityProvider_NetworkInterfaceCountProperty `field:"optional" json:"networkInterfaceCount" yaml:"networkInterfaceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#on_demand_max_price_percentage_over_lowest_price AwsCapacityProvider#on_demand_max_price_percentage_over_lowest_price}.
	// Experimental.
	OnDemandMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"onDemandMaxPricePercentageOverLowestPrice" yaml:"onDemandMaxPricePercentageOverLowestPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#require_hibernate_support AwsCapacityProvider#require_hibernate_support}.
	// Experimental.
	RequireHibernateSupport interface{} `field:"optional" json:"requireHibernateSupport" yaml:"requireHibernateSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#spot_max_price_percentage_over_lowest_price AwsCapacityProvider#spot_max_price_percentage_over_lowest_price}.
	// Experimental.
	SpotMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"spotMaxPricePercentageOverLowestPrice" yaml:"spotMaxPricePercentageOverLowestPrice"`
	// total_local_storage_gb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#total_local_storage_gb AwsCapacityProvider#total_local_storage_gb}
	// Experimental.
	TotalLocalStorageGb *AwsCapacityProvider_TotalLocalStorageGbProperty `field:"optional" json:"totalLocalStorageGb" yaml:"totalLocalStorageGb"`
}

