package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/instance aws_instance}.
// Experimental.
type DataTfInstance interface {
	cdktn.TerraformDataSource
	// Experimental.
	Ami() *string
	// Experimental.
	Arn() *string
	// Experimental.
	AssociatePublicIpAddress() cdktn.IResolvable
	// Experimental.
	AvailabilityZone() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	CreditSpecification() DataTfInstance_CreditSpecificationPropertyList
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DisableApiStop() cdktn.IResolvable
	// Experimental.
	DisableApiTermination() cdktn.IResolvable
	// Experimental.
	EbsBlockDevice() DataTfInstance_EbsBlockDevicePropertyList
	// Experimental.
	EbsOptimized() cdktn.IResolvable
	// Experimental.
	EnclaveOptions() DataTfInstance_EnclaveOptionsPropertyList
	// Experimental.
	EphemeralBlockDevice() DataTfInstance_EphemeralBlockDevicePropertyList
	// Experimental.
	FetchPasswordData() interface{}
	// Experimental.
	SetFetchPasswordData(val interface{})
	// Experimental.
	FetchPasswordDataInput() interface{}
	// Experimental.
	FetchUserData() interface{}
	// Experimental.
	SetFetchUserData(val interface{})
	// Experimental.
	FetchUserDataInput() interface{}
	// Experimental.
	Filter() DataTfInstance_FilterPropertyList
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
	HostId() *string
	// Experimental.
	HostResourceGroupArn() *string
	// Experimental.
	IamInstanceProfile() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	InstanceId() *string
	// Experimental.
	SetInstanceId(val *string)
	// Experimental.
	InstanceIdInput() *string
	// Experimental.
	InstanceState() *string
	// Experimental.
	InstanceTags() *map[string]*string
	// Experimental.
	SetInstanceTags(val *map[string]*string)
	// Experimental.
	InstanceTagsInput() *map[string]*string
	// Experimental.
	InstanceType() *string
	// Experimental.
	Ipv6Addresses() *[]*string
	// Experimental.
	KeyName() *string
	// Experimental.
	LaunchTime() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MaintenanceOptions() DataTfInstance_MaintenanceOptionsPropertyList
	// Experimental.
	MetadataOptions() DataTfInstance_MetadataOptionsPropertyList
	// Experimental.
	Monitoring() cdktn.IResolvable
	// Experimental.
	NetworkInterfaceId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OutpostArn() *string
	// Experimental.
	PasswordData() *string
	// Experimental.
	PlacementGroup() *string
	// Experimental.
	PlacementGroupId() *string
	// Experimental.
	PlacementPartitionNumber() *float64
	// Experimental.
	PrivateDns() *string
	// Experimental.
	PrivateDnsNameOptions() DataTfInstance_PrivateDnsNameOptionsPropertyList
	// Experimental.
	PrivateIp() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	PublicDns() *string
	// Experimental.
	PublicIp() *string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	RootBlockDevice() DataTfInstance_RootBlockDevicePropertyList
	// Experimental.
	SecondaryPrivateIps() *[]*string
	// Experimental.
	SecurityGroups() *[]*string
	// Experimental.
	SourceDestCheck() cdktn.IResolvable
	// Experimental.
	SubnetId() *string
	// Experimental.
	Tags() *map[string]*string
	// Experimental.
	SetTags(val *map[string]*string)
	// Experimental.
	TagsInput() *map[string]*string
	// Experimental.
	Tenancy() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() DataTfInstance_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	UserData() *string
	// Experimental.
	UserDataBase64() *string
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
	PutTimeouts(value *DataTfInstance_TimeoutsProperty)
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
	ResetFetchPasswordData()
	// Experimental.
	ResetFetchUserData()
	// Experimental.
	ResetFilter()
	// Experimental.
	ResetId()
	// Experimental.
	ResetInstanceId()
	// Experimental.
	ResetInstanceTags()
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

// The jsii proxy struct for DataTfInstance
type jsiiProxy_DataTfInstance struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataTfInstance) Ami() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ami",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) AssociatePublicIpAddress() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) CreditSpecification() DataTfInstance_CreditSpecificationPropertyList {
	var returns DataTfInstance_CreditSpecificationPropertyList
	_jsii_.Get(
		j,
		"creditSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) DisableApiStop() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"disableApiStop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) DisableApiTermination() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"disableApiTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) EbsBlockDevice() DataTfInstance_EbsBlockDevicePropertyList {
	var returns DataTfInstance_EbsBlockDevicePropertyList
	_jsii_.Get(
		j,
		"ebsBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) EbsOptimized() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) EnclaveOptions() DataTfInstance_EnclaveOptionsPropertyList {
	var returns DataTfInstance_EnclaveOptionsPropertyList
	_jsii_.Get(
		j,
		"enclaveOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) EphemeralBlockDevice() DataTfInstance_EphemeralBlockDevicePropertyList {
	var returns DataTfInstance_EphemeralBlockDevicePropertyList
	_jsii_.Get(
		j,
		"ephemeralBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) FetchPasswordData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchPasswordData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) FetchPasswordDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchPasswordDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) FetchUserData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchUserData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) FetchUserDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchUserDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) Filter() DataTfInstance_FilterPropertyList {
	var returns DataTfInstance_FilterPropertyList
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) HostId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) HostResourceGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostResourceGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) InstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) InstanceState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) InstanceTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"instanceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) InstanceTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"instanceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) Ipv6Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) LaunchTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) MaintenanceOptions() DataTfInstance_MaintenanceOptionsPropertyList {
	var returns DataTfInstance_MaintenanceOptionsPropertyList
	_jsii_.Get(
		j,
		"maintenanceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) MetadataOptions() DataTfInstance_MetadataOptionsPropertyList {
	var returns DataTfInstance_MetadataOptionsPropertyList
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) Monitoring() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) NetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) OutpostArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) PasswordData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) PlacementGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) PlacementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) PlacementPartitionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"placementPartitionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) PrivateDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) PrivateDnsNameOptions() DataTfInstance_PrivateDnsNameOptionsPropertyList {
	var returns DataTfInstance_PrivateDnsNameOptionsPropertyList
	_jsii_.Get(
		j,
		"privateDnsNameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) PublicDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) PublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) RootBlockDevice() DataTfInstance_RootBlockDevicePropertyList {
	var returns DataTfInstance_RootBlockDevicePropertyList
	_jsii_.Get(
		j,
		"rootBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) SecondaryPrivateIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secondaryPrivateIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) SourceDestCheck() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"sourceDestCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) Tenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) Timeouts() DataTfInstance_TimeoutsPropertyOutputReference {
	var returns DataTfInstance_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) UserDataBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfInstance) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/instance aws_instance} Data Source.
