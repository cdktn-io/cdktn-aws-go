package awsec2


// Experimental.
type TfFleet_InstanceRequirementsProperty struct {
	// memory_mib block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#memory_mib TfFleet#memory_mib}
	// Experimental.
	MemoryMib *TfFleet_MemoryMibProperty `field:"required" json:"memoryMib" yaml:"memoryMib"`
	// vcpu_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#vcpu_count TfFleet#vcpu_count}
	// Experimental.
	VcpuCount *TfFleet_VcpuCountProperty `field:"required" json:"vcpuCount" yaml:"vcpuCount"`
	// accelerator_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#accelerator_count TfFleet#accelerator_count}
	// Experimental.
	AcceleratorCount *TfFleet_AcceleratorCountProperty `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#accelerator_manufacturers TfFleet#accelerator_manufacturers}.
	// Experimental.
	AcceleratorManufacturers *[]*string `field:"optional" json:"acceleratorManufacturers" yaml:"acceleratorManufacturers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#accelerator_names TfFleet#accelerator_names}.
	// Experimental.
	AcceleratorNames *[]*string `field:"optional" json:"acceleratorNames" yaml:"acceleratorNames"`
	// accelerator_total_memory_mib block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#accelerator_total_memory_mib TfFleet#accelerator_total_memory_mib}
	// Experimental.
	AcceleratorTotalMemoryMib *TfFleet_AcceleratorTotalMemoryMibProperty `field:"optional" json:"acceleratorTotalMemoryMib" yaml:"acceleratorTotalMemoryMib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#accelerator_types TfFleet#accelerator_types}.
	// Experimental.
	AcceleratorTypes *[]*string `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#allowed_instance_types TfFleet#allowed_instance_types}.
	// Experimental.
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#bare_metal TfFleet#bare_metal}.
	// Experimental.
	BareMetal *string `field:"optional" json:"bareMetal" yaml:"bareMetal"`
	// baseline_ebs_bandwidth_mbps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#baseline_ebs_bandwidth_mbps TfFleet#baseline_ebs_bandwidth_mbps}
	// Experimental.
	BaselineEbsBandwidthMbps *TfFleet_BaselineEbsBandwidthMbpsProperty `field:"optional" json:"baselineEbsBandwidthMbps" yaml:"baselineEbsBandwidthMbps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#burstable_performance TfFleet#burstable_performance}.
	// Experimental.
	BurstablePerformance *string `field:"optional" json:"burstablePerformance" yaml:"burstablePerformance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#cpu_manufacturers TfFleet#cpu_manufacturers}.
	// Experimental.
	CpuManufacturers *[]*string `field:"optional" json:"cpuManufacturers" yaml:"cpuManufacturers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#excluded_instance_types TfFleet#excluded_instance_types}.
	// Experimental.
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#instance_generations TfFleet#instance_generations}.
	// Experimental.
	InstanceGenerations *[]*string `field:"optional" json:"instanceGenerations" yaml:"instanceGenerations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#local_storage TfFleet#local_storage}.
	// Experimental.
	LocalStorage *string `field:"optional" json:"localStorage" yaml:"localStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#local_storage_types TfFleet#local_storage_types}.
	// Experimental.
	LocalStorageTypes *[]*string `field:"optional" json:"localStorageTypes" yaml:"localStorageTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#max_spot_price_as_percentage_of_optimal_on_demand_price TfFleet#max_spot_price_as_percentage_of_optimal_on_demand_price}.
	// Experimental.
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice *float64 `field:"optional" json:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice" yaml:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice"`
	// memory_gib_per_vcpu block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#memory_gib_per_vcpu TfFleet#memory_gib_per_vcpu}
	// Experimental.
	MemoryGibPerVcpu *TfFleet_MemoryGibPerVcpuProperty `field:"optional" json:"memoryGibPerVcpu" yaml:"memoryGibPerVcpu"`
	// network_bandwidth_gbps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#network_bandwidth_gbps TfFleet#network_bandwidth_gbps}
	// Experimental.
	NetworkBandwidthGbps *TfFleet_NetworkBandwidthGbpsProperty `field:"optional" json:"networkBandwidthGbps" yaml:"networkBandwidthGbps"`
	// network_interface_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#network_interface_count TfFleet#network_interface_count}
	// Experimental.
	NetworkInterfaceCount *TfFleet_NetworkInterfaceCountProperty `field:"optional" json:"networkInterfaceCount" yaml:"networkInterfaceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#on_demand_max_price_percentage_over_lowest_price TfFleet#on_demand_max_price_percentage_over_lowest_price}.
	// Experimental.
	OnDemandMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"onDemandMaxPricePercentageOverLowestPrice" yaml:"onDemandMaxPricePercentageOverLowestPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#require_hibernate_support TfFleet#require_hibernate_support}.
	// Experimental.
	RequireHibernateSupport interface{} `field:"optional" json:"requireHibernateSupport" yaml:"requireHibernateSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#spot_max_price_percentage_over_lowest_price TfFleet#spot_max_price_percentage_over_lowest_price}.
	// Experimental.
	SpotMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"spotMaxPricePercentageOverLowestPrice" yaml:"spotMaxPricePercentageOverLowestPrice"`
	// total_local_storage_gb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#total_local_storage_gb TfFleet#total_local_storage_gb}
	// Experimental.
	TotalLocalStorageGb *TfFleet_TotalLocalStorageGbProperty `field:"optional" json:"totalLocalStorageGb" yaml:"totalLocalStorageGb"`
}

