package awsmanagedgrafana

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmanagedgrafana/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsmanagedgrafana/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace aws_grafana_workspace}.
// Experimental.
type TfWorkspace interface {
	cdktn.TerraformResource
	// Experimental.
	AccountAccessType() *string
	// Experimental.
	SetAccountAccessType(val *string)
	// Experimental.
	AccountAccessTypeInput() *string
	// Experimental.
	Arn() *string
	// Experimental.
	AuthenticationProviders() *[]*string
	// Experimental.
	SetAuthenticationProviders(val *[]*string)
	// Experimental.
	AuthenticationProvidersInput() *[]*string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Configuration() *string
	// Experimental.
	SetConfiguration(val *string)
	// Experimental.
	ConfigurationInput() *string
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
	DataSources() *[]*string
	// Experimental.
	SetDataSources(val *[]*string)
	// Experimental.
	DataSourcesInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	Endpoint() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	GrafanaVersion() *string
	// Experimental.
	SetGrafanaVersion(val *string)
	// Experimental.
	GrafanaVersionInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	KmsKeyId() *string
	// Experimental.
	SetKmsKeyId(val *string)
	// Experimental.
	KmsKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NetworkAccessControl() TfWorkspace_NetworkAccessControlPropertyOutputReference
	// Experimental.
	NetworkAccessControlInput() *TfWorkspace_NetworkAccessControlProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	NotificationDestinations() *[]*string
	// Experimental.
	SetNotificationDestinations(val *[]*string)
	// Experimental.
	NotificationDestinationsInput() *[]*string
	// Experimental.
	OrganizationalUnits() *[]*string
	// Experimental.
	SetOrganizationalUnits(val *[]*string)
	// Experimental.
	OrganizationalUnitsInput() *[]*string
	// Experimental.
	OrganizationRoleName() *string
	// Experimental.
	SetOrganizationRoleName(val *string)
	// Experimental.
	OrganizationRoleNameInput() *string
	// Experimental.
	PermissionType() *string
	// Experimental.
	SetPermissionType(val *string)
	// Experimental.
	PermissionTypeInput() *string
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
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	SamlConfigurationStatus() *string
	// Experimental.
	StackSetName() *string
	// Experimental.
	SetStackSetName(val *string)
	// Experimental.
	StackSetNameInput() *string
	// Experimental.
	Tags() *map[string]*string
	// Experimental.
	SetTags(val *map[string]*string)
	// Experimental.
	TagsAll() *map[string]*string
	// Experimental.
	SetTagsAll(val *map[string]*string)
	// Experimental.
	TagsAllInput() *map[string]*string
	// Experimental.
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() TfWorkspace_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	VpcConfiguration() TfWorkspace_VpcConfigurationPropertyOutputReference
	// Experimental.
	VpcConfigurationInput() *TfWorkspace_VpcConfigurationProperty
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
	PutNetworkAccessControl(value *TfWorkspace_NetworkAccessControlProperty)
	// Experimental.
	PutTimeouts(value *TfWorkspace_TimeoutsProperty)
	// Experimental.
	PutVpcConfiguration(value *TfWorkspace_VpcConfigurationProperty)
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
	ResetConfiguration()
	// Experimental.
	ResetDataSources()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetGrafanaVersion()
	// Experimental.
	ResetId()
	// Experimental.
	ResetKmsKeyId()
	// Experimental.
	ResetName()
	// Experimental.
	ResetNetworkAccessControl()
	// Experimental.
	ResetNotificationDestinations()
	// Experimental.
	ResetOrganizationalUnits()
	// Experimental.
	ResetOrganizationRoleName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRoleArn()
	// Experimental.
	ResetStackSetName()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetVpcConfiguration()
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

// The jsii proxy struct for TfWorkspace
type jsiiProxy_TfWorkspace struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfWorkspace) AccountAccessType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountAccessType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) AccountAccessTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountAccessTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) AuthenticationProviders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authenticationProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) AuthenticationProvidersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authenticationProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) Configuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) ConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) DataSources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) DataSourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) GrafanaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) GrafanaVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) NetworkAccessControl() TfWorkspace_NetworkAccessControlPropertyOutputReference {
	var returns TfWorkspace_NetworkAccessControlPropertyOutputReference
	_jsii_.Get(
		j,
		"networkAccessControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) NetworkAccessControlInput() *TfWorkspace_NetworkAccessControlProperty {
	var returns *TfWorkspace_NetworkAccessControlProperty
	_jsii_.Get(
		j,
		"networkAccessControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) NotificationDestinations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) NotificationDestinationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationDestinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) OrganizationalUnits() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationalUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) OrganizationalUnitsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationalUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) OrganizationRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) OrganizationRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) PermissionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) PermissionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) SamlConfigurationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlConfigurationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) StackSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) StackSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) Timeouts() TfWorkspace_TimeoutsPropertyOutputReference {
	var returns TfWorkspace_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) VpcConfiguration() TfWorkspace_VpcConfigurationPropertyOutputReference {
	var returns TfWorkspace_VpcConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"vpcConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkspace) VpcConfigurationInput() *TfWorkspace_VpcConfigurationProperty {
	var returns *TfWorkspace_VpcConfigurationProperty
	_jsii_.Get(
		j,
		"vpcConfigurationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace aws_grafana_workspace} Resource.
// Experimental.
func NewTfWorkspace(scope constructs.Construct, id *string, config *TfWorkspaceConfig) TfWorkspace {
	_init_.Initialize()

	if err := validateNewTfWorkspaceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWorkspace{}

	_jsii_.Create(
		"@cdktn/aws-managed-grafana.TfWorkspace",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace aws_grafana_workspace} Resource.
// Experimental.
func NewTfWorkspace_Override(t TfWorkspace, scope constructs.Construct, id *string, config *TfWorkspaceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-managed-grafana.TfWorkspace",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfWorkspace)SetAccountAccessType(val *string) {
	if err := j.validateSetAccountAccessTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountAccessType",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetAuthenticationProviders(val *[]*string) {
	if err := j.validateSetAuthenticationProvidersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationProviders",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetConfiguration(val *string) {
	if err := j.validateSetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetDataSources(val *[]*string) {
	if err := j.validateSetDataSourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSources",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetGrafanaVersion(val *string) {
	if err := j.validateSetGrafanaVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grafanaVersion",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetNotificationDestinations(val *[]*string) {
	if err := j.validateSetNotificationDestinationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationDestinations",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetOrganizationalUnits(val *[]*string) {
	if err := j.validateSetOrganizationalUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationalUnits",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetOrganizationRoleName(val *string) {
	if err := j.validateSetOrganizationRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationRoleName",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetPermissionType(val *string) {
	if err := j.validateSetPermissionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissionType",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetStackSetName(val *string) {
	if err := j.validateSetStackSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackSetName",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfWorkspace)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a TfWorkspace resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfWorkspace_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfWorkspace_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-managed-grafana.TfWorkspace",
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
func TfWorkspace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfWorkspace_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-managed-grafana.TfWorkspace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfWorkspace_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfWorkspace_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-managed-grafana.TfWorkspace",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfWorkspace_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfWorkspace_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-managed-grafana.TfWorkspace",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfWorkspace_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-managed-grafana.TfWorkspace",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfWorkspace) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfWorkspace) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfWorkspace) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWorkspace) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkspace) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWorkspace) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWorkspace) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWorkspace) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWorkspace) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWorkspace) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWorkspace) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWorkspace) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkspace) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfWorkspace) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkspace) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfWorkspace) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfWorkspace) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfWorkspace) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfWorkspace) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfWorkspace) PutNetworkAccessControl(value *TfWorkspace_NetworkAccessControlProperty) {
	if err := t.validatePutNetworkAccessControlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkAccessControl",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkspace) PutTimeouts(value *TfWorkspace_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkspace) PutVpcConfiguration(value *TfWorkspace_VpcConfigurationProperty) {
	if err := t.validatePutVpcConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVpcConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkspace) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfWorkspace) ResetConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspace) ResetDataSources() {
	_jsii_.InvokeVoid(
		t,
		"resetDataSources",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspace) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspace) ResetGrafanaVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetGrafanaVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspace) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspace) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspace) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspace) ResetNetworkAccessControl() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkAccessControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspace) ResetNotificationDestinations() {
	_jsii_.InvokeVoid(
		t,
		"resetNotificationDestinations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspace) ResetOrganizationalUnits() {
	_jsii_.InvokeVoid(
		t,
		"resetOrganizationalUnits",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspace) ResetOrganizationRoleName() {
	_jsii_.InvokeVoid(
		t,
		"resetOrganizationRoleName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspace) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspace) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspace) ResetRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspace) ResetStackSetName() {
	_jsii_.InvokeVoid(
		t,
		"resetStackSetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspace) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspace) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspace) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspace) ResetVpcConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkspace) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkspace) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkspace) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkspace) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkspace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkspace) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkspace) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

