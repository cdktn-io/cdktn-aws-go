package neptune

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/neptune/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/neptune/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/neptune_engine_version aws_neptune_engine_version}.
// Experimental.
type DataAwsEngineVersion interface {
	cdktn.TerraformDataSource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DefaultCharacterSet() *string
	// Experimental.
	DefaultOnly() interface{}
	// Experimental.
	SetDefaultOnly(val interface{})
	// Experimental.
	DefaultOnlyInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	Engine() *string
	// Experimental.
	SetEngine(val *string)
	// Experimental.
	EngineDescription() *string
	// Experimental.
	EngineInput() *string
	// Experimental.
	ExportableLogTypes() *[]*string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HasMajorTarget() interface{}
	// Experimental.
	SetHasMajorTarget(val interface{})
	// Experimental.
	HasMajorTargetInput() interface{}
	// Experimental.
	HasMinorTarget() interface{}
	// Experimental.
	SetHasMinorTarget(val interface{})
	// Experimental.
	HasMinorTargetInput() interface{}
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	Latest() interface{}
	// Experimental.
	SetLatest(val interface{})
	// Experimental.
	LatestInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	ParameterGroupFamily() *string
	// Experimental.
	SetParameterGroupFamily(val *string)
	// Experimental.
	ParameterGroupFamilyInput() *string
	// Experimental.
	PreferredMajorTargets() *[]*string
	// Experimental.
	SetPreferredMajorTargets(val *[]*string)
	// Experimental.
	PreferredMajorTargetsInput() *[]*string
	// Experimental.
	PreferredUpgradeTargets() *[]*string
	// Experimental.
	SetPreferredUpgradeTargets(val *[]*string)
	// Experimental.
	PreferredUpgradeTargetsInput() *[]*string
	// Experimental.
	PreferredVersions() *[]*string
	// Experimental.
	SetPreferredVersions(val *[]*string)
	// Experimental.
	PreferredVersionsInput() *[]*string
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
	SupportedCharacterSets() *[]*string
	// Experimental.
	SupportedTimezones() *[]*string
	// Experimental.
	SupportsGlobalDatabases() cdktn.IResolvable
	// Experimental.
	SupportsLogExportsToCloudwatch() cdktn.IResolvable
	// Experimental.
	SupportsReadReplica() cdktn.IResolvable
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	ValidMajorTargets() *[]*string
	// Experimental.
	ValidMinorTargets() *[]*string
	// Experimental.
	ValidUpgradeTargets() *[]*string
	// Experimental.
	Version() *string
	// Experimental.
	SetVersion(val *string)
	// Experimental.
	VersionActual() *string
	// Experimental.
	VersionDescription() *string
	// Experimental.
	VersionInput() *string
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
	ResetDefaultOnly()
	// Experimental.
	ResetEngine()
	// Experimental.
	ResetHasMajorTarget()
	// Experimental.
	ResetHasMinorTarget()
	// Experimental.
	ResetId()
	// Experimental.
	ResetLatest()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetParameterGroupFamily()
	// Experimental.
	ResetPreferredMajorTargets()
	// Experimental.
	ResetPreferredUpgradeTargets()
	// Experimental.
	ResetPreferredVersions()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetVersion()
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

// The jsii proxy struct for DataAwsEngineVersion
type jsiiProxy_DataAwsEngineVersion struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataAwsEngineVersion) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) DefaultCharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultCharacterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) DefaultOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) DefaultOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) EngineDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) ExportableLogTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exportableLogTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) HasMajorTarget() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasMajorTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) HasMajorTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasMajorTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) HasMinorTarget() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasMinorTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) HasMinorTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasMinorTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) Latest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"latest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) LatestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"latestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) ParameterGroupFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) ParameterGroupFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) PreferredMajorTargets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredMajorTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) PreferredMajorTargetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredMajorTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) PreferredUpgradeTargets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredUpgradeTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) PreferredUpgradeTargetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredUpgradeTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) PreferredVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) PreferredVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) SupportedCharacterSets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedCharacterSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) SupportedTimezones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedTimezones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) SupportsGlobalDatabases() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"supportsGlobalDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) SupportsLogExportsToCloudwatch() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"supportsLogExportsToCloudwatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) SupportsReadReplica() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"supportsReadReplica",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) ValidMajorTargets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validMajorTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) ValidMinorTargets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validMinorTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) ValidUpgradeTargets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validUpgradeTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) VersionActual() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionActual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) VersionDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEngineVersion) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/neptune_engine_version aws_neptune_engine_version} Data Source.
// Experimental.
func NewDataAwsEngineVersion(scope constructs.Construct, id *string, config *DataAwsEngineVersionConfig) DataAwsEngineVersion {
	_init_.Initialize()

	if err := validateNewDataAwsEngineVersionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsEngineVersion{}

	_jsii_.Create(
		"@cdktn/aws-neptune.DataAwsEngineVersion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/neptune_engine_version aws_neptune_engine_version} Data Source.
// Experimental.
func NewDataAwsEngineVersion_Override(d DataAwsEngineVersion, scope constructs.Construct, id *string, config *DataAwsEngineVersionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-neptune.DataAwsEngineVersion",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsEngineVersion)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsEngineVersion)SetDefaultOnly(val interface{}) {
	if err := j.validateSetDefaultOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultOnly",
		val,
	)
}

func (j *jsiiProxy_DataAwsEngineVersion)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsEngineVersion)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_DataAwsEngineVersion)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsEngineVersion)SetHasMajorTarget(val interface{}) {
	if err := j.validateSetHasMajorTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasMajorTarget",
		val,
	)
}

