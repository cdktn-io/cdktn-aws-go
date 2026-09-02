package awsmanagedgrafana

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmanagedgrafana/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsmanagedgrafana/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace_saml_configuration aws_grafana_workspace_saml_configuration}.
// Experimental.
type TfWorkspaceSamlConfiguration interface {
	cdktn.TerraformResource
	// Experimental.
	AdminRoleValues() *[]*string
	// Experimental.
	SetAdminRoleValues(val *[]*string)
	// Experimental.
	AdminRoleValuesInput() *[]*string
	// Experimental.
	AllowedOrganizations() *[]*string
	// Experimental.
	SetAllowedOrganizations(val *[]*string)
	// Experimental.
	AllowedOrganizationsInput() *[]*string
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
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EditorRoleValues() *[]*string
	// Experimental.
	SetEditorRoleValues(val *[]*string)
	// Experimental.
	EditorRoleValuesInput() *[]*string
	// Experimental.
	EmailAssertion() *string
	// Experimental.
	SetEmailAssertion(val *string)
	// Experimental.
	EmailAssertionInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	GroupsAssertion() *string
	// Experimental.
	SetGroupsAssertion(val *string)
	// Experimental.
	GroupsAssertionInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	IdpMetadataUrl() *string
	// Experimental.
	SetIdpMetadataUrl(val *string)
	// Experimental.
	IdpMetadataUrlInput() *string
	// Experimental.
	IdpMetadataXml() *string
	// Experimental.
	SetIdpMetadataXml(val *string)
	// Experimental.
	IdpMetadataXmlInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LoginAssertion() *string
	// Experimental.
	SetLoginAssertion(val *string)
	// Experimental.
	LoginAssertionInput() *string
	// Experimental.
	LoginValidityDuration() *float64
	// Experimental.
	SetLoginValidityDuration(val *float64)
	// Experimental.
	LoginValidityDurationInput() *float64
	// Experimental.
	NameAssertion() *string
	// Experimental.
	SetNameAssertion(val *string)
	// Experimental.
	NameAssertionInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OrgAssertion() *string
	// Experimental.
	SetOrgAssertion(val *string)
	// Experimental.
	OrgAssertionInput() *string
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
	RoleAssertion() *string
	// Experimental.
	SetRoleAssertion(val *string)
	// Experimental.
	RoleAssertionInput() *string
	// Experimental.
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() TfWorkspaceSamlConfiguration_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	WorkspaceId() *string
	// Experimental.
	SetWorkspaceId(val *string)
	// Experimental.
	WorkspaceIdInput() *string
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
	PutTimeouts(value *TfWorkspaceSamlConfiguration_TimeoutsProperty)
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
	ResetAdminRoleValues()
	// Experimental.
	ResetAllowedOrganizations()
	// Experimental.
	ResetEmailAssertion()
	// Experimental.
	ResetGroupsAssertion()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIdpMetadataUrl()
	// Experimental.
	ResetIdpMetadataXml()
	// Experimental.
	ResetLoginAssertion()
	// Experimental.
	ResetLoginValidityDuration()
	// Experimental.
	ResetNameAssertion()
	// Experimental.
	ResetOrgAssertion()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRoleAssertion()
	// Experimental.
	ResetTimeouts()
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

// The jsii proxy struct for TfWorkspaceSamlConfiguration
type jsiiProxy_TfWorkspaceSamlConfiguration struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) AdminRoleValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminRoleValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) AdminRoleValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminRoleValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) AllowedOrganizations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrganizations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) AllowedOrganizationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrganizationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) EditorRoleValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"editorRoleValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) EditorRoleValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"editorRoleValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) EmailAssertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) EmailAssertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAssertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) GroupsAssertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) GroupsAssertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAssertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) IdpMetadataUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpMetadataUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) IdpMetadataUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpMetadataUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) IdpMetadataXml() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpMetadataXml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) IdpMetadataXmlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpMetadataXmlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) LoginAssertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) LoginAssertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginAssertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) LoginValidityDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loginValidityDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) LoginValidityDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loginValidityDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) NameAssertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) NameAssertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameAssertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) OrgAssertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) OrgAssertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgAssertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) RoleAssertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) RoleAssertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleAssertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) Timeouts() TfWorkspaceSamlConfiguration_TimeoutsPropertyOutputReference {
	var returns TfWorkspaceSamlConfiguration_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration) WorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace_saml_configuration aws_grafana_workspace_saml_configuration} Resource.
