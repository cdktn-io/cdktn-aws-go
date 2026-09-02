package awskms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskms/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awskms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store aws_kms_custom_key_store}.
// Experimental.
type TfCustomKeyStore interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CloudHsmClusterId() *string
	// Experimental.
	SetCloudHsmClusterId(val *string)
	// Experimental.
	CloudHsmClusterIdInput() *string
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
	CustomKeyStoreName() *string
	// Experimental.
	SetCustomKeyStoreName(val *string)
	// Experimental.
	CustomKeyStoreNameInput() *string
	// Experimental.
	CustomKeyStoreType() *string
	// Experimental.
	SetCustomKeyStoreType(val *string)
	// Experimental.
	CustomKeyStoreTypeInput() *string
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
	KeyStorePassword() *string
	// Experimental.
	SetKeyStorePassword(val *string)
	// Experimental.
	KeyStorePasswordInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() TfCustomKeyStore_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TrustAnchorCertificate() *string
	// Experimental.
	SetTrustAnchorCertificate(val *string)
	// Experimental.
	TrustAnchorCertificateInput() *string
	// Experimental.
	XksProxyAuthenticationCredential() TfCustomKeyStore_XksProxyAuthenticationCredentialPropertyOutputReference
	// Experimental.
	XksProxyAuthenticationCredentialInput() *TfCustomKeyStore_XksProxyAuthenticationCredentialProperty
	// Experimental.
	XksProxyConnectivity() *string
	// Experimental.
	SetXksProxyConnectivity(val *string)
	// Experimental.
	XksProxyConnectivityInput() *string
	// Experimental.
	XksProxyUriEndpoint() *string
	// Experimental.
	SetXksProxyUriEndpoint(val *string)
	// Experimental.
	XksProxyUriEndpointInput() *string
	// Experimental.
	XksProxyUriPath() *string
	// Experimental.
	SetXksProxyUriPath(val *string)
	// Experimental.
	XksProxyUriPathInput() *string
	// Experimental.
	XksProxyVpcEndpointServiceName() *string
	// Experimental.
	SetXksProxyVpcEndpointServiceName(val *string)
	// Experimental.
	XksProxyVpcEndpointServiceNameInput() *string
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
	PutTimeouts(value *TfCustomKeyStore_TimeoutsProperty)
	// Experimental.
	PutXksProxyAuthenticationCredential(value *TfCustomKeyStore_XksProxyAuthenticationCredentialProperty)
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
	ResetCloudHsmClusterId()
	// Experimental.
	ResetCustomKeyStoreType()
	// Experimental.
	ResetId()
	// Experimental.
	ResetKeyStorePassword()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTrustAnchorCertificate()
	// Experimental.
	ResetXksProxyAuthenticationCredential()
	// Experimental.
	ResetXksProxyConnectivity()
	// Experimental.
	ResetXksProxyUriEndpoint()
	// Experimental.
	ResetXksProxyUriPath()
	// Experimental.
	ResetXksProxyVpcEndpointServiceName()
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

// The jsii proxy struct for TfCustomKeyStore
type jsiiProxy_TfCustomKeyStore struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfCustomKeyStore) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) CloudHsmClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudHsmClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) CloudHsmClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudHsmClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) CustomKeyStoreName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customKeyStoreName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) CustomKeyStoreNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customKeyStoreNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) CustomKeyStoreType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customKeyStoreType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) CustomKeyStoreTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customKeyStoreTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) KeyStorePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStorePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) KeyStorePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStorePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) Timeouts() TfCustomKeyStore_TimeoutsPropertyOutputReference {
	var returns TfCustomKeyStore_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) TrustAnchorCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustAnchorCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) TrustAnchorCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustAnchorCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) XksProxyAuthenticationCredential() TfCustomKeyStore_XksProxyAuthenticationCredentialPropertyOutputReference {
	var returns TfCustomKeyStore_XksProxyAuthenticationCredentialPropertyOutputReference
	_jsii_.Get(
		j,
		"xksProxyAuthenticationCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) XksProxyAuthenticationCredentialInput() *TfCustomKeyStore_XksProxyAuthenticationCredentialProperty {
	var returns *TfCustomKeyStore_XksProxyAuthenticationCredentialProperty
	_jsii_.Get(
		j,
		"xksProxyAuthenticationCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) XksProxyConnectivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xksProxyConnectivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) XksProxyConnectivityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xksProxyConnectivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) XksProxyUriEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xksProxyUriEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) XksProxyUriEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xksProxyUriEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) XksProxyUriPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xksProxyUriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) XksProxyUriPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xksProxyUriPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) XksProxyVpcEndpointServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xksProxyVpcEndpointServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomKeyStore) XksProxyVpcEndpointServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xksProxyVpcEndpointServiceNameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store aws_kms_custom_key_store} Resource.
// Experimental.
func NewTfCustomKeyStore(scope constructs.Construct, id *string, config *TfCustomKeyStoreConfig) TfCustomKeyStore {
	_init_.Initialize()

	if err := validateNewTfCustomKeyStoreParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCustomKeyStore{}

	_jsii_.Create(
		"@cdktn/aws-kms.TfCustomKeyStore",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store aws_kms_custom_key_store} Resource.
// Experimental.
func NewTfCustomKeyStore_Override(t TfCustomKeyStore, scope constructs.Construct, id *string, config *TfCustomKeyStoreConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kms.TfCustomKeyStore",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfCustomKeyStore)SetCloudHsmClusterId(val *string) {
	if err := j.validateSetCloudHsmClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudHsmClusterId",
		val,
	)
}

func (j *jsiiProxy_TfCustomKeyStore)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfCustomKeyStore)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfCustomKeyStore)SetCustomKeyStoreName(val *string) {
	if err := j.validateSetCustomKeyStoreNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customKeyStoreName",
		val,
	)
}

