package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_instance_type aws_ec2_instance_type}.
// Experimental.
type DataTfInstanceType interface {
	cdktn.TerraformDataSource
	// Experimental.
	AutoRecoverySupported() cdktn.IResolvable
	// Experimental.
	BandwidthWeightings() *[]*string
	// Experimental.
	BareMetal() cdktn.IResolvable
	// Experimental.
	BootModes() *[]*string
	// Experimental.
	BurstablePerformanceSupported() cdktn.IResolvable
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	CurrentGeneration() cdktn.IResolvable
	// Experimental.
	DedicatedHostsSupported() cdktn.IResolvable
	// Experimental.
	DefaultCores() *float64
	// Experimental.
	DefaultNetworkCardIndex() *float64
	// Experimental.
	DefaultThreadsPerCore() *float64
	// Experimental.
	DefaultVcpus() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EbsEncryptionSupport() *string
	// Experimental.
	EbsNvmeSupport() *string
	// Experimental.
	EbsOptimizedSupport() *string
	// Experimental.
	EbsPerformanceBaselineBandwidth() *float64
	// Experimental.
	EbsPerformanceBaselineIops() *float64
	// Experimental.
	EbsPerformanceBaselineThroughput() *float64
	// Experimental.
	EbsPerformanceMaximumBandwidth() *float64
	// Experimental.
	EbsPerformanceMaximumIops() *float64
	// Experimental.
	EbsPerformanceMaximumThroughput() *float64
	// Experimental.
	EfaMaximumInterfaces() *float64
	// Experimental.
	EfaSupported() cdktn.IResolvable
	// Experimental.
	EnaSrdSupported() cdktn.IResolvable
	// Experimental.
	EnaSupport() *string
	// Experimental.
	EncryptionInTransitSupported() cdktn.IResolvable
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fpgas() DataTfInstanceType_FpgasPropertyList
	// Experimental.
	Fqn() *string
	// Experimental.
	FreeTierEligible() cdktn.IResolvable
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Gpus() DataTfInstanceType_GpusPropertyList
	// Experimental.
	HibernationSupported() cdktn.IResolvable
	// Experimental.
	Hypervisor() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	InferenceAccelerators() DataTfInstanceType_InferenceAcceleratorsPropertyList
	// Experimental.
	InstanceDisks() DataTfInstanceType_InstanceDisksPropertyList
	// Experimental.
	InstanceStorageSupported() cdktn.IResolvable
	// Experimental.
	InstanceType() *string
	// Experimental.
	SetInstanceType(val *string)
	// Experimental.
	InstanceTypeInput() *string
	// Experimental.
	Ipv6Supported() cdktn.IResolvable
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MaximumIpv4AddressesPerInterface() *float64
	// Experimental.
	MaximumIpv6AddressesPerInterface() *float64
	// Experimental.
	MaximumNetworkCards() *float64
	// Experimental.
	MaximumNetworkInterfaces() *float64
	// Experimental.
	MediaAccelerators() DataTfInstanceType_MediaAcceleratorsPropertyList
	// Experimental.
	MemorySize() *float64
	// Experimental.
	NetworkCards() DataTfInstanceType_NetworkCardsPropertyList
	// Experimental.
	NetworkPerformance() *string
	// Experimental.
	NeuronDevices() DataTfInstanceType_NeuronDevicesPropertyList
	// Experimental.
	NitroEnclavesSupport() *string
	// Experimental.
	NitroTpmSupport() *string
	// Experimental.
	NitroTpmSupportedVersions() *[]*string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PhcSupport() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	SupportedArchitectures() *[]*string
	// Experimental.
	SupportedCpuFeatures() *[]*string
	// Experimental.
	SupportedPlacementStrategies() *[]*string
	// Experimental.
	SupportedRootDeviceTypes() *[]*string
	// Experimental.
	SupportedUsagesClasses() *[]*string
	// Experimental.
	SupportedVirtualizationTypes() *[]*string
	// Experimental.
	SustainedClockSpeed() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() DataTfInstanceType_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TotalFpgaMemory() *float64
	// Experimental.
	TotalGpuMemory() *float64
	// Experimental.
	TotalInferenceMemory() *float64
	// Experimental.
	TotalInstanceStorage() *float64
	// Experimental.
	TotalMediaMemory() *float64
	// Experimental.
	TotalNeuronDeviceMemory() *float64
	// Experimental.
	ValidCores() *[]*float64
	// Experimental.
	ValidThreadsPerCore() *[]*float64
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Experimental.
	PutTimeouts(value *DataTfInstanceType_TimeoutsProperty)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	// Experimental.
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for DataTfInstanceType
type jsiiProxy_DataTfInstanceType struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataTfInstanceType) AutoRecoverySupported() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"autoRecoverySupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) BandwidthWeightings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bandwidthWeightings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) BareMetal() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"bareMetal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) BootModes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bootModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) BurstablePerformanceSupported() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"burstablePerformanceSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) CurrentGeneration() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"currentGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) DedicatedHostsSupported() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"dedicatedHostsSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) DefaultCores() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultCores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) DefaultNetworkCardIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultNetworkCardIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) DefaultThreadsPerCore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultThreadsPerCore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) DefaultVcpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultVcpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) EbsEncryptionSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryptionSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) EbsNvmeSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsNvmeSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) EbsOptimizedSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsOptimizedSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) EbsPerformanceBaselineBandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsPerformanceBaselineBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) EbsPerformanceBaselineIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsPerformanceBaselineIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) EbsPerformanceBaselineThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsPerformanceBaselineThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) EbsPerformanceMaximumBandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsPerformanceMaximumBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) EbsPerformanceMaximumIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsPerformanceMaximumIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) EbsPerformanceMaximumThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsPerformanceMaximumThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) EfaMaximumInterfaces() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"efaMaximumInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) EfaSupported() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"efaSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) EnaSrdSupported() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enaSrdSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) EnaSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enaSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) EncryptionInTransitSupported() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"encryptionInTransitSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) Fpgas() DataTfInstanceType_FpgasPropertyList {
	var returns DataTfInstanceType_FpgasPropertyList
	_jsii_.Get(
		j,
		"fpgas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) FreeTierEligible() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"freeTierEligible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) Gpus() DataTfInstanceType_GpusPropertyList {
	var returns DataTfInstanceType_GpusPropertyList
	_jsii_.Get(
		j,
		"gpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) HibernationSupported() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"hibernationSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) Hypervisor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hypervisor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) InferenceAccelerators() DataTfInstanceType_InferenceAcceleratorsPropertyList {
	var returns DataTfInstanceType_InferenceAcceleratorsPropertyList
	_jsii_.Get(
		j,
		"inferenceAccelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) InstanceDisks() DataTfInstanceType_InstanceDisksPropertyList {
	var returns DataTfInstanceType_InstanceDisksPropertyList
	_jsii_.Get(
		j,
		"instanceDisks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) InstanceStorageSupported() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"instanceStorageSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) Ipv6Supported() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"ipv6Supported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) MaximumIpv4AddressesPerInterface() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumIpv4AddressesPerInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) MaximumIpv6AddressesPerInterface() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumIpv6AddressesPerInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) MaximumNetworkCards() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumNetworkCards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) MaximumNetworkInterfaces() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumNetworkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) MediaAccelerators() DataTfInstanceType_MediaAcceleratorsPropertyList {
	var returns DataTfInstanceType_MediaAcceleratorsPropertyList
	_jsii_.Get(
		j,
		"mediaAccelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) MemorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) NetworkCards() DataTfInstanceType_NetworkCardsPropertyList {
	var returns DataTfInstanceType_NetworkCardsPropertyList
	_jsii_.Get(
		j,
		"networkCards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) NetworkPerformance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPerformance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) NeuronDevices() DataTfInstanceType_NeuronDevicesPropertyList {
	var returns DataTfInstanceType_NeuronDevicesPropertyList
	_jsii_.Get(
		j,
		"neuronDevices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) NitroEnclavesSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nitroEnclavesSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) NitroTpmSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nitroTpmSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) NitroTpmSupportedVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nitroTpmSupportedVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) PhcSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phcSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) SupportedArchitectures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedArchitectures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) SupportedCpuFeatures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedCpuFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) SupportedPlacementStrategies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedPlacementStrategies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) SupportedRootDeviceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedRootDeviceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) SupportedUsagesClasses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedUsagesClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) SupportedVirtualizationTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedVirtualizationTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) SustainedClockSpeed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sustainedClockSpeed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) Timeouts() DataTfInstanceType_TimeoutsPropertyOutputReference {
	var returns DataTfInstanceType_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) TotalFpgaMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalFpgaMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) TotalGpuMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalGpuMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) TotalInferenceMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalInferenceMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) TotalInstanceStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalInstanceStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) TotalMediaMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalMediaMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) TotalNeuronDeviceMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalNeuronDeviceMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) ValidCores() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"validCores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstanceType) ValidThreadsPerCore() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"validThreadsPerCore",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_instance_type aws_ec2_instance_type} Data Source.
