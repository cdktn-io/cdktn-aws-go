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
type AwsGrafanaWorkspaceSamlConfiguration interface {
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
	Timeouts() AwsGrafanaWorkspaceSamlConfiguration_TimeoutsPropertyOutputReference
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
	PutTimeouts(value *AwsGrafanaWorkspaceSamlConfiguration_TimeoutsProperty)
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

// The jsii proxy struct for AwsGrafanaWorkspaceSamlConfiguration
type jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) AdminRoleValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminRoleValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) AdminRoleValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminRoleValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) AllowedOrganizations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrganizations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) AllowedOrganizationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrganizationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) EditorRoleValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"editorRoleValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) EditorRoleValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"editorRoleValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) EmailAssertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) EmailAssertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAssertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) GroupsAssertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) GroupsAssertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAssertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) IdpMetadataUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpMetadataUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) IdpMetadataUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpMetadataUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) IdpMetadataXml() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpMetadataXml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) IdpMetadataXmlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpMetadataXmlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) LoginAssertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) LoginAssertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginAssertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) LoginValidityDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loginValidityDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) LoginValidityDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loginValidityDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) NameAssertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) NameAssertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameAssertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) OrgAssertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) OrgAssertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgAssertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) RoleAssertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) RoleAssertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleAssertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) Timeouts() AwsGrafanaWorkspaceSamlConfiguration_TimeoutsPropertyOutputReference {
	var returns AwsGrafanaWorkspaceSamlConfiguration_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) WorkspaceIdInput() *string {
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
func NewAwsGrafanaWorkspaceSamlConfiguration(scope constructs.Construct, id *string, config *AwsGrafanaWorkspaceSamlConfigurationConfig) AwsGrafanaWorkspaceSamlConfiguration {
	_init_.Initialize()

	if err := validateNewAwsGrafanaWorkspaceSamlConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration{}

	_jsii_.Create(
		"@cdktn/aws-managed-grafana.AwsGrafanaWorkspaceSamlConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace_saml_configuration aws_grafana_workspace_saml_configuration} Resource.
// Experimental.
func NewAwsGrafanaWorkspaceSamlConfiguration_Override(a AwsGrafanaWorkspaceSamlConfiguration, scope constructs.Construct, id *string, config *AwsGrafanaWorkspaceSamlConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-managed-grafana.AwsGrafanaWorkspaceSamlConfiguration",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetAdminRoleValues(val *[]*string) {
	if err := j.validateSetAdminRoleValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminRoleValues",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetAllowedOrganizations(val *[]*string) {
	if err := j.validateSetAllowedOrganizationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOrganizations",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetEditorRoleValues(val *[]*string) {
	if err := j.validateSetEditorRoleValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"editorRoleValues",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetEmailAssertion(val *string) {
	if err := j.validateSetEmailAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAssertion",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetGroupsAssertion(val *string) {
	if err := j.validateSetGroupsAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsAssertion",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetIdpMetadataUrl(val *string) {
	if err := j.validateSetIdpMetadataUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpMetadataUrl",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetIdpMetadataXml(val *string) {
	if err := j.validateSetIdpMetadataXmlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpMetadataXml",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetLoginAssertion(val *string) {
	if err := j.validateSetLoginAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginAssertion",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetLoginValidityDuration(val *float64) {
	if err := j.validateSetLoginValidityDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginValidityDuration",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetNameAssertion(val *string) {
	if err := j.validateSetNameAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameAssertion",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetOrgAssertion(val *string) {
	if err := j.validateSetOrgAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orgAssertion",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetRoleAssertion(val *string) {
	if err := j.validateSetRoleAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleAssertion",
		val,
	)
}

func (j *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration)SetWorkspaceId(val *string) {
	if err := j.validateSetWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceId",
		val,
	)
}

// Generates CDKTN code for importing a AwsGrafanaWorkspaceSamlConfiguration resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsGrafanaWorkspaceSamlConfiguration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsGrafanaWorkspaceSamlConfiguration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-managed-grafana.AwsGrafanaWorkspaceSamlConfiguration",
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
func AwsGrafanaWorkspaceSamlConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsGrafanaWorkspaceSamlConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-managed-grafana.AwsGrafanaWorkspaceSamlConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsGrafanaWorkspaceSamlConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsGrafanaWorkspaceSamlConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-managed-grafana.AwsGrafanaWorkspaceSamlConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsGrafanaWorkspaceSamlConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsGrafanaWorkspaceSamlConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-managed-grafana.AwsGrafanaWorkspaceSamlConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsGrafanaWorkspaceSamlConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-managed-grafana.AwsGrafanaWorkspaceSamlConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := a.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) PutTimeouts(value *AwsGrafanaWorkspaceSamlConfiguration_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ResetAdminRoleValues() {
	_jsii_.InvokeVoid(
		a,
		"resetAdminRoleValues",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ResetAllowedOrganizations() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedOrganizations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ResetEmailAssertion() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailAssertion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ResetGroupsAssertion() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupsAssertion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ResetIdpMetadataUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetIdpMetadataUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ResetIdpMetadataXml() {
	_jsii_.InvokeVoid(
		a,
		"resetIdpMetadataXml",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ResetLoginAssertion() {
	_jsii_.InvokeVoid(
		a,
		"resetLoginAssertion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ResetLoginValidityDuration() {
	_jsii_.InvokeVoid(
		a,
		"resetLoginValidityDuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ResetNameAssertion() {
	_jsii_.InvokeVoid(
		a,
		"resetNameAssertion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ResetOrgAssertion() {
	_jsii_.InvokeVoid(
		a,
		"resetOrgAssertion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ResetRoleAssertion() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleAssertion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGrafanaWorkspaceSamlConfiguration) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

