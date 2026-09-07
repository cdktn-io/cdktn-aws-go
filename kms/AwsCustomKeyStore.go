package kms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/kms/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/kms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store aws_kms_custom_key_store}.
// Experimental.
type AwsCustomKeyStore interface {
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
	Timeouts() AwsCustomKeyStore_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TrustAnchorCertificate() *string
	// Experimental.
	SetTrustAnchorCertificate(val *string)
	// Experimental.
	TrustAnchorCertificateInput() *string
	// Experimental.
	XksProxyAuthenticationCredential() AwsCustomKeyStore_XksProxyAuthenticationCredentialPropertyOutputReference
	// Experimental.
	XksProxyAuthenticationCredentialInput() *AwsCustomKeyStore_XksProxyAuthenticationCredentialProperty
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
	PutTimeouts(value *AwsCustomKeyStore_TimeoutsProperty)
	// Experimental.
	PutXksProxyAuthenticationCredential(value *AwsCustomKeyStore_XksProxyAuthenticationCredentialProperty)
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

// The jsii proxy struct for AwsCustomKeyStore
type jsiiProxy_AwsCustomKeyStore struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsCustomKeyStore) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) CloudHsmClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudHsmClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) CloudHsmClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudHsmClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) CustomKeyStoreName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customKeyStoreName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) CustomKeyStoreNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customKeyStoreNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) CustomKeyStoreType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customKeyStoreType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) CustomKeyStoreTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customKeyStoreTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) KeyStorePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStorePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) KeyStorePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStorePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) Timeouts() AwsCustomKeyStore_TimeoutsPropertyOutputReference {
	var returns AwsCustomKeyStore_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) TrustAnchorCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustAnchorCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) TrustAnchorCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustAnchorCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) XksProxyAuthenticationCredential() AwsCustomKeyStore_XksProxyAuthenticationCredentialPropertyOutputReference {
	var returns AwsCustomKeyStore_XksProxyAuthenticationCredentialPropertyOutputReference
	_jsii_.Get(
		j,
		"xksProxyAuthenticationCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) XksProxyAuthenticationCredentialInput() *AwsCustomKeyStore_XksProxyAuthenticationCredentialProperty {
	var returns *AwsCustomKeyStore_XksProxyAuthenticationCredentialProperty
	_jsii_.Get(
		j,
		"xksProxyAuthenticationCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) XksProxyConnectivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xksProxyConnectivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) XksProxyConnectivityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xksProxyConnectivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) XksProxyUriEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xksProxyUriEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) XksProxyUriEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xksProxyUriEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) XksProxyUriPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xksProxyUriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) XksProxyUriPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xksProxyUriPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) XksProxyVpcEndpointServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xksProxyVpcEndpointServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomKeyStore) XksProxyVpcEndpointServiceNameInput() *string {
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
func NewAwsCustomKeyStore(scope constructs.Construct, id *string, config *AwsCustomKeyStoreConfig) AwsCustomKeyStore {
	_init_.Initialize()

	if err := validateNewAwsCustomKeyStoreParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCustomKeyStore{}

	_jsii_.Create(
		"@cdktn/aws-kms.AwsCustomKeyStore",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store aws_kms_custom_key_store} Resource.
// Experimental.
func NewAwsCustomKeyStore_Override(a AwsCustomKeyStore, scope constructs.Construct, id *string, config *AwsCustomKeyStoreConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kms.AwsCustomKeyStore",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsCustomKeyStore)SetCloudHsmClusterId(val *string) {
	if err := j.validateSetCloudHsmClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudHsmClusterId",
		val,
	)
}

func (j *jsiiProxy_AwsCustomKeyStore)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsCustomKeyStore)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsCustomKeyStore)SetCustomKeyStoreName(val *string) {
	if err := j.validateSetCustomKeyStoreNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customKeyStoreName",
		val,
	)
}

func (j *jsiiProxy_AwsCustomKeyStore)SetCustomKeyStoreType(val *string) {
	if err := j.validateSetCustomKeyStoreTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customKeyStoreType",
		val,
	)
}

func (j *jsiiProxy_AwsCustomKeyStore)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsCustomKeyStore)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsCustomKeyStore)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsCustomKeyStore)SetKeyStorePassword(val *string) {
	if err := j.validateSetKeyStorePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyStorePassword",
		val,
	)
}

func (j *jsiiProxy_AwsCustomKeyStore)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsCustomKeyStore)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsCustomKeyStore)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsCustomKeyStore)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsCustomKeyStore)SetTrustAnchorCertificate(val *string) {
	if err := j.validateSetTrustAnchorCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustAnchorCertificate",
		val,
	)
}

func (j *jsiiProxy_AwsCustomKeyStore)SetXksProxyConnectivity(val *string) {
	if err := j.validateSetXksProxyConnectivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xksProxyConnectivity",
		val,
	)
}

func (j *jsiiProxy_AwsCustomKeyStore)SetXksProxyUriEndpoint(val *string) {
	if err := j.validateSetXksProxyUriEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xksProxyUriEndpoint",
		val,
	)
}

func (j *jsiiProxy_AwsCustomKeyStore)SetXksProxyUriPath(val *string) {
	if err := j.validateSetXksProxyUriPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xksProxyUriPath",
		val,
	)
}

func (j *jsiiProxy_AwsCustomKeyStore)SetXksProxyVpcEndpointServiceName(val *string) {
	if err := j.validateSetXksProxyVpcEndpointServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xksProxyVpcEndpointServiceName",
		val,
	)
}

// Generates CDKTN code for importing a AwsCustomKeyStore resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsCustomKeyStore_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsCustomKeyStore_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-kms.AwsCustomKeyStore",
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
func AwsCustomKeyStore_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCustomKeyStore_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-kms.AwsCustomKeyStore",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsCustomKeyStore_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCustomKeyStore_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-kms.AwsCustomKeyStore",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsCustomKeyStore_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCustomKeyStore_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-kms.AwsCustomKeyStore",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsCustomKeyStore_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-kms.AwsCustomKeyStore",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsCustomKeyStore) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCustomKeyStore) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCustomKeyStore) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCustomKeyStore) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCustomKeyStore) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCustomKeyStore) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCustomKeyStore) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCustomKeyStore) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCustomKeyStore) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCustomKeyStore) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomKeyStore) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCustomKeyStore) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsCustomKeyStore) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) PutTimeouts(value *AwsCustomKeyStore_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) PutXksProxyAuthenticationCredential(value *AwsCustomKeyStore_XksProxyAuthenticationCredentialProperty) {
	if err := a.validatePutXksProxyAuthenticationCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putXksProxyAuthenticationCredential",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) ResetCloudHsmClusterId() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudHsmClusterId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) ResetCustomKeyStoreType() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomKeyStoreType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) ResetKeyStorePassword() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyStorePassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) ResetTrustAnchorCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetTrustAnchorCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) ResetXksProxyAuthenticationCredential() {
	_jsii_.InvokeVoid(
		a,
		"resetXksProxyAuthenticationCredential",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) ResetXksProxyConnectivity() {
	_jsii_.InvokeVoid(
		a,
		"resetXksProxyConnectivity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) ResetXksProxyUriEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetXksProxyUriEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) ResetXksProxyUriPath() {
	_jsii_.InvokeVoid(
		a,
		"resetXksProxyUriPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) ResetXksProxyVpcEndpointServiceName() {
	_jsii_.InvokeVoid(
		a,
		"resetXksProxyVpcEndpointServiceName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomKeyStore) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomKeyStore) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomKeyStore) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomKeyStore) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomKeyStore) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomKeyStore) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomKeyStore) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