// Experimental.
func NewDataTfInstanceType(scope constructs.Construct, id *string, config *DataTfInstanceTypeConfig) DataTfInstanceType {
	_init_.Initialize()

	if err := validateNewDataTfInstanceTypeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfInstanceType{}

	_jsii_.Create(
		"@cdktn/aws-ec2.DataTfInstanceType",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_instance_type aws_ec2_instance_type} Data Source.
// Experimental.
func NewDataTfInstanceType_Override(d DataTfInstanceType, scope constructs.Construct, id *string, config *DataTfInstanceTypeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.DataTfInstanceType",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataTfInstanceType)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataTfInstanceType)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataTfInstanceType)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataTfInstanceType)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataTfInstanceType)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_DataTfInstanceType)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataTfInstanceType)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataTfInstanceType)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTN code for importing a DataTfInstanceType resource upon running "cdktn plan <stack-name>".
// Experimental.
func DataTfInstanceType_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataTfInstanceType_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.DataTfInstanceType",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func DataTfInstanceType_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfInstanceType_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.DataTfInstanceType",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataTfInstanceType_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfInstanceType_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.DataTfInstanceType",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataTfInstanceType_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfInstanceType_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.DataTfInstanceType",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTfInstanceType_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-ec2.DataTfInstanceType",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTfInstanceType) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTfInstanceType) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstanceType) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstanceType) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstanceType) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstanceType) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstanceType) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstanceType) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstanceType) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstanceType) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstanceType) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstanceType) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTfInstanceType) PutTimeouts(value *DataTfInstanceType_TimeoutsProperty) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfInstanceType) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataTfInstanceType) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfInstanceType) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfInstanceType) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfInstanceType) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfInstanceType) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstanceType) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstanceType) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstanceType) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstanceType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstanceType) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstanceType) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

