package awsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsworkspacesweb/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsworkspacesweb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings aws_workspacesweb_user_settings}.
// Experimental.
type AwsWorkspaceswebUserSettings interface {
	cdktn.TerraformResource
	// Experimental.
	AdditionalEncryptionContext() *map[string]*string
	// Experimental.
	SetAdditionalEncryptionContext(val *map[string]*string)
	// Experimental.
	AdditionalEncryptionContextInput() *map[string]*string
	// Experimental.
	AssociatedPortalArns() *[]*string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	CookieSynchronizationConfiguration() AwsWorkspaceswebUserSettings_CookieSynchronizationConfigurationPropertyList
	// Experimental.
	CookieSynchronizationConfigurationInput() interface{}
	// Experimental.
	CopyAllowed() *string
	// Experimental.
	SetCopyAllowed(val *string)
	// Experimental.
	CopyAllowedInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	CustomerManagedKey() *string
	// Experimental.
	SetCustomerManagedKey(val *string)
	// Experimental.
	CustomerManagedKeyInput() *string
	// Experimental.
	DeepLinkAllowed() *string
	// Experimental.
	SetDeepLinkAllowed(val *string)
	// Experimental.
	DeepLinkAllowedInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DisconnectTimeoutInMinutes() *float64
	// Experimental.
	SetDisconnectTimeoutInMinutes(val *float64)
	// Experimental.
	DisconnectTimeoutInMinutesInput() *float64
	// Experimental.
	DownloadAllowed() *string
	// Experimental.
	SetDownloadAllowed(val *string)
	// Experimental.
	DownloadAllowedInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	IdleDisconnectTimeoutInMinutes() *float64
	// Experimental.
	SetIdleDisconnectTimeoutInMinutes(val *float64)
	// Experimental.
	IdleDisconnectTimeoutInMinutesInput() *float64
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PasteAllowed() *string
	// Experimental.
	SetPasteAllowed(val *string)
	// Experimental.
	PasteAllowedInput() *string
	// Experimental.
	PrintAllowed() *string
	// Experimental.
	SetPrintAllowed(val *string)
	// Experimental.
	PrintAllowedInput() *string
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
	Tags() *map[string]*string
	// Experimental.
	SetTags(val *map[string]*string)
	// Experimental.
	TagsAll() cdktn.StringMap
	// Experimental.
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	ToolbarConfiguration() AwsWorkspaceswebUserSettings_ToolbarConfigurationPropertyList
	// Experimental.
	ToolbarConfigurationInput() interface{}
	// Experimental.
	UploadAllowed() *string
	// Experimental.
	SetUploadAllowed(val *string)
	// Experimental.
	UploadAllowedInput() *string
	// Experimental.
	UserSettingsArn() *string
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
	PutCookieSynchronizationConfiguration(value interface{})
	// Experimental.
	PutToolbarConfiguration(value interface{})
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
	ResetAdditionalEncryptionContext()
	// Experimental.
	ResetCookieSynchronizationConfiguration()
	// Experimental.
	ResetCustomerManagedKey()
	// Experimental.
	ResetDeepLinkAllowed()
	// Experimental.
	ResetDisconnectTimeoutInMinutes()
	// Experimental.
	ResetIdleDisconnectTimeoutInMinutes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetToolbarConfiguration()
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

// The jsii proxy struct for AwsWorkspaceswebUserSettings
type jsiiProxy_AwsWorkspaceswebUserSettings struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) AdditionalEncryptionContext() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalEncryptionContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) AdditionalEncryptionContextInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalEncryptionContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) AssociatedPortalArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"associatedPortalArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) CookieSynchronizationConfiguration() AwsWorkspaceswebUserSettings_CookieSynchronizationConfigurationPropertyList {
	var returns AwsWorkspaceswebUserSettings_CookieSynchronizationConfigurationPropertyList
	_jsii_.Get(
		j,
		"cookieSynchronizationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) CookieSynchronizationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cookieSynchronizationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) CopyAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) CopyAllowedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) CustomerManagedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerManagedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) CustomerManagedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerManagedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) DeepLinkAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deepLinkAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) DeepLinkAllowedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deepLinkAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) DisconnectTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"disconnectTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) DisconnectTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"disconnectTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) DownloadAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"downloadAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) DownloadAllowedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"downloadAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) IdleDisconnectTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleDisconnectTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) IdleDisconnectTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleDisconnectTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) PasteAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pasteAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) PasteAllowedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pasteAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) PrintAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"printAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) PrintAllowedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"printAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) ToolbarConfiguration() AwsWorkspaceswebUserSettings_ToolbarConfigurationPropertyList {
	var returns AwsWorkspaceswebUserSettings_ToolbarConfigurationPropertyList
	_jsii_.Get(
		j,
		"toolbarConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) ToolbarConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"toolbarConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) UploadAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) UploadAllowedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings) UserSettingsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userSettingsArn",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings aws_workspacesweb_user_settings} Resource.