// Experimental.
func NewTfWorkspaceSamlConfiguration(scope constructs.Construct, id *string, config *TfWorkspaceSamlConfigurationConfig) TfWorkspaceSamlConfiguration {
	_init_.Initialize()

	if err := validateNewTfWorkspaceSamlConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWorkspaceSamlConfiguration{}

	_jsii_.Create(
		"@cdktn/aws-managed-grafana.TfWorkspaceSamlConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace_saml_configuration aws_grafana_workspace_saml_configuration} Resource.
// Experimental.
func NewTfWorkspaceSamlConfiguration_Override(t TfWorkspaceSamlConfiguration, scope constructs.Construct, id *string, config *TfWorkspaceSamlConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-managed-grafana.TfWorkspaceSamlConfiguration",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetAdminRoleValues(val *[]*string) {
	if err := j.validateSetAdminRoleValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminRoleValues",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetAllowedOrganizations(val *[]*string) {
	if err := j.validateSetAllowedOrganizationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOrganizations",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetEditorRoleValues(val *[]*string) {
	if err := j.validateSetEditorRoleValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"editorRoleValues",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetEmailAssertion(val *string) {
	if err := j.validateSetEmailAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAssertion",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetGroupsAssertion(val *string) {
	if err := j.validateSetGroupsAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsAssertion",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetIdpMetadataUrl(val *string) {
	if err := j.validateSetIdpMetadataUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpMetadataUrl",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetIdpMetadataXml(val *string) {
	if err := j.validateSetIdpMetadataXmlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpMetadataXml",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetLoginAssertion(val *string) {
	if err := j.validateSetLoginAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginAssertion",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetLoginValidityDuration(val *float64) {
	if err := j.validateSetLoginValidityDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginValidityDuration",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetNameAssertion(val *string) {
	if err := j.validateSetNameAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameAssertion",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetOrgAssertion(val *string) {
	if err := j.validateSetOrgAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orgAssertion",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetRoleAssertion(val *string) {
	if err := j.validateSetRoleAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleAssertion",
		val,
	)
}

func (j *jsiiProxy_TfWorkspaceSamlConfiguration)SetWorkspaceId(val *string) {
	if err := j.validateSetWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceId",
		val,
	)
}

// Generates CDKTN code for importing a TfWorkspaceSamlConfiguration resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfWorkspaceSamlConfiguration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfWorkspaceSamlConfiguration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-managed-grafana.TfWorkspaceSamlConfiguration",
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
func TfWorkspaceSamlConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfWorkspaceSamlConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-managed-grafana.TfWorkspaceSamlConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfWorkspaceSamlConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfWorkspaceSamlConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-managed-grafana.TfWorkspaceSamlConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfWorkspaceSamlConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfWorkspaceSamlConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-managed-grafana.TfWorkspaceSamlConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfWorkspaceSamlConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-managed-grafana.TfWorkspaceSamlConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) PutTimeouts(value *TfWorkspaceSamlConfiguration_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) ResetAdminRoleValues() {
	_jsii_.InvokeVoid(
		t,
		"resetAdminRoleValues",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) ResetAllowedOrganizations() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowedOrganizations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) ResetEmailAssertion() {
	_jsii_.InvokeVoid(
		t,
		"resetEmailAssertion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) ResetGroupsAssertion() {
	_jsii_.InvokeVoid(
		t,
		"resetGroupsAssertion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) ResetIdpMetadataUrl() {
	_jsii_.InvokeVoid(
		t,
		"resetIdpMetadataUrl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) ResetIdpMetadataXml() {
	_jsii_.InvokeVoid(
		t,
		"resetIdpMetadataXml",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) ResetLoginAssertion() {
	_jsii_.InvokeVoid(
		t,
		"resetLoginAssertion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) ResetLoginValidityDuration() {
	_jsii_.InvokeVoid(
		t,
		"resetLoginValidityDuration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) ResetNameAssertion() {
	_jsii_.InvokeVoid(
		t,
		"resetNameAssertion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) ResetOrgAssertion() {
	_jsii_.InvokeVoid(
		t,
		"resetOrgAssertion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) ResetRoleAssertion() {
	_jsii_.InvokeVoid(
		t,
		"resetRoleAssertion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkspaceSamlConfiguration) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

