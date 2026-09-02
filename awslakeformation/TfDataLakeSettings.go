package awslakeformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslakeformation/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awslakeformation/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings aws_lakeformation_data_lake_settings}.
// Experimental.
type TfDataLakeSettings interface {
	cdktn.TerraformResource
	// Experimental.
	Admins() *[]*string
	// Experimental.
	SetAdmins(val *[]*string)
	// Experimental.
	AdminsInput() *[]*string
	// Experimental.
	AllowExternalDataFiltering() interface{}
	// Experimental.
	SetAllowExternalDataFiltering(val interface{})
	// Experimental.
	AllowExternalDataFilteringInput() interface{}
	// Experimental.
	AllowFullTableExternalDataAccess() interface{}
	// Experimental.
	SetAllowFullTableExternalDataAccess(val interface{})
	// Experimental.
	AllowFullTableExternalDataAccessInput() interface{}
	// Experimental.
	AuthorizedSessionTagValueList() *[]*string
	// Experimental.
	SetAuthorizedSessionTagValueList(val *[]*string)
	// Experimental.
	AuthorizedSessionTagValueListInput() *[]*string
	// Experimental.
	CatalogId() *string
	// Experimental.
	SetCatalogId(val *string)
	// Experimental.
	CatalogIdInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	CreateDatabaseDefaultPermissions() TfDataLakeSettings_CreateDatabaseDefaultPermissionsPropertyList
	// Experimental.
	CreateDatabaseDefaultPermissionsInput() interface{}
	// Experimental.
	CreateTableDefaultPermissions() TfDataLakeSettings_CreateTableDefaultPermissionsPropertyList
	// Experimental.
	CreateTableDefaultPermissionsInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ExternalDataFilteringAllowList() *[]*string
	// Experimental.
	SetExternalDataFilteringAllowList(val *[]*string)
	// Experimental.
	ExternalDataFilteringAllowListInput() *[]*string
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
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Parameters() *map[string]*string
	// Experimental.
	SetParameters(val *map[string]*string)
	// Experimental.
	ParametersInput() *map[string]*string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	ReadOnlyAdmins() *[]*string
	// Experimental.
	SetReadOnlyAdmins(val *[]*string)
	// Experimental.
	ReadOnlyAdminsInput() *[]*string
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	TrustedResourceOwners() *[]*string
	// Experimental.
	SetTrustedResourceOwners(val *[]*string)
	// Experimental.
	TrustedResourceOwnersInput() *[]*string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Experimental.
	PutCreateDatabaseDefaultPermissions(value interface{})
	// Experimental.
	PutCreateTableDefaultPermissions(value interface{})
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
	ResetAdmins()
	// Experimental.
	ResetAllowExternalDataFiltering()
	// Experimental.
	ResetAllowFullTableExternalDataAccess()
	// Experimental.
	ResetAuthorizedSessionTagValueList()
	// Experimental.
	ResetCatalogId()
	// Experimental.
	ResetCreateDatabaseDefaultPermissions()
	// Experimental.
	ResetCreateTableDefaultPermissions()
	// Experimental.
	ResetExternalDataFilteringAllowList()
	// Experimental.
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetParameters()
	// Experimental.
	ResetReadOnlyAdmins()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetTrustedResourceOwners()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	SynthesizeHclAttributes() *map[string]interface{}
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

// The jsii proxy struct for TfDataLakeSettings
type jsiiProxy_TfDataLakeSettings struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfDataLakeSettings) Admins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"admins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) AdminsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) AllowExternalDataFiltering() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowExternalDataFiltering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) AllowExternalDataFilteringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowExternalDataFilteringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) AllowFullTableExternalDataAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowFullTableExternalDataAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) AllowFullTableExternalDataAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowFullTableExternalDataAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) AuthorizedSessionTagValueList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizedSessionTagValueList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) AuthorizedSessionTagValueListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizedSessionTagValueListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) CreateDatabaseDefaultPermissions() TfDataLakeSettings_CreateDatabaseDefaultPermissionsPropertyList {
	var returns TfDataLakeSettings_CreateDatabaseDefaultPermissionsPropertyList
	_jsii_.Get(
		j,
		"createDatabaseDefaultPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) CreateDatabaseDefaultPermissionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createDatabaseDefaultPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) CreateTableDefaultPermissions() TfDataLakeSettings_CreateTableDefaultPermissionsPropertyList {
	var returns TfDataLakeSettings_CreateTableDefaultPermissionsPropertyList
	_jsii_.Get(
		j,
		"createTableDefaultPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) CreateTableDefaultPermissionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createTableDefaultPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) ExternalDataFilteringAllowList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalDataFilteringAllowList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) ExternalDataFilteringAllowListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalDataFilteringAllowListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) ReadOnlyAdmins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readOnlyAdmins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) ReadOnlyAdminsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readOnlyAdminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) TrustedResourceOwners() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedResourceOwners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataLakeSettings) TrustedResourceOwnersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedResourceOwnersInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings aws_lakeformation_data_lake_settings} Resource.
// Experimental.
func NewTfDataLakeSettings(scope constructs.Construct, id *string, config *TfDataLakeSettingsConfig) TfDataLakeSettings {
	_init_.Initialize()

	if err := validateNewTfDataLakeSettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDataLakeSettings{}

	_jsii_.Create(
		"@cdktn/aws-lake-formation.TfDataLakeSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings aws_lakeformation_data_lake_settings} Resource.
// Experimental.
func NewTfDataLakeSettings_Override(t TfDataLakeSettings, scope constructs.Construct, id *string, config *TfDataLakeSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lake-formation.TfDataLakeSettings",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfDataLakeSettings)SetAdmins(val *[]*string) {
	if err := j.validateSetAdminsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"admins",
		val,
	)
}

func (j *jsiiProxy_TfDataLakeSettings)SetAllowExternalDataFiltering(val interface{}) {
	if err := j.validateSetAllowExternalDataFilteringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowExternalDataFiltering",
		val,
	)
}

func (j *jsiiProxy_TfDataLakeSettings)SetAllowFullTableExternalDataAccess(val interface{}) {
	if err := j.validateSetAllowFullTableExternalDataAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowFullTableExternalDataAccess",
		val,
	)
}

func (j *jsiiProxy_TfDataLakeSettings)SetAuthorizedSessionTagValueList(val *[]*string) {
	if err := j.validateSetAuthorizedSessionTagValueListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizedSessionTagValueList",
		val,
	)
}