// Experimental.
func NewDataTfInstance(scope constructs.Construct, id *string, config *DataTfInstanceConfig) DataTfInstance {
	_init_.Initialize()

	if err := validateNewDataTfInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfInstance{}

	_jsii_.Create(
		"@cdktn/aws-ec2.DataTfInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/instance aws_instance} Data Source.
// Experimental.
func NewDataTfInstance_Override(d DataTfInstance, scope constructs.Construct, id *string, config *DataTfInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.DataTfInstance",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataTfInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataTfInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataTfInstance)SetFetchPasswordData(val interface{}) {
	if err := j.validateSetFetchPasswordDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fetchPasswordData",
		val,
	)
}

func (j *jsiiProxy_DataTfInstance)SetFetchUserData(val interface{}) {
	if err := j.validateSetFetchUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fetchUserData",
		val,
	)
}

func (j *jsiiProxy_DataTfInstance)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataTfInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataTfInstance)SetInstanceId(val *string) {
	if err := j.validateSetInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_DataTfInstance)SetInstanceTags(val *map[string]*string) {
	if err := j.validateSetInstanceTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTags",
		val,
	)
}

func (j *jsiiProxy_DataTfInstance)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataTfInstance)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataTfInstance)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataTfInstance)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTN code for importing a DataTfInstance resource upon running "cdktn plan <stack-name>".
// Experimental.
func DataTfInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataTfInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.DataTfInstance",
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
func DataTfInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.DataTfInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataTfInstance_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfInstance_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.DataTfInstance",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataTfInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.DataTfInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTfInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-ec2.DataTfInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTfInstance) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTfInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataTfInstance) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataTfInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataTfInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataTfInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataTfInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataTfInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataTfInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataTfInstance) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfInstance) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTfInstance) PutFilter(value interface{}) {
	if err := d.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFilter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfInstance) PutTimeouts(value *DataTfInstance_TimeoutsProperty) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfInstance) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataTfInstance) ResetFetchPasswordData() {
	_jsii_.InvokeVoid(
		d,
		"resetFetchPasswordData",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfInstance) ResetFetchUserData() {
	_jsii_.InvokeVoid(
		d,
		"resetFetchUserData",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfInstance) ResetFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfInstance) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfInstance) ResetInstanceId() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfInstance) ResetInstanceTags() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfInstance) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfInstance) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfInstance) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

