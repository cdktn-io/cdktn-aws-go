package awslakeformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslakeformation/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awslakeformation/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lakeformation_permissions aws_lakeformation_permissions}.
// Experimental.
type DataTfPermissions interface {
	cdktn.TerraformDataSource
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
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	Database() DataTfPermissions_DatabasePropertyOutputReference
	// Experimental.
	DatabaseInput() *DataTfPermissions_DatabaseProperty
	// Experimental.
	DataCellsFilter() DataTfPermissions_DataCellsFilterPropertyOutputReference
	// Experimental.
	DataCellsFilterInput() *DataTfPermissions_DataCellsFilterProperty
	// Experimental.
	DataLocation() DataTfPermissions_DataLocationPropertyOutputReference
	// Experimental.
	DataLocationInput() *DataTfPermissions_DataLocationProperty
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
	LfTag() DataTfPermissions_LfTagPropertyOutputReference
	// Experimental.
	LfTagInput() *DataTfPermissions_LfTagProperty
	// Experimental.
	LfTagPolicy() DataTfPermissions_LfTagPolicyPropertyOutputReference
	// Experimental.
	LfTagPolicyInput() *DataTfPermissions_LfTagPolicyProperty
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
	PermissionsWithGrantOption() *[]*string
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
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	Table() DataTfPermissions_TablePropertyOutputReference
	// Experimental.
	TableInput() *DataTfPermissions_TableProperty
	// Experimental.
	TableWithColumns() DataTfPermissions_TableWithColumnsPropertyOutputReference
	// Experimental.
	TableWithColumnsInput() *DataTfPermissions_TableWithColumnsProperty
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutDatabase(value *DataTfPermissions_DatabaseProperty)
	// Experimental.
	PutDataCellsFilter(value *DataTfPermissions_DataCellsFilterProperty)
	// Experimental.
	PutDataLocation(value *DataTfPermissions_DataLocationProperty)
	// Experimental.
	PutLfTag(value *DataTfPermissions_LfTagProperty)
	// Experimental.
	PutLfTagPolicy(value *DataTfPermissions_LfTagPolicyProperty)
	// Experimental.
	PutTable(value *DataTfPermissions_TableProperty)
	// Experimental.
	PutTableWithColumns(value *DataTfPermissions_TableWithColumnsProperty)
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
	ResetRegion()
	// Experimental.
	ResetTable()
	// Experimental.
	ResetTableWithColumns()
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

// The jsii proxy struct for DataTfPermissions
type jsiiProxy_DataTfPermissions struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataTfPermissions) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) CatalogResource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) CatalogResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) Database() DataTfPermissions_DatabasePropertyOutputReference {
	var returns DataTfPermissions_DatabasePropertyOutputReference
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) DatabaseInput() *DataTfPermissions_DatabaseProperty {
	var returns *DataTfPermissions_DatabaseProperty
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) DataCellsFilter() DataTfPermissions_DataCellsFilterPropertyOutputReference {
	var returns DataTfPermissions_DataCellsFilterPropertyOutputReference
	_jsii_.Get(
		j,
		"dataCellsFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) DataCellsFilterInput() *DataTfPermissions_DataCellsFilterProperty {
	var returns *DataTfPermissions_DataCellsFilterProperty
	_jsii_.Get(
		j,
		"dataCellsFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) DataLocation() DataTfPermissions_DataLocationPropertyOutputReference {
	var returns DataTfPermissions_DataLocationPropertyOutputReference
	_jsii_.Get(
		j,
		"dataLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) DataLocationInput() *DataTfPermissions_DataLocationProperty {
	var returns *DataTfPermissions_DataLocationProperty
	_jsii_.Get(
		j,
		"dataLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) LfTag() DataTfPermissions_LfTagPropertyOutputReference {
	var returns DataTfPermissions_LfTagPropertyOutputReference
	_jsii_.Get(
		j,
		"lfTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) LfTagInput() *DataTfPermissions_LfTagProperty {
	var returns *DataTfPermissions_LfTagProperty
	_jsii_.Get(
		j,
		"lfTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) LfTagPolicy() DataTfPermissions_LfTagPolicyPropertyOutputReference {
	var returns DataTfPermissions_LfTagPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"lfTagPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) LfTagPolicyInput() *DataTfPermissions_LfTagPolicyProperty {
	var returns *DataTfPermissions_LfTagPolicyProperty
	_jsii_.Get(
		j,
		"lfTagPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) Permissions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) PermissionsWithGrantOption() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissionsWithGrantOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) PrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) Table() DataTfPermissions_TablePropertyOutputReference {
	var returns DataTfPermissions_TablePropertyOutputReference
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) TableInput() *DataTfPermissions_TableProperty {
	var returns *DataTfPermissions_TableProperty
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) TableWithColumns() DataTfPermissions_TableWithColumnsPropertyOutputReference {
	var returns DataTfPermissions_TableWithColumnsPropertyOutputReference
	_jsii_.Get(
		j,
		"tableWithColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) TableWithColumnsInput() *DataTfPermissions_TableWithColumnsProperty {
	var returns *DataTfPermissions_TableWithColumnsProperty
	_jsii_.Get(
		j,
		"tableWithColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPermissions) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lakeformation_permissions aws_lakeformation_permissions} Data Source.
// Experimental.
func NewDataTfPermissions(scope constructs.Construct, id *string, config *DataTfPermissionsConfig) DataTfPermissions {
	_init_.Initialize()

	if err := validateNewDataTfPermissionsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfPermissions{}

	_jsii_.Create(
		"@cdktn/aws-lake-formation.DataTfPermissions",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lakeformation_permissions aws_lakeformation_permissions} Data Source.
// Experimental.
func NewDataTfPermissions_Override(d DataTfPermissions, scope constructs.Construct, id *string, config *DataTfPermissionsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lake-formation.DataTfPermissions",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataTfPermissions)SetCatalogId(val *string) {
	if err := j.validateSetCatalogIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_DataTfPermissions)SetCatalogResource(val interface{}) {
	if err := j.validateSetCatalogResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogResource",
		val,
	)
}

func (j *jsiiProxy_DataTfPermissions)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataTfPermissions)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataTfPermissions)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataTfPermissions)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataTfPermissions)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataTfPermissions)SetPrincipal(val *string) {
	if err := j.validateSetPrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_DataTfPermissions)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataTfPermissions)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTN code for importing a DataTfPermissions resource upon running "cdktn plan <stack-name>".
// Experimental.
func DataTfPermissions_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataTfPermissions_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-lake-formation.DataTfPermissions",
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
func DataTfPermissions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfPermissions_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lake-formation.DataTfPermissions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataTfPermissions_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfPermissions_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lake-formation.DataTfPermissions",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataTfPermissions_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfPermissions_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lake-formation.DataTfPermissions",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTfPermissions_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-lake-formation.DataTfPermissions",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTfPermissions) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTfPermissions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataTfPermissions) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfPermissions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataTfPermissions) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataTfPermissions) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataTfPermissions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataTfPermissions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataTfPermissions) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataTfPermissions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataTfPermissions) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfPermissions) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTfPermissions) PutDatabase(value *DataTfPermissions_DatabaseProperty) {
	if err := d.validatePutDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDatabase",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfPermissions) PutDataCellsFilter(value *DataTfPermissions_DataCellsFilterProperty) {
	if err := d.validatePutDataCellsFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDataCellsFilter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfPermissions) PutDataLocation(value *DataTfPermissions_DataLocationProperty) {
	if err := d.validatePutDataLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDataLocation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfPermissions) PutLfTag(value *DataTfPermissions_LfTagProperty) {
	if err := d.validatePutLfTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLfTag",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfPermissions) PutLfTagPolicy(value *DataTfPermissions_LfTagPolicyProperty) {
	if err := d.validatePutLfTagPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLfTagPolicy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfPermissions) PutTable(value *DataTfPermissions_TableProperty) {
	if err := d.validatePutTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfPermissions) PutTableWithColumns(value *DataTfPermissions_TableWithColumnsProperty) {
	if err := d.validatePutTableWithColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTableWithColumns",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfPermissions) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataTfPermissions) ResetCatalogId() {
	_jsii_.InvokeVoid(
		d,
		"resetCatalogId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPermissions) ResetCatalogResource() {
	_jsii_.InvokeVoid(
		d,
		"resetCatalogResource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPermissions) ResetDatabase() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPermissions) ResetDataCellsFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetDataCellsFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPermissions) ResetDataLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetDataLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPermissions) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPermissions) ResetLfTag() {
	_jsii_.InvokeVoid(
		d,
		"resetLfTag",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPermissions) ResetLfTagPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetLfTagPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPermissions) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPermissions) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPermissions) ResetTable() {
	_jsii_.InvokeVoid(
		d,
		"resetTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPermissions) ResetTableWithColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetTableWithColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPermissions) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfPermissions) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfPermissions) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfPermissions) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfPermissions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfPermissions) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfPermissions) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