func (j *jsiiProxy_TfDataLakeSettings)SetCatalogId(val *string) {
	if err := j.validateSetCatalogIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_TfDataLakeSettings)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfDataLakeSettings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfDataLakeSettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfDataLakeSettings)SetExternalDataFilteringAllowList(val *[]*string) {
	if err := j.validateSetExternalDataFilteringAllowListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalDataFilteringAllowList",
		val,
	)
}

func (j *jsiiProxy_TfDataLakeSettings)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfDataLakeSettings)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfDataLakeSettings)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfDataLakeSettings)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_TfDataLakeSettings)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfDataLakeSettings)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfDataLakeSettings)SetReadOnlyAdmins(val *[]*string) {
	if err := j.validateSetReadOnlyAdminsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnlyAdmins",
		val,
	)
}

func (j *jsiiProxy_TfDataLakeSettings)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfDataLakeSettings)SetTrustedResourceOwners(val *[]*string) {
	if err := j.validateSetTrustedResourceOwnersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedResourceOwners",
		val,
	)
}

// Generates CDKTN code for importing a TfDataLakeSettings resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfDataLakeSettings_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfDataLakeSettings_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-lake-formation.TfDataLakeSettings",
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
func TfDataLakeSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDataLakeSettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lake-formation.TfDataLakeSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfDataLakeSettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDataLakeSettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lake-formation.TfDataLakeSettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfDataLakeSettings_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDataLakeSettings_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lake-formation.TfDataLakeSettings",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfDataLakeSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-lake-formation.TfDataLakeSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfDataLakeSettings) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfDataLakeSettings) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfDataLakeSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataLakeSettings) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataLakeSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataLakeSettings) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataLakeSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataLakeSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataLakeSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataLakeSettings) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataLakeSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataLakeSettings) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataLakeSettings) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfDataLakeSettings) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataLakeSettings) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := t.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataLakeSettings) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfDataLakeSettings) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfDataLakeSettings) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfDataLakeSettings) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfDataLakeSettings) PutCreateDatabaseDefaultPermissions(value interface{}) {
	if err := t.validatePutCreateDatabaseDefaultPermissionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCreateDatabaseDefaultPermissions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataLakeSettings) PutCreateTableDefaultPermissions(value interface{}) {
	if err := t.validatePutCreateTableDefaultPermissionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCreateTableDefaultPermissions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataLakeSettings) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfDataLakeSettings) ResetAdmins() {
	_jsii_.InvokeVoid(
		t,
		"resetAdmins",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataLakeSettings) ResetAllowExternalDataFiltering() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowExternalDataFiltering",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataLakeSettings) ResetAllowFullTableExternalDataAccess() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowFullTableExternalDataAccess",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataLakeSettings) ResetAuthorizedSessionTagValueList() {
	_jsii_.InvokeVoid(
		t,
		"resetAuthorizedSessionTagValueList",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataLakeSettings) ResetCatalogId() {
	_jsii_.InvokeVoid(
		t,
		"resetCatalogId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataLakeSettings) ResetCreateDatabaseDefaultPermissions() {
	_jsii_.InvokeVoid(
		t,
		"resetCreateDatabaseDefaultPermissions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataLakeSettings) ResetCreateTableDefaultPermissions() {
	_jsii_.InvokeVoid(
		t,
		"resetCreateTableDefaultPermissions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataLakeSettings) ResetExternalDataFilteringAllowList() {
	_jsii_.InvokeVoid(
		t,
		"resetExternalDataFilteringAllowList",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataLakeSettings) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataLakeSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataLakeSettings) ResetParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataLakeSettings) ResetReadOnlyAdmins() {
	_jsii_.InvokeVoid(
		t,
		"resetReadOnlyAdmins",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataLakeSettings) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataLakeSettings) ResetTrustedResourceOwners() {
	_jsii_.InvokeVoid(
		t,
		"resetTrustedResourceOwners",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataLakeSettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataLakeSettings) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataLakeSettings) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataLakeSettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataLakeSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataLakeSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataLakeSettings) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		t,
		"with",
		args,
		&returns,
	)

	return returns
}

