package rds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/rds/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/rds/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance aws_rds_orderable_db_instance}.
// Experimental.
type DataAwsOrderableDbInstance interface {
	cdktn.TerraformDataSource
	// Experimental.
	AvailabilityZoneGroup() *string
	// Experimental.
	SetAvailabilityZoneGroup(val *string)
	// Experimental.
	AvailabilityZoneGroupInput() *string
	// Experimental.
	AvailabilityZones() *[]*string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	Engine() *string
	// Experimental.
	SetEngine(val *string)
	// Experimental.
	EngineInput() *string
	// Experimental.
	EngineLatestVersion() interface{}
	// Experimental.
	SetEngineLatestVersion(val interface{})
	// Experimental.
	EngineLatestVersionInput() interface{}
	// Experimental.
	EngineVersion() *string
	// Experimental.
	SetEngineVersion(val *string)
	// Experimental.
	EngineVersionInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	InstanceClass() *string
	// Experimental.
	SetInstanceClass(val *string)
	// Experimental.
	InstanceClassInput() *string
	// Experimental.
	LicenseModel() *string
	// Experimental.
	SetLicenseModel(val *string)
	// Experimental.
	LicenseModelInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MaxIopsPerDbInstance() *float64
	// Experimental.
	MaxIopsPerGib() *float64
	// Experimental.
	MaxStorageSize() *float64
	// Experimental.
	MinIopsPerDbInstance() *float64
	// Experimental.
	MinIopsPerGib() *float64
	// Experimental.
	MinStorageSize() *float64
	// Experimental.
	MultiAzCapable() cdktn.IResolvable
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OutpostCapable() cdktn.IResolvable
	// Experimental.
	PreferredEngineVersions() *[]*string
	// Experimental.
	SetPreferredEngineVersions(val *[]*string)
	// Experimental.
	PreferredEngineVersionsInput() *[]*string
	// Experimental.
	PreferredInstanceClasses() *[]*string
	// Experimental.
	SetPreferredInstanceClasses(val *[]*string)
	// Experimental.
	PreferredInstanceClassesInput() *[]*string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	ReadReplicaCapable() interface{}
	// Experimental.
	SetReadReplicaCapable(val interface{})
	// Experimental.
	ReadReplicaCapableInput() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	StorageType() *string
	// Experimental.
	SetStorageType(val *string)
	// Experimental.
	StorageTypeInput() *string
	// Experimental.
	SupportedEngineModes() *[]*string
	// Experimental.
	SetSupportedEngineModes(val *[]*string)
	// Experimental.
	SupportedEngineModesInput() *[]*string
	// Experimental.
	SupportedNetworkTypes() *[]*string
	// Experimental.
	SetSupportedNetworkTypes(val *[]*string)
	// Experimental.
	SupportedNetworkTypesInput() *[]*string
	// Experimental.
	SupportsClusters() interface{}
	// Experimental.
	SetSupportsClusters(val interface{})
	// Experimental.
	SupportsClustersInput() interface{}
	// Experimental.
	SupportsEnhancedMonitoring() interface{}
	// Experimental.
	SetSupportsEnhancedMonitoring(val interface{})
	// Experimental.
	SupportsEnhancedMonitoringInput() interface{}
	// Experimental.
	SupportsGlobalDatabases() interface{}
	// Experimental.
	SetSupportsGlobalDatabases(val interface{})
	// Experimental.
	SupportsGlobalDatabasesInput() interface{}
	// Experimental.
	SupportsIamDatabaseAuthentication() interface{}
	// Experimental.
	SetSupportsIamDatabaseAuthentication(val interface{})
	// Experimental.
	SupportsIamDatabaseAuthenticationInput() interface{}
	// Experimental.
	SupportsIops() interface{}
	// Experimental.
	SetSupportsIops(val interface{})
	// Experimental.
	SupportsIopsInput() interface{}
	// Experimental.
	SupportsKerberosAuthentication() interface{}
	// Experimental.
	SetSupportsKerberosAuthentication(val interface{})
	// Experimental.
	SupportsKerberosAuthenticationInput() interface{}
	// Experimental.
	SupportsMultiAz() interface{}
	// Experimental.
	SetSupportsMultiAz(val interface{})
	// Experimental.
	SupportsMultiAzInput() interface{}
	// Experimental.
	SupportsPerformanceInsights() interface{}
	// Experimental.
	SetSupportsPerformanceInsights(val interface{})
	// Experimental.
	SupportsPerformanceInsightsInput() interface{}
	// Experimental.
	SupportsStorageAutoscaling() interface{}
	// Experimental.
	SetSupportsStorageAutoscaling(val interface{})
	// Experimental.
	SupportsStorageAutoscalingInput() interface{}
	// Experimental.
	SupportsStorageEncryption() interface{}
	// Experimental.
	SetSupportsStorageEncryption(val interface{})
	// Experimental.
	SupportsStorageEncryptionInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Vpc() interface{}
	// Experimental.
	SetVpc(val interface{})
	// Experimental.
	VpcInput() interface{}
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
	// Experimental.
	ResetAvailabilityZoneGroup()
	// Experimental.
	ResetEngineLatestVersion()
	// Experimental.
	ResetEngineVersion()
	// Experimental.
	ResetId()
	// Experimental.
	ResetInstanceClass()
	// Experimental.
	ResetLicenseModel()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPreferredEngineVersions()
	// Experimental.
	ResetPreferredInstanceClasses()
	// Experimental.
	ResetReadReplicaCapable()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetStorageType()
	// Experimental.
	ResetSupportedEngineModes()
	// Experimental.
	ResetSupportedNetworkTypes()
	// Experimental.
	ResetSupportsClusters()
	// Experimental.
	ResetSupportsEnhancedMonitoring()
	// Experimental.
	ResetSupportsGlobalDatabases()
	// Experimental.
	ResetSupportsIamDatabaseAuthentication()
	// Experimental.
	ResetSupportsIops()
	// Experimental.
	ResetSupportsKerberosAuthentication()
	// Experimental.
	ResetSupportsMultiAz()
	// Experimental.
	ResetSupportsPerformanceInsights()
	// Experimental.
	ResetSupportsStorageAutoscaling()
	// Experimental.
	ResetSupportsStorageEncryption()
	// Experimental.
	ResetVpc()
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

// The jsii proxy struct for DataAwsOrderableDbInstance
type jsiiProxy_DataAwsOrderableDbInstance struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) AvailabilityZoneGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) AvailabilityZoneGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) EngineLatestVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"engineLatestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) EngineLatestVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"engineLatestVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) InstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) InstanceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) LicenseModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) MaxIopsPerDbInstance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIopsPerDbInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) MaxIopsPerGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIopsPerGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) MaxStorageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStorageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) MinIopsPerDbInstance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIopsPerDbInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) MinIopsPerGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIopsPerGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) MinStorageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minStorageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) MultiAzCapable() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"multiAzCapable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) OutpostCapable() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"outpostCapable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) PreferredEngineVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredEngineVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) PreferredEngineVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredEngineVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) PreferredInstanceClasses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredInstanceClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) PreferredInstanceClassesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredInstanceClassesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) ReadReplicaCapable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readReplicaCapable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) ReadReplicaCapableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readReplicaCapableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportedEngineModes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedEngineModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportedEngineModesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedEngineModesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportedNetworkTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedNetworkTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportedNetworkTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedNetworkTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportsClusters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsClusters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportsClustersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsClustersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportsEnhancedMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsEnhancedMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportsEnhancedMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsEnhancedMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportsGlobalDatabases() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsGlobalDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportsGlobalDatabasesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsGlobalDatabasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportsIamDatabaseAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsIamDatabaseAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportsIamDatabaseAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsIamDatabaseAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportsIops() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportsIopsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsIopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportsKerberosAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsKerberosAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportsKerberosAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsKerberosAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportsMultiAz() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsMultiAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportsMultiAzInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsMultiAzInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportsPerformanceInsights() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsPerformanceInsights",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportsPerformanceInsightsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsPerformanceInsightsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportsStorageAutoscaling() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsStorageAutoscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportsStorageAutoscalingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsStorageAutoscalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportsStorageEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsStorageEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) SupportsStorageEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsStorageEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) Vpc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrderableDbInstance) VpcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance aws_rds_orderable_db_instance} Data Source.
// Experimental.
func NewDataAwsOrderableDbInstance(scope constructs.Construct, id *string, config *DataAwsOrderableDbInstanceConfig) DataAwsOrderableDbInstance {
	_init_.Initialize()

	if err := validateNewDataAwsOrderableDbInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsOrderableDbInstance{}

	_jsii_.Create(
		"@cdktn/aws-rds.DataAwsOrderableDbInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance aws_rds_orderable_db_instance} Data Source.
// Experimental.
func NewDataAwsOrderableDbInstance_Override(d DataAwsOrderableDbInstance, scope constructs.Construct, id *string, config *DataAwsOrderableDbInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-rds.DataAwsOrderableDbInstance",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetAvailabilityZoneGroup(val *string) {
	if err := j.validateSetAvailabilityZoneGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneGroup",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetEngineLatestVersion(val interface{}) {
	if err := j.validateSetEngineLatestVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineLatestVersion",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetInstanceClass(val *string) {
	if err := j.validateSetInstanceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceClass",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetLicenseModel(val *string) {
	if err := j.validateSetLicenseModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetPreferredEngineVersions(val *[]*string) {
	if err := j.validateSetPreferredEngineVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredEngineVersions",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetPreferredInstanceClasses(val *[]*string) {
	if err := j.validateSetPreferredInstanceClassesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredInstanceClasses",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetReadReplicaCapable(val interface{}) {
	if err := j.validateSetReadReplicaCapableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readReplicaCapable",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetSupportedEngineModes(val *[]*string) {
	if err := j.validateSetSupportedEngineModesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedEngineModes",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetSupportedNetworkTypes(val *[]*string) {
	if err := j.validateSetSupportedNetworkTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedNetworkTypes",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetSupportsClusters(val interface{}) {
	if err := j.validateSetSupportsClustersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsClusters",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetSupportsEnhancedMonitoring(val interface{}) {
	if err := j.validateSetSupportsEnhancedMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsEnhancedMonitoring",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetSupportsGlobalDatabases(val interface{}) {
	if err := j.validateSetSupportsGlobalDatabasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsGlobalDatabases",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetSupportsIamDatabaseAuthentication(val interface{}) {
	if err := j.validateSetSupportsIamDatabaseAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsIamDatabaseAuthentication",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetSupportsIops(val interface{}) {
	if err := j.validateSetSupportsIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsIops",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetSupportsKerberosAuthentication(val interface{}) {
	if err := j.validateSetSupportsKerberosAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsKerberosAuthentication",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetSupportsMultiAz(val interface{}) {
	if err := j.validateSetSupportsMultiAzParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsMultiAz",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetSupportsPerformanceInsights(val interface{}) {
	if err := j.validateSetSupportsPerformanceInsightsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsPerformanceInsights",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetSupportsStorageAutoscaling(val interface{}) {
	if err := j.validateSetSupportsStorageAutoscalingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsStorageAutoscaling",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetSupportsStorageEncryption(val interface{}) {
	if err := j.validateSetSupportsStorageEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsStorageEncryption",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrderableDbInstance)SetVpc(val interface{}) {
	if err := j.validateSetVpcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpc",
		val,
	)
}

// Generates CDKTN code for importing a DataAwsOrderableDbInstance resource upon running "cdktn plan <stack-name>".
// Experimental.
func DataAwsOrderableDbInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsOrderableDbInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-rds.DataAwsOrderableDbInstance",
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
func DataAwsOrderableDbInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsOrderableDbInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-rds.DataAwsOrderableDbInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsOrderableDbInstance_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsOrderableDbInstance_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-rds.DataAwsOrderableDbInstance",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsOrderableDbInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsOrderableDbInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-rds.DataAwsOrderableDbInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsOrderableDbInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-rds.DataAwsOrderableDbInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsOrderableDbInstance) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsOrderableDbInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsOrderableDbInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsOrderableDbInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsOrderableDbInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsOrderableDbInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsOrderableDbInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsOrderableDbInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsOrderableDbInstance) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsOrderableDbInstance) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetAvailabilityZoneGroup() {
	_jsii_.InvokeVoid(
		d,
		"resetAvailabilityZoneGroup",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetEngineLatestVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetEngineLatestVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetInstanceClass() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceClass",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetLicenseModel() {
	_jsii_.InvokeVoid(
		d,
		"resetLicenseModel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetPreferredEngineVersions() {
	_jsii_.InvokeVoid(
		d,
		"resetPreferredEngineVersions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetPreferredInstanceClasses() {
	_jsii_.InvokeVoid(
		d,
		"resetPreferredInstanceClasses",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetReadReplicaCapable() {
	_jsii_.InvokeVoid(
		d,
		"resetReadReplicaCapable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetStorageType() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetSupportedEngineModes() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportedEngineModes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetSupportedNetworkTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportedNetworkTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetSupportsClusters() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportsClusters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetSupportsEnhancedMonitoring() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportsEnhancedMonitoring",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetSupportsGlobalDatabases() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportsGlobalDatabases",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetSupportsIamDatabaseAuthentication() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportsIamDatabaseAuthentication",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetSupportsIops() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportsIops",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetSupportsKerberosAuthentication() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportsKerberosAuthentication",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetSupportsMultiAz() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportsMultiAz",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetSupportsPerformanceInsights() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportsPerformanceInsights",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetSupportsStorageAutoscaling() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportsStorageAutoscaling",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetSupportsStorageEncryption() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportsStorageEncryption",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ResetVpc() {
	_jsii_.InvokeVoid(
		d,
		"resetVpc",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOrderableDbInstance) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