// Experimental.
func NewAwsWorkspaceswebUserSettings(scope constructs.Construct, id *string, config *AwsWorkspaceswebUserSettingsConfig) AwsWorkspaceswebUserSettings {
	_init_.Initialize()

	if err := validateNewAwsWorkspaceswebUserSettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWorkspaceswebUserSettings{}

	_jsii_.Create(
		"@cdktn/aws-workspaces-web.AwsWorkspaceswebUserSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings aws_workspacesweb_user_settings} Resource.
// Experimental.
func NewAwsWorkspaceswebUserSettings_Override(a AwsWorkspaceswebUserSettings, scope constructs.Construct, id *string, config *AwsWorkspaceswebUserSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-workspaces-web.AwsWorkspaceswebUserSettings",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings)SetAdditionalEncryptionContext(val *map[string]*string) {
	if err := j.validateSetAdditionalEncryptionContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalEncryptionContext",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings)SetCopyAllowed(val *string) {
	if err := j.validateSetCopyAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyAllowed",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings)SetCustomerManagedKey(val *string) {
	if err := j.validateSetCustomerManagedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerManagedKey",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings)SetDeepLinkAllowed(val *string) {
	if err := j.validateSetDeepLinkAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deepLinkAllowed",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings)SetDisconnectTimeoutInMinutes(val *float64) {
	if err := j.validateSetDisconnectTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disconnectTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings)SetDownloadAllowed(val *string) {
	if err := j.validateSetDownloadAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"downloadAllowed",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings)SetIdleDisconnectTimeoutInMinutes(val *float64) {
	if err := j.validateSetIdleDisconnectTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleDisconnectTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings)SetPasteAllowed(val *string) {
	if err := j.validateSetPasteAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pasteAllowed",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings)SetPrintAllowed(val *string) {
	if err := j.validateSetPrintAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"printAllowed",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspaceswebUserSettings)SetUploadAllowed(val *string) {
	if err := j.validateSetUploadAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uploadAllowed",
		val,
	)
}

// Generates CDKTN code for importing a AwsWorkspaceswebUserSettings resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsWorkspaceswebUserSettings_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsWorkspaceswebUserSettings_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-workspaces-web.AwsWorkspaceswebUserSettings",
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
func AwsWorkspaceswebUserSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsWorkspaceswebUserSettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-workspaces-web.AwsWorkspaceswebUserSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsWorkspaceswebUserSettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsWorkspaceswebUserSettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-workspaces-web.AwsWorkspaceswebUserSettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsWorkspaceswebUserSettings_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsWorkspaceswebUserSettings_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-workspaces-web.AwsWorkspaceswebUserSettings",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsWorkspaceswebUserSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-workspaces-web.AwsWorkspaceswebUserSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) PutCookieSynchronizationConfiguration(value interface{}) {
	if err := a.validatePutCookieSynchronizationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCookieSynchronizationConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) PutToolbarConfiguration(value interface{}) {
	if err := a.validatePutToolbarConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putToolbarConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) ResetAdditionalEncryptionContext() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalEncryptionContext",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) ResetCookieSynchronizationConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCookieSynchronizationConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) ResetCustomerManagedKey() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomerManagedKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) ResetDeepLinkAllowed() {
	_jsii_.InvokeVoid(
		a,
		"resetDeepLinkAllowed",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) ResetDisconnectTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetDisconnectTimeoutInMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) ResetIdleDisconnectTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetIdleDisconnectTimeoutInMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) ResetToolbarConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetToolbarConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWorkspaceswebUserSettings) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