func (j *jsiiProxy_TfCustomKeyStore)SetCustomKeyStoreType(val *string) {
	if err := j.validateSetCustomKeyStoreTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customKeyStoreType",
		val,
	)
}

func (j *jsiiProxy_TfCustomKeyStore)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfCustomKeyStore)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfCustomKeyStore)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfCustomKeyStore)SetKeyStorePassword(val *string) {
	if err := j.validateSetKeyStorePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyStorePassword",
		val,
	)
}

func (j *jsiiProxy_TfCustomKeyStore)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfCustomKeyStore)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfCustomKeyStore)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfCustomKeyStore)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfCustomKeyStore)SetTrustAnchorCertificate(val *string) {
	if err := j.validateSetTrustAnchorCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustAnchorCertificate",
		val,
	)
}

func (j *jsiiProxy_TfCustomKeyStore)SetXksProxyConnectivity(val *string) {
	if err := j.validateSetXksProxyConnectivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xksProxyConnectivity",
		val,
	)
}

func (j *jsiiProxy_TfCustomKeyStore)SetXksProxyUriEndpoint(val *string) {
	if err := j.validateSetXksProxyUriEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xksProxyUriEndpoint",
		val,
	)
}

func (j *jsiiProxy_TfCustomKeyStore)SetXksProxyUriPath(val *string) {
	if err := j.validateSetXksProxyUriPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xksProxyUriPath",
		val,
	)
}

func (j *jsiiProxy_TfCustomKeyStore)SetXksProxyVpcEndpointServiceName(val *string) {
	if err := j.validateSetXksProxyVpcEndpointServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xksProxyVpcEndpointServiceName",
		val,
	)
}

// Generates CDKTN code for importing a TfCustomKeyStore resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfCustomKeyStore_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfCustomKeyStore_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-kms.TfCustomKeyStore",
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
func TfCustomKeyStore_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCustomKeyStore_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-kms.TfCustomKeyStore",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfCustomKeyStore_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCustomKeyStore_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-kms.TfCustomKeyStore",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfCustomKeyStore_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCustomKeyStore_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-kms.TfCustomKeyStore",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfCustomKeyStore_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-kms.TfCustomKeyStore",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfCustomKeyStore) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfCustomKeyStore) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfCustomKeyStore) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCustomKeyStore) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCustomKeyStore) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCustomKeyStore) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCustomKeyStore) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCustomKeyStore) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCustomKeyStore) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCustomKeyStore) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCustomKeyStore) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCustomKeyStore) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomKeyStore) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfCustomKeyStore) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCustomKeyStore) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfCustomKeyStore) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfCustomKeyStore) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfCustomKeyStore) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfCustomKeyStore) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfCustomKeyStore) PutTimeouts(value *TfCustomKeyStore_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCustomKeyStore) PutXksProxyAuthenticationCredential(value *TfCustomKeyStore_XksProxyAuthenticationCredentialProperty) {
	if err := t.validatePutXksProxyAuthenticationCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putXksProxyAuthenticationCredential",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCustomKeyStore) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfCustomKeyStore) ResetCloudHsmClusterId() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudHsmClusterId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomKeyStore) ResetCustomKeyStoreType() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomKeyStoreType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomKeyStore) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomKeyStore) ResetKeyStorePassword() {
	_jsii_.InvokeVoid(
		t,
		"resetKeyStorePassword",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomKeyStore) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomKeyStore) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomKeyStore) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomKeyStore) ResetTrustAnchorCertificate() {
	_jsii_.InvokeVoid(
		t,
		"resetTrustAnchorCertificate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomKeyStore) ResetXksProxyAuthenticationCredential() {
	_jsii_.InvokeVoid(
		t,
		"resetXksProxyAuthenticationCredential",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomKeyStore) ResetXksProxyConnectivity() {
	_jsii_.InvokeVoid(
		t,
		"resetXksProxyConnectivity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomKeyStore) ResetXksProxyUriEndpoint() {
	_jsii_.InvokeVoid(
		t,
		"resetXksProxyUriEndpoint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomKeyStore) ResetXksProxyUriPath() {
	_jsii_.InvokeVoid(
		t,
		"resetXksProxyUriPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomKeyStore) ResetXksProxyVpcEndpointServiceName() {
	_jsii_.InvokeVoid(
		t,
		"resetXksProxyVpcEndpointServiceName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomKeyStore) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomKeyStore) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomKeyStore) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomKeyStore) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomKeyStore) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomKeyStore) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomKeyStore) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