func (j *jsiiProxy_DataAwsEngineVersion)SetHasMinorTarget(val interface{}) {
	if err := j.validateSetHasMinorTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasMinorTarget",
		val,
	)
}

func (j *jsiiProxy_DataAwsEngineVersion)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsEngineVersion)SetLatest(val interface{}) {
	if err := j.validateSetLatestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latest",
		val,
	)
}

func (j *jsiiProxy_DataAwsEngineVersion)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsEngineVersion)SetParameterGroupFamily(val *string) {
	if err := j.validateSetParameterGroupFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterGroupFamily",
		val,
	)
}

func (j *jsiiProxy_DataAwsEngineVersion)SetPreferredMajorTargets(val *[]*string) {
	if err := j.validateSetPreferredMajorTargetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredMajorTargets",
		val,
	)
}

func (j *jsiiProxy_DataAwsEngineVersion)SetPreferredUpgradeTargets(val *[]*string) {
	if err := j.validateSetPreferredUpgradeTargetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredUpgradeTargets",
		val,
	)
}

func (j *jsiiProxy_DataAwsEngineVersion)SetPreferredVersions(val *[]*string) {
	if err := j.validateSetPreferredVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredVersions",
		val,
	)
}

func (j *jsiiProxy_DataAwsEngineVersion)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsEngineVersion)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataAwsEngineVersion)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTN code for importing a DataAwsEngineVersion resource upon running "cdktn plan <stack-name>".
// Experimental.
func DataAwsEngineVersion_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsEngineVersion_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-neptune.DataAwsEngineVersion",
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
func DataAwsEngineVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsEngineVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-neptune.DataAwsEngineVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsEngineVersion_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsEngineVersion_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-neptune.DataAwsEngineVersion",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsEngineVersion_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsEngineVersion_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-neptune.DataAwsEngineVersion",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsEngineVersion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-neptune.DataAwsEngineVersion",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsEngineVersion) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsEngineVersion) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsEngineVersion) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsEngineVersion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsEngineVersion) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsEngineVersion) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsEngineVersion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsEngineVersion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsEngineVersion) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsEngineVersion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsEngineVersion) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsEngineVersion) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsEngineVersion) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataAwsEngineVersion) ResetDefaultOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEngineVersion) ResetEngine() {
	_jsii_.InvokeVoid(
		d,
		"resetEngine",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEngineVersion) ResetHasMajorTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetHasMajorTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEngineVersion) ResetHasMinorTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetHasMinorTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEngineVersion) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEngineVersion) ResetLatest() {
	_jsii_.InvokeVoid(
		d,
		"resetLatest",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEngineVersion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEngineVersion) ResetParameterGroupFamily() {
	_jsii_.InvokeVoid(
		d,
		"resetParameterGroupFamily",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEngineVersion) ResetPreferredMajorTargets() {
	_jsii_.InvokeVoid(
		d,
		"resetPreferredMajorTargets",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEngineVersion) ResetPreferredUpgradeTargets() {
	_jsii_.InvokeVoid(
		d,
		"resetPreferredUpgradeTargets",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEngineVersion) ResetPreferredVersions() {
	_jsii_.InvokeVoid(
		d,
		"resetPreferredVersions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEngineVersion) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEngineVersion) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEngineVersion) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEngineVersion) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEngineVersion) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEngineVersion) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEngineVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEngineVersion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEngineVersion) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

