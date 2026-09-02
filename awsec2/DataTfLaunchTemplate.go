package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/launch_template aws_launch_template}.
// Experimental.
type DataTfLaunchTemplate interface {
	cdktn.TerraformDataSource
	// Experimental.
	Arn() *string
	// Experimental.
	BlockDeviceMappings() DataTfLaunchTemplate_BlockDeviceMappingsPropertyList
	// Experimental.
	CapacityReservationSpecification() DataTfLaunchTemplate_CapacityReservationSpecificationPropertyList
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	CpuOptions() DataTfLaunchTemplate_CpuOptionsPropertyList
	// Experimental.
	CreditSpecification() DataTfLaunchTemplate_CreditSpecificationPropertyList
	// Experimental.
	DefaultVersion() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	Description() *string
	// Experimental.
	DisableApiStop() cdktn.IResolvable
	// Experimental.
	DisableApiTermination() cdktn.IResolvable
	// Experimental.
	EbsOptimized() *string
	// Experimental.
	EnclaveOptions() DataTfLaunchTemplate_EnclaveOptionsPropertyList
	// Experimental.
	Filter() DataTfLaunchTemplate_FilterPropertyList
	// Experimental.
	FilterInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HibernationOptions() DataTfLaunchTemplate_HibernationOptionsPropertyList
	// Experimental.
	IamInstanceProfile() DataTfLaunchTemplate_IamInstanceProfilePropertyList
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	ImageId() *string
	// Experimental.
	InstanceInitiatedShutdownBehavior() *string
	// Experimental.
	InstanceMarketOptions() DataTfLaunchTemplate_InstanceMarketOptionsPropertyList
	// Experimental.
	InstanceRequirements() DataTfLaunchTemplate_InstanceRequirementsPropertyList
	// Experimental.
	InstanceType() *string
	// Experimental.
	KernelId() *string
	// Experimental.
	KeyName() *string
	// Experimental.
	LatestVersion() *float64
	// Experimental.
	LicenseSpecification() DataTfLaunchTemplate_LicenseSpecificationPropertyList
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MaintenanceOptions() DataTfLaunchTemplate_MaintenanceOptionsPropertyList
	// Experimental.
	MetadataOptions() DataTfLaunchTemplate_MetadataOptionsPropertyList
	// Experimental.
	Monitoring() DataTfLaunchTemplate_MonitoringPropertyList
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NetworkInterfaces() DataTfLaunchTemplate_NetworkInterfacesPropertyList
	// Experimental.
	NetworkPerformanceOptions() DataTfLaunchTemplate_NetworkPerformanceOptionsPropertyList
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Placement() DataTfLaunchTemplate_PlacementPropertyList
	// Experimental.
	PrivateDnsNameOptions() DataTfLaunchTemplate_PrivateDnsNameOptionsPropertyList
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RamDiskId() *string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	SecondaryInterfaces() DataTfLaunchTemplate_SecondaryInterfacesPropertyList
	// Experimental.
	SecurityGroupNames() *[]*string
	// Experimental.
	Tags() *map[string]*string
	// Experimental.
	SetTags(val *map[string]*string)
	// Experimental.
	TagsInput() *map[string]*string
	// Experimental.
	TagSpecifications() DataTfLaunchTemplate_TagSpecificationsPropertyList
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() DataTfLaunchTemplate_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	UserData() *string
	// Experimental.
	VpcSecurityGroupIds() *[]*string
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
	PutFilter(value interface{})
	// Experimental.
	PutTimeouts(value *DataTfLaunchTemplate_TimeoutsProperty)
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
	ResetFilter()
	// Experimental.
	ResetId()
	// Experimental.
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetTags()
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

// The jsii proxy struct for DataTfLaunchTemplate
type jsiiProxy_DataTfLaunchTemplate struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataTfLaunchTemplate) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) BlockDeviceMappings() DataTfLaunchTemplate_BlockDeviceMappingsPropertyList {
	var returns DataTfLaunchTemplate_BlockDeviceMappingsPropertyList
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) CapacityReservationSpecification() DataTfLaunchTemplate_CapacityReservationSpecificationPropertyList {
	var returns DataTfLaunchTemplate_CapacityReservationSpecificationPropertyList
	_jsii_.Get(
		j,
		"capacityReservationSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) CpuOptions() DataTfLaunchTemplate_CpuOptionsPropertyList {
	var returns DataTfLaunchTemplate_CpuOptionsPropertyList
	_jsii_.Get(
		j,
		"cpuOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) CreditSpecification() DataTfLaunchTemplate_CreditSpecificationPropertyList {
	var returns DataTfLaunchTemplate_CreditSpecificationPropertyList
	_jsii_.Get(
		j,
		"creditSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) DefaultVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) DisableApiStop() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"disableApiStop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) DisableApiTermination() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"disableApiTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) EbsOptimized() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) EnclaveOptions() DataTfLaunchTemplate_EnclaveOptionsPropertyList {
	var returns DataTfLaunchTemplate_EnclaveOptionsPropertyList
	_jsii_.Get(
		j,
		"enclaveOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) Filter() DataTfLaunchTemplate_FilterPropertyList {
	var returns DataTfLaunchTemplate_FilterPropertyList
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) HibernationOptions() DataTfLaunchTemplate_HibernationOptionsPropertyList {
	var returns DataTfLaunchTemplate_HibernationOptionsPropertyList
	_jsii_.Get(
		j,
		"hibernationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) IamInstanceProfile() DataTfLaunchTemplate_IamInstanceProfilePropertyList {
	var returns DataTfLaunchTemplate_IamInstanceProfilePropertyList
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) InstanceInitiatedShutdownBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) InstanceMarketOptions() DataTfLaunchTemplate_InstanceMarketOptionsPropertyList {
	var returns DataTfLaunchTemplate_InstanceMarketOptionsPropertyList
	_jsii_.Get(
		j,
		"instanceMarketOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) InstanceRequirements() DataTfLaunchTemplate_InstanceRequirementsPropertyList {
	var returns DataTfLaunchTemplate_InstanceRequirementsPropertyList
	_jsii_.Get(
		j,
		"instanceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) KernelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) LatestVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) LicenseSpecification() DataTfLaunchTemplate_LicenseSpecificationPropertyList {
	var returns DataTfLaunchTemplate_LicenseSpecificationPropertyList
	_jsii_.Get(
		j,
		"licenseSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) MaintenanceOptions() DataTfLaunchTemplate_MaintenanceOptionsPropertyList {
	var returns DataTfLaunchTemplate_MaintenanceOptionsPropertyList
	_jsii_.Get(
		j,
		"maintenanceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) MetadataOptions() DataTfLaunchTemplate_MetadataOptionsPropertyList {
	var returns DataTfLaunchTemplate_MetadataOptionsPropertyList
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) Monitoring() DataTfLaunchTemplate_MonitoringPropertyList {
	var returns DataTfLaunchTemplate_MonitoringPropertyList
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) NetworkInterfaces() DataTfLaunchTemplate_NetworkInterfacesPropertyList {
	var returns DataTfLaunchTemplate_NetworkInterfacesPropertyList
	_jsii_.Get(
		j,
		"networkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) NetworkPerformanceOptions() DataTfLaunchTemplate_NetworkPerformanceOptionsPropertyList {
	var returns DataTfLaunchTemplate_NetworkPerformanceOptionsPropertyList
	_jsii_.Get(
		j,
		"networkPerformanceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) Placement() DataTfLaunchTemplate_PlacementPropertyList {
	var returns DataTfLaunchTemplate_PlacementPropertyList
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) PrivateDnsNameOptions() DataTfLaunchTemplate_PrivateDnsNameOptionsPropertyList {
	var returns DataTfLaunchTemplate_PrivateDnsNameOptionsPropertyList
	_jsii_.Get(
		j,
		"privateDnsNameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) RamDiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramDiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) SecondaryInterfaces() DataTfLaunchTemplate_SecondaryInterfacesPropertyList {
	var returns DataTfLaunchTemplate_SecondaryInterfacesPropertyList
	_jsii_.Get(
		j,
		"secondaryInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) SecurityGroupNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) TagSpecifications() DataTfLaunchTemplate_TagSpecificationsPropertyList {
	var returns DataTfLaunchTemplate_TagSpecificationsPropertyList
	_jsii_.Get(
		j,
		"tagSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) Timeouts() DataTfLaunchTemplate_TimeoutsPropertyOutputReference {
	var returns DataTfLaunchTemplate_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfLaunchTemplate) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/launch_template aws_launch_template} Data Source.
// Experimental.
func NewDataTfLaunchTemplate(scope constructs.Construct, id *string, config *DataTfLaunchTemplateConfig) DataTfLaunchTemplate {
	_init_.Initialize()

	if err := validateNewDataTfLaunchTemplateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfLaunchTemplate{}

	_jsii_.Create(
		"@cdktn/aws-ec2.DataTfLaunchTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/launch_template aws_launch_template} Data Source.
// Experimental.
func NewDataTfLaunchTemplate_Override(d DataTfLaunchTemplate, scope constructs.Construct, id *string, config *DataTfLaunchTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.DataTfLaunchTemplate",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataTfLaunchTemplate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataTfLaunchTemplate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataTfLaunchTemplate)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataTfLaunchTemplate)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataTfLaunchTemplate)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataTfLaunchTemplate)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataTfLaunchTemplate)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataTfLaunchTemplate)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataTfLaunchTemplate)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTN code for importing a DataTfLaunchTemplate resource upon running "cdktn plan <stack-name>".
// Experimental.
func DataTfLaunchTemplate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataTfLaunchTemplate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.DataTfLaunchTemplate",
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
func DataTfLaunchTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfLaunchTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.DataTfLaunchTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataTfLaunchTemplate_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfLaunchTemplate_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.DataTfLaunchTemplate",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataTfLaunchTemplate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfLaunchTemplate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.DataTfLaunchTemplate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTfLaunchTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-ec2.DataTfLaunchTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTfLaunchTemplate) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTfLaunchTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataTfLaunchTemplate) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfLaunchTemplate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataTfLaunchTemplate) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataTfLaunchTemplate) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataTfLaunchTemplate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataTfLaunchTemplate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataTfLaunchTemplate) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataTfLaunchTemplate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataTfLaunchTemplate) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfLaunchTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTfLaunchTemplate) PutFilter(value interface{}) {
	if err := d.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFilter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfLaunchTemplate) PutTimeouts(value *DataTfLaunchTemplate_TimeoutsProperty) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfLaunchTemplate) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataTfLaunchTemplate) ResetFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfLaunchTemplate) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfLaunchTemplate) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfLaunchTemplate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfLaunchTemplate) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfLaunchTemplate) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfLaunchTemplate) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfLaunchTemplate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfLaunchTemplate) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfLaunchTemplate) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfLaunchTemplate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfLaunchTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfLaunchTemplate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfLaunchTemplate) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

