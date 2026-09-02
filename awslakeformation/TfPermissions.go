package awslakeformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslakeformation/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awslakeformation/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions aws_lakeformation_permissions}.
// Experimental.
type TfPermissions interface {
	cdktn.TerraformResource
	// Experimental.
	CatalogId() *string
	// Experimental.
	SetCatalogId(val *string)
	// Experimental.
	CatalogIdInput() *string
	// Experimental.
	CatalogResource() interface{}
	// Experimental.
	SetCatalogResource(val interface{})
	// Experimental.
	CatalogResourceInput() interface{}
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
	Database() TfPermissions_DatabasePropertyOutputReference
	// Experimental.
	DatabaseInput() *TfPermissions_DatabaseProperty
	// Experimental.
	DataCellsFilter() TfPermissions_DataCellsFilterPropertyOutputReference
	// Experimental.
	DataCellsFilterInput() *TfPermissions_DataCellsFilterProperty
	// Experimental.
	DataLocation() TfPermissions_DataLocationPropertyOutputReference
	// Experimental.
	DataLocationInput() *TfPermissions_DataLocationProperty
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	LfTag() TfPermissions_LfTagPropertyOutputReference
	// Experimental.
	LfTagInput() *TfPermissions_LfTagProperty
	// Experimental.
	LfTagPolicy() TfPermissions_LfTagPolicyPropertyOutputReference
	// Experimental.
	LfTagPolicyInput() *TfPermissions_LfTagPolicyProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Permissions() *[]*string
	// Experimental.
	SetPermissions(val *[]*string)
	// Experimental.
	PermissionsInput() *[]*string
	// Experimental.
	PermissionsWithGrantOption() *[]*string
	// Experimental.
	SetPermissionsWithGrantOption(val *[]*string)
	// Experimental.
	PermissionsWithGrantOptionInput() *[]*string
	// Experimental.
	Principal() *string
	// Experimental.
	SetPrincipal(val *string)
	// Experimental.
	PrincipalInput() *string
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
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	Table() TfPermissions_TablePropertyOutputReference
	// Experimental.
	TableInput() *TfPermissions_TableProperty
	// Experimental.
	TableWithColumns() TfPermissions_TableWithColumnsPropertyOutputReference
	// Experimental.
	TableWithColumnsInput() *TfPermissions_TableWithColumnsProperty
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutDatabase(value *TfPermissions_DatabaseProperty)
	// Experimental.
	PutDataCellsFilter(value *TfPermissions_DataCellsFilterProperty)
	// Experimental.
	PutDataLocation(value *TfPermissions_DataLocationProperty)
	// Experimental.
	PutLfTag(value *TfPermissions_LfTagProperty)
	// Experimental.
	PutLfTagPolicy(value *TfPermissions_LfTagPolicyProperty)
	// Experimental.
	PutTable(value *TfPermissions_TableProperty)
	// Experimental.
	PutTableWithColumns(value *TfPermissions_TableWithColumnsProperty)
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
	ResetCatalogId()
	// Experimental.
	ResetCatalogResource()
	// Experimental.
	ResetDatabase()
	// Experimental.
	ResetDataCellsFilter()
	// Experimental.
	ResetDataLocation()
	// Experimental.
	ResetId()
	// Experimental.
	ResetLfTag()
	// Experimental.
	ResetLfTagPolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPermissionsWithGrantOption()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetTable()
	// Experimental.
	ResetTableWithColumns()
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

// The jsii proxy struct for TfPermissions
type jsiiProxy_TfPermissions struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfPermissions) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) CatalogResource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) CatalogResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) Database() TfPermissions_DatabasePropertyOutputReference {
	var returns TfPermissions_DatabasePropertyOutputReference
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) DatabaseInput() *TfPermissions_DatabaseProperty {
	var returns *TfPermissions_DatabaseProperty
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) DataCellsFilter() TfPermissions_DataCellsFilterPropertyOutputReference {
	var returns TfPermissions_DataCellsFilterPropertyOutputReference
	_jsii_.Get(
		j,
		"dataCellsFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) DataCellsFilterInput() *TfPermissions_DataCellsFilterProperty {
	var returns *TfPermissions_DataCellsFilterProperty
	_jsii_.Get(
		j,
		"dataCellsFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) DataLocation() TfPermissions_DataLocationPropertyOutputReference {
	var returns TfPermissions_DataLocationPropertyOutputReference
	_jsii_.Get(
		j,
		"dataLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) DataLocationInput() *TfPermissions_DataLocationProperty {
	var returns *TfPermissions_DataLocationProperty
	_jsii_.Get(
		j,
		"dataLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) LfTag() TfPermissions_LfTagPropertyOutputReference {
	var returns TfPermissions_LfTagPropertyOutputReference
	_jsii_.Get(
		j,
		"lfTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) LfTagInput() *TfPermissions_LfTagProperty {
	var returns *TfPermissions_LfTagProperty
	_jsii_.Get(
		j,
		"lfTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) LfTagPolicy() TfPermissions_LfTagPolicyPropertyOutputReference {
	var returns TfPermissions_LfTagPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"lfTagPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) LfTagPolicyInput() *TfPermissions_LfTagPolicyProperty {
	var returns *TfPermissions_LfTagPolicyProperty
	_jsii_.Get(
		j,
		"lfTagPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) Permissions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) PermissionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) PermissionsWithGrantOption() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissionsWithGrantOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) PermissionsWithGrantOptionInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissionsWithGrantOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) PrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) Table() TfPermissions_TablePropertyOutputReference {
	var returns TfPermissions_TablePropertyOutputReference
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) TableInput() *TfPermissions_TableProperty {
	var returns *TfPermissions_TableProperty
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) TableWithColumns() TfPermissions_TableWithColumnsPropertyOutputReference {
	var returns TfPermissions_TableWithColumnsPropertyOutputReference
	_jsii_.Get(
		j,
		"tableWithColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) TableWithColumnsInput() *TfPermissions_TableWithColumnsProperty {
	var returns *TfPermissions_TableWithColumnsProperty
	_jsii_.Get(
		j,
		"tableWithColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissions) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions aws_lakeformation_permissions} Resource.
// Experimental.
func NewTfPermissions(scope constructs.Construct, id *string, config *TfPermissionsConfig) TfPermissions {
	_init_.Initialize()

	if err := validateNewTfPermissionsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPermissions{}

	_jsii_.Create(
		"@cdktn/aws-lake-formation.TfPermissions",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions aws_lakeformation_permissions} Resource.
// Experimental.
func NewTfPermissions_Override(t TfPermissions, scope constructs.Construct, id *string, config *TfPermissionsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lake-formation.TfPermissions",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfPermissions)SetCatalogId(val *string) {
	if err := j.validateSetCatalogIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_TfPermissions)SetCatalogResource(val interface{}) {
	if err := j.validateSetCatalogResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogResource",
		val,
	)
}

func (j *jsiiProxy_TfPermissions)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfPermissions)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfPermissions)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfPermissions)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfPermissions)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfPermissions)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfPermissions)SetPermissions(val *[]*string) {
	if err := j.validateSetPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_TfPermissions)SetPermissionsWithGrantOption(val *[]*string) {
	if err := j.validateSetPermissionsWithGrantOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissionsWithGrantOption",
		val,
	)
}

func (j *jsiiProxy_TfPermissions)SetPrincipal(val *string) {
	if err := j.validateSetPrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_TfPermissions)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfPermissions)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfPermissions)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTN code for importing a TfPermissions resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfPermissions_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfPermissions_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-lake-formation.TfPermissions",
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
func TfPermissions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfPermissions_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lake-formation.TfPermissions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfPermissions_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfPermissions_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lake-formation.TfPermissions",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfPermissions_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfPermissions_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lake-formation.TfPermissions",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfPermissions_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-lake-formation.TfPermissions",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfPermissions) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfPermissions) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfPermissions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPermissions) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPermissions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPermissions) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPermissions) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPermissions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPermissions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPermissions) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPermissions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPermissions) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPermissions) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfPermissions) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPermissions) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfPermissions) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfPermissions) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfPermissions) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfPermissions) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfPermissions) PutDatabase(value *TfPermissions_DatabaseProperty) {
	if err := t.validatePutDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDatabase",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPermissions) PutDataCellsFilter(value *TfPermissions_DataCellsFilterProperty) {
	if err := t.validatePutDataCellsFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDataCellsFilter",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPermissions) PutDataLocation(value *TfPermissions_DataLocationProperty) {
	if err := t.validatePutDataLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDataLocation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPermissions) PutLfTag(value *TfPermissions_LfTagProperty) {
	if err := t.validatePutLfTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLfTag",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPermissions) PutLfTagPolicy(value *TfPermissions_LfTagPolicyProperty) {
	if err := t.validatePutLfTagPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLfTagPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPermissions) PutTable(value *TfPermissions_TableProperty) {
	if err := t.validatePutTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTable",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPermissions) PutTableWithColumns(value *TfPermissions_TableWithColumnsProperty) {
	if err := t.validatePutTableWithColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTableWithColumns",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPermissions) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfPermissions) ResetCatalogId() {
	_jsii_.InvokeVoid(
		t,
		"resetCatalogId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPermissions) ResetCatalogResource() {
	_jsii_.InvokeVoid(
		t,
		"resetCatalogResource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPermissions) ResetDatabase() {
	_jsii_.InvokeVoid(
		t,
		"resetDatabase",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPermissions) ResetDataCellsFilter() {
	_jsii_.InvokeVoid(
		t,
		"resetDataCellsFilter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPermissions) ResetDataLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetDataLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPermissions) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPermissions) ResetLfTag() {
	_jsii_.InvokeVoid(
		t,
		"resetLfTag",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPermissions) ResetLfTagPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetLfTagPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPermissions) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPermissions) ResetPermissionsWithGrantOption() {
	_jsii_.InvokeVoid(
		t,
		"resetPermissionsWithGrantOption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPermissions) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPermissions) ResetTable() {
	_jsii_.InvokeVoid(
		t,
		"resetTable",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPermissions) ResetTableWithColumns() {
	_jsii_.InvokeVoid(
		t,
		"resetTableWithColumns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPermissions) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPermissions) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPermissions) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPermissions) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPermissions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPermissions) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPermissions) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

