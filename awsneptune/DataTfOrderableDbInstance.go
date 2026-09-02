package awsneptune

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsneptune/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsneptune/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/neptune_orderable_db_instance aws_neptune_orderable_db_instance}.
// Experimental.
type DataTfOrderableDbInstance interface {
	cdktn.TerraformDataSource
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
	ReadReplicaCapable() cdktn.IResolvable
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	StorageType() *string
	// Experimental.
	SupportsEnhancedMonitoring() cdktn.IResolvable
	// Experimental.
	SupportsIamDatabaseAuthentication() cdktn.IResolvable
	// Experimental.
	SupportsIops() cdktn.IResolvable
	// Experimental.
	SupportsPerformanceInsights() cdktn.IResolvable
	// Experimental.
	SupportsStorageEncryption() cdktn.IResolvable
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
	ResetEngine()
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
	ResetPreferredInstanceClasses()
	// Experimental.
	ResetRegion()
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

// The jsii proxy struct for DataTfOrderableDbInstance
type jsiiProxy_DataTfOrderableDbInstance struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataTfOrderableDbInstance) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) InstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) InstanceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) LicenseModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) MaxIopsPerDbInstance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIopsPerDbInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) MaxIopsPerGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIopsPerGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) MaxStorageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStorageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) MinIopsPerDbInstance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIopsPerDbInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) MinIopsPerGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIopsPerGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) MinStorageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minStorageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) MultiAzCapable() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"multiAzCapable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) PreferredInstanceClasses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredInstanceClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) PreferredInstanceClassesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredInstanceClassesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) ReadReplicaCapable() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"readReplicaCapable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) SupportsEnhancedMonitoring() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"supportsEnhancedMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) SupportsIamDatabaseAuthentication() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"supportsIamDatabaseAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) SupportsIops() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"supportsIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) SupportsPerformanceInsights() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"supportsPerformanceInsights",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) SupportsStorageEncryption() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"supportsStorageEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) Vpc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfOrderableDbInstance) VpcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/neptune_orderable_db_instance aws_neptune_orderable_db_instance} Data Source.
// Experimental.
func NewDataTfOrderableDbInstance(scope constructs.Construct, id *string, config *DataTfOrderableDbInstanceConfig) DataTfOrderableDbInstance {
	_init_.Initialize()

	if err := validateNewDataTfOrderableDbInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfOrderableDbInstance{}

	_jsii_.Create(
		"@cdktn/aws-neptune.DataTfOrderableDbInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/neptune_orderable_db_instance aws_neptune_orderable_db_instance} Data Source.
// Experimental.
func NewDataTfOrderableDbInstance_Override(d DataTfOrderableDbInstance, scope constructs.Construct, id *string, config *DataTfOrderableDbInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-neptune.DataTfOrderableDbInstance",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataTfOrderableDbInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataTfOrderableDbInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataTfOrderableDbInstance)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_DataTfOrderableDbInstance)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_DataTfOrderableDbInstance)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataTfOrderableDbInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataTfOrderableDbInstance)SetInstanceClass(val *string) {
	if err := j.validateSetInstanceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceClass",
		val,
	)
}

func (j *jsiiProxy_DataTfOrderableDbInstance)SetLicenseModel(val *string) {
	if err := j.validateSetLicenseModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_DataTfOrderableDbInstance)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataTfOrderableDbInstance)SetPreferredInstanceClasses(val *[]*string) {
	if err := j.validateSetPreferredInstanceClassesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredInstanceClasses",
		val,
	)
}

func (j *jsiiProxy_DataTfOrderableDbInstance)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataTfOrderableDbInstance)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataTfOrderableDbInstance)SetVpc(val interface{}) {
	if err := j.validateSetVpcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpc",
		val,
	)
}

// Generates CDKTN code for importing a DataTfOrderableDbInstance resource upon running "cdktn plan <stack-name>".
// Experimental.
func DataTfOrderableDbInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataTfOrderableDbInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-neptune.DataTfOrderableDbInstance",
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
func DataTfOrderableDbInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfOrderableDbInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-neptune.DataTfOrderableDbInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataTfOrderableDbInstance_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfOrderableDbInstance_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-neptune.DataTfOrderableDbInstance",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataTfOrderableDbInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfOrderableDbInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-neptune.DataTfOrderableDbInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTfOrderableDbInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-neptune.DataTfOrderableDbInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTfOrderableDbInstance) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTfOrderableDbInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataTfOrderableDbInstance) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfOrderableDbInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataTfOrderableDbInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataTfOrderableDbInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataTfOrderableDbInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataTfOrderableDbInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataTfOrderableDbInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataTfOrderableDbInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataTfOrderableDbInstance) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfOrderableDbInstance) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTfOrderableDbInstance) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataTfOrderableDbInstance) ResetEngine() {
	_jsii_.InvokeVoid(
		d,
		"resetEngine",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfOrderableDbInstance) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfOrderableDbInstance) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfOrderableDbInstance) ResetInstanceClass() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceClass",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfOrderableDbInstance) ResetLicenseModel() {
	_jsii_.InvokeVoid(
		d,
		"resetLicenseModel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfOrderableDbInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfOrderableDbInstance) ResetPreferredInstanceClasses() {
	_jsii_.InvokeVoid(
		d,
		"resetPreferredInstanceClasses",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfOrderableDbInstance) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfOrderableDbInstance) ResetVpc() {
	_jsii_.InvokeVoid(
		d,
		"resetVpc",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfOrderableDbInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfOrderableDbInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfOrderableDbInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfOrderableDbInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfOrderableDbInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfOrderableDbInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfOrderableDbInstance) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

