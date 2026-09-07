package oracledatabaseaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/oracledatabaseaws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/oracledatabaseaws/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/odb_cloud_autonomous_vm_cluster aws_odb_cloud_autonomous_vm_cluster}.
// Experimental.
type DataAwsCloudAutonomousVmCluster interface {
	cdktn.TerraformDataSource
	// Experimental.
	Arn() *string
	// Experimental.
	AutonomousDataStoragePercentage() *float64
	// Experimental.
	AutonomousDataStorageSizeInTbs() *float64
	// Experimental.
	AvailableAutonomousDataStorageSizeInTbs() *float64
	// Experimental.
	AvailableContainerDatabases() *float64
	// Experimental.
	AvailableCpus() *float64
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CloudExadataInfrastructureArn() *string
	// Experimental.
	CloudExadataInfrastructureId() *string
	// Experimental.
	ComputeModel() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	CpuCoreCount() *float64
	// Experimental.
	CpuCoreCountPerNode() *float64
	// Experimental.
	CpuPercentage() *float64
	// Experimental.
	CreatedAt() *string
	// Experimental.
	DataStorageSizeInGbs() *float64
	// Experimental.
	DataStorageSizeInTbs() *float64
	// Experimental.
	DbServers() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	Description() *string
	// Experimental.
	DisplayName() *string
	// Experimental.
	Domain() *string
	// Experimental.
	ExadataStorageInTbsLowestScaledValue() *float64
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Hostname() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	IsMtlsEnabledVmCluster() cdktn.IResolvable
	// Experimental.
	LicenseModel() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MaintenanceWindow() DataAwsCloudAutonomousVmCluster_MaintenanceWindowPropertyList
	// Experimental.
	MaxAcdsLowestScaledValue() *float64
	// Experimental.
	MemoryPerOracleComputeUnitInGbs() *float64
	// Experimental.
	MemorySizeInGbs() *float64
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	NodeCount() *float64
	// Experimental.
	NonProvisionableAutonomousContainerDatabases() *float64
	// Experimental.
	Ocid() *string
	// Experimental.
	OciResourceAnchorName() *string
	// Experimental.
	OciUrl() *string
	// Experimental.
	OdbNetworkArn() *string
	// Experimental.
	OdbNetworkId() *string
	// Experimental.
	OdbNodeStorageSizeInGbs() *float64
	// Experimental.
	PercentProgress() *float64
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	ProvisionableAutonomousContainerDatabases() *float64
	// Experimental.
	ProvisionedAutonomousContainerDatabases() *float64
	// Experimental.
	ProvisionedCpus() *float64
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	ReclaimableCpus() *float64
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	ReservedCpus() *float64
	// Experimental.
	ScanListenerPortNonTls() *float64
	// Experimental.
	ScanListenerPortTls() *float64
	// Experimental.
	Shape() *string
	// Experimental.
	Status() *string
	// Experimental.
	StatusReason() *string
	// Experimental.
	Tags() cdktn.StringMap
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	TimeDatabaseSslCertificateExpires() *string
	// Experimental.
	TimeOrdsCertificateExpires() *string
	// Experimental.
	TimeZone() *string
	// Experimental.
	TotalContainerDatabases() *float64
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
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

// The jsii proxy struct for DataAwsCloudAutonomousVmCluster
type jsiiProxy_DataAwsCloudAutonomousVmCluster struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) AutonomousDataStoragePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autonomousDataStoragePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) AutonomousDataStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autonomousDataStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) AvailableAutonomousDataStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableAutonomousDataStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) AvailableContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) AvailableCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) CloudExadataInfrastructureArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) CloudExadataInfrastructureId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) ComputeModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) CpuCoreCountPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCountPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) CpuPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) DataStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) DataStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) DbServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) ExadataStorageInTbsLowestScaledValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"exadataStorageInTbsLowestScaledValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) IsMtlsEnabledVmCluster() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isMtlsEnabledVmCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) MaintenanceWindow() DataAwsCloudAutonomousVmCluster_MaintenanceWindowPropertyList {
	var returns DataAwsCloudAutonomousVmCluster_MaintenanceWindowPropertyList
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) MaxAcdsLowestScaledValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAcdsLowestScaledValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) MemoryPerOracleComputeUnitInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerOracleComputeUnitInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) MemorySizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) NonProvisionableAutonomousContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nonProvisionableAutonomousContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) OciResourceAnchorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociResourceAnchorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) OdbNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) OdbNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) OdbNodeStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"odbNodeStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) PercentProgress() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentProgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) ProvisionableAutonomousContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionableAutonomousContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) ProvisionedAutonomousContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedAutonomousContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) ProvisionedCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) ReclaimableCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reclaimableCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) ReservedCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) ScanListenerPortNonTls() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortNonTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) ScanListenerPortTls() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) Tags() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) TimeDatabaseSslCertificateExpires() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeDatabaseSslCertificateExpires",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) TimeOrdsCertificateExpires() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOrdsCertificateExpires",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster) TotalContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalContainerDatabases",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/odb_cloud_autonomous_vm_cluster aws_odb_cloud_autonomous_vm_cluster} Data Source.
// Experimental.
func NewDataAwsCloudAutonomousVmCluster(scope constructs.Construct, id *string, config *DataAwsCloudAutonomousVmClusterConfig) DataAwsCloudAutonomousVmCluster {
	_init_.Initialize()

	if err := validateNewDataAwsCloudAutonomousVmClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsCloudAutonomousVmCluster{}

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.DataAwsCloudAutonomousVmCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/odb_cloud_autonomous_vm_cluster aws_odb_cloud_autonomous_vm_cluster} Data Source.
// Experimental.
func NewDataAwsCloudAutonomousVmCluster_Override(d DataAwsCloudAutonomousVmCluster, scope constructs.Construct, id *string, config *DataAwsCloudAutonomousVmClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.DataAwsCloudAutonomousVmCluster",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudAutonomousVmCluster)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTN code for importing a DataAwsCloudAutonomousVmCluster resource upon running "cdktn plan <stack-name>".
// Experimental.
func DataAwsCloudAutonomousVmCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsCloudAutonomousVmCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.DataAwsCloudAutonomousVmCluster",
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
func DataAwsCloudAutonomousVmCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsCloudAutonomousVmCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.DataAwsCloudAutonomousVmCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsCloudAutonomousVmCluster_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsCloudAutonomousVmCluster_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.DataAwsCloudAutonomousVmCluster",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsCloudAutonomousVmCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsCloudAutonomousVmCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.DataAwsCloudAutonomousVmCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsCloudAutonomousVmCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-oracle-database-aws.DataAwsCloudAutonomousVmCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudAutonomousVmCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

