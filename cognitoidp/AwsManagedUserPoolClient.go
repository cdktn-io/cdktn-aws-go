package cognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cognitoidp/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/cognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_managed_user_pool_client aws_cognito_managed_user_pool_client}.
// Experimental.
type AwsManagedUserPoolClient interface {
	cdktn.TerraformResource
	// Experimental.
	AccessTokenValidity() *float64
	// Experimental.
	SetAccessTokenValidity(val *float64)
	// Experimental.
	AccessTokenValidityInput() *float64
	// Experimental.
	AllowedOauthFlows() *[]*string
	// Experimental.
	SetAllowedOauthFlows(val *[]*string)
	// Experimental.
	AllowedOauthFlowsInput() *[]*string
	// Experimental.
	AllowedOauthFlowsUserPoolClient() interface{}
	// Experimental.
	SetAllowedOauthFlowsUserPoolClient(val interface{})
	// Experimental.
	AllowedOauthFlowsUserPoolClientInput() interface{}
	// Experimental.
	AllowedOauthScopes() *[]*string
	// Experimental.
	SetAllowedOauthScopes(val *[]*string)
	// Experimental.
	AllowedOauthScopesInput() *[]*string
	// Experimental.
	AnalyticsConfiguration() AwsManagedUserPoolClient_AnalyticsConfigurationPropertyList
	// Experimental.
	AnalyticsConfigurationInput() interface{}
	// Experimental.
	AuthSessionValidity() *float64
	// Experimental.
	SetAuthSessionValidity(val *float64)
	// Experimental.
	AuthSessionValidityInput() *float64
	// Experimental.
	CallbackUrls() *[]*string
	// Experimental.
	SetCallbackUrls(val *[]*string)
	// Experimental.
	CallbackUrlsInput() *[]*string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ClientSecret() *string
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
	DefaultRedirectUri() *string
	// Experimental.
	SetDefaultRedirectUri(val *string)
	// Experimental.
	DefaultRedirectUriInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EnablePropagateAdditionalUserContextData() interface{}
	// Experimental.
	SetEnablePropagateAdditionalUserContextData(val interface{})
	// Experimental.
	EnablePropagateAdditionalUserContextDataInput() interface{}
	// Experimental.
	EnableTokenRevocation() interface{}
	// Experimental.
	SetEnableTokenRevocation(val interface{})
	// Experimental.
	EnableTokenRevocationInput() interface{}
	// Experimental.
	ExplicitAuthFlows() *[]*string
	// Experimental.
	SetExplicitAuthFlows(val *[]*string)
	// Experimental.
	ExplicitAuthFlowsInput() *[]*string
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
	IdTokenValidity() *float64
	// Experimental.
	SetIdTokenValidity(val *float64)
	// Experimental.
	IdTokenValidityInput() *float64
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LogoutUrls() *[]*string
	// Experimental.
	SetLogoutUrls(val *[]*string)
	// Experimental.
	LogoutUrlsInput() *[]*string
	// Experimental.
	Name() *string
	// Experimental.
	NamePattern() *string
	// Experimental.
	SetNamePattern(val *string)
	// Experimental.
	NamePatternInput() *string
	// Experimental.
	NamePrefix() *string
	// Experimental.
	SetNamePrefix(val *string)
	// Experimental.
	NamePrefixInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PreventUserExistenceErrors() *string
	// Experimental.
	SetPreventUserExistenceErrors(val *string)
	// Experimental.
	PreventUserExistenceErrorsInput() *string
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
	ReadAttributes() *[]*string
	// Experimental.
	SetReadAttributes(val *[]*string)
	// Experimental.
	ReadAttributesInput() *[]*string
	// Experimental.
	RefreshTokenRotation() AwsManagedUserPoolClient_RefreshTokenRotationPropertyList
	// Experimental.
	RefreshTokenRotationInput() interface{}
	// Experimental.
	RefreshTokenValidity() *float64
	// Experimental.
	SetRefreshTokenValidity(val *float64)
	// Experimental.
	RefreshTokenValidityInput() *float64
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	SupportedIdentityProviders() *[]*string
	// Experimental.
	SetSupportedIdentityProviders(val *[]*string)
	// Experimental.
	SupportedIdentityProvidersInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	TokenValidityUnits() AwsManagedUserPoolClient_TokenValidityUnitsPropertyList
	// Experimental.
	TokenValidityUnitsInput() interface{}
	// Experimental.
	UserPoolId() *string
	// Experimental.
	SetUserPoolId(val *string)
	// Experimental.
	UserPoolIdInput() *string
	// Experimental.
	WriteAttributes() *[]*string
	// Experimental.
	SetWriteAttributes(val *[]*string)
	// Experimental.
	WriteAttributesInput() *[]*string
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
	PutAnalyticsConfiguration(value interface{})
	// Experimental.
	PutRefreshTokenRotation(value interface{})
	// Experimental.
	PutTokenValidityUnits(value interface{})
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
	ResetAccessTokenValidity()
	// Experimental.
	ResetAllowedOauthFlows()
	// Experimental.
	ResetAllowedOauthFlowsUserPoolClient()
	// Experimental.
	ResetAllowedOauthScopes()
	// Experimental.
	ResetAnalyticsConfiguration()
	// Experimental.
	ResetAuthSessionValidity()
	// Experimental.
	ResetCallbackUrls()
	// Experimental.
	ResetDefaultRedirectUri()
	// Experimental.
	ResetEnablePropagateAdditionalUserContextData()
	// Experimental.
	ResetEnableTokenRevocation()
	// Experimental.
	ResetExplicitAuthFlows()
	// Experimental.
	ResetIdTokenValidity()
	// Experimental.
	ResetLogoutUrls()
	// Experimental.
	ResetNamePattern()
	// Experimental.
	ResetNamePrefix()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPreventUserExistenceErrors()
	// Experimental.
	ResetReadAttributes()
	// Experimental.
	ResetRefreshTokenRotation()
	// Experimental.
	ResetRefreshTokenValidity()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSupportedIdentityProviders()
	// Experimental.
	ResetTokenValidityUnits()
	// Experimental.
	ResetWriteAttributes()
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

// The jsii proxy struct for AwsManagedUserPoolClient
type jsiiProxy_AwsManagedUserPoolClient struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsManagedUserPoolClient) AccessTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) AccessTokenValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessTokenValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) AllowedOauthFlows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOauthFlows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) AllowedOauthFlowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOauthFlowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) AllowedOauthFlowsUserPoolClient() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedOauthFlowsUserPoolClient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) AllowedOauthFlowsUserPoolClientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedOauthFlowsUserPoolClientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) AllowedOauthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) AllowedOauthScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOauthScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) AnalyticsConfiguration() AwsManagedUserPoolClient_AnalyticsConfigurationPropertyList {
	var returns AwsManagedUserPoolClient_AnalyticsConfigurationPropertyList
	_jsii_.Get(
		j,
		"analyticsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) AnalyticsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"analyticsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) AuthSessionValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authSessionValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) AuthSessionValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authSessionValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) CallbackUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"callbackUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) CallbackUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"callbackUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) DefaultRedirectUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRedirectUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) DefaultRedirectUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRedirectUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) EnablePropagateAdditionalUserContextData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePropagateAdditionalUserContextData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) EnablePropagateAdditionalUserContextDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePropagateAdditionalUserContextDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) EnableTokenRevocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTokenRevocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) EnableTokenRevocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTokenRevocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) ExplicitAuthFlows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"explicitAuthFlows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) ExplicitAuthFlowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"explicitAuthFlowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) IdTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) IdTokenValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idTokenValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) LogoutUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logoutUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) LogoutUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logoutUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) NamePattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) NamePatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) PreventUserExistenceErrors() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preventUserExistenceErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) PreventUserExistenceErrorsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preventUserExistenceErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) ReadAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) ReadAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) RefreshTokenRotation() AwsManagedUserPoolClient_RefreshTokenRotationPropertyList {
	var returns AwsManagedUserPoolClient_RefreshTokenRotationPropertyList
	_jsii_.Get(
		j,
		"refreshTokenRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) RefreshTokenRotationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"refreshTokenRotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) RefreshTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) RefreshTokenValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTokenValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) SupportedIdentityProviders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedIdentityProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) SupportedIdentityProvidersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedIdentityProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) TokenValidityUnits() AwsManagedUserPoolClient_TokenValidityUnitsPropertyList {
	var returns AwsManagedUserPoolClient_TokenValidityUnitsPropertyList
	_jsii_.Get(
		j,
		"tokenValidityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) TokenValidityUnitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenValidityUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) UserPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) WriteAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"writeAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsManagedUserPoolClient) WriteAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"writeAttributesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_managed_user_pool_client aws_cognito_managed_user_pool_client} Resource.
// Experimental.
func NewAwsManagedUserPoolClient(scope constructs.Construct, id *string, config *AwsManagedUserPoolClientConfig) AwsManagedUserPoolClient {
	_init_.Initialize()

	if err := validateNewAwsManagedUserPoolClientParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsManagedUserPoolClient{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsManagedUserPoolClient",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_managed_user_pool_client aws_cognito_managed_user_pool_client} Resource.
// Experimental.
func NewAwsManagedUserPoolClient_Override(a AwsManagedUserPoolClient, scope constructs.Construct, id *string, config *AwsManagedUserPoolClientConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsManagedUserPoolClient",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetAccessTokenValidity(val *float64) {
	if err := j.validateSetAccessTokenValidityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessTokenValidity",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetAllowedOauthFlows(val *[]*string) {
	if err := j.validateSetAllowedOauthFlowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOauthFlows",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetAllowedOauthFlowsUserPoolClient(val interface{}) {
	if err := j.validateSetAllowedOauthFlowsUserPoolClientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOauthFlowsUserPoolClient",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetAllowedOauthScopes(val *[]*string) {
	if err := j.validateSetAllowedOauthScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOauthScopes",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetAuthSessionValidity(val *float64) {
	if err := j.validateSetAuthSessionValidityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authSessionValidity",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetCallbackUrls(val *[]*string) {
	if err := j.validateSetCallbackUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"callbackUrls",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetDefaultRedirectUri(val *string) {
	if err := j.validateSetDefaultRedirectUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRedirectUri",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetEnablePropagateAdditionalUserContextData(val interface{}) {
	if err := j.validateSetEnablePropagateAdditionalUserContextDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePropagateAdditionalUserContextData",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetEnableTokenRevocation(val interface{}) {
	if err := j.validateSetEnableTokenRevocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableTokenRevocation",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetExplicitAuthFlows(val *[]*string) {
	if err := j.validateSetExplicitAuthFlowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"explicitAuthFlows",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetIdTokenValidity(val *float64) {
	if err := j.validateSetIdTokenValidityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idTokenValidity",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetLogoutUrls(val *[]*string) {
	if err := j.validateSetLogoutUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logoutUrls",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetNamePattern(val *string) {
	if err := j.validateSetNamePatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePattern",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetPreventUserExistenceErrors(val *string) {
	if err := j.validateSetPreventUserExistenceErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventUserExistenceErrors",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetReadAttributes(val *[]*string) {
	if err := j.validateSetReadAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readAttributes",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetRefreshTokenValidity(val *float64) {
	if err := j.validateSetRefreshTokenValidityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshTokenValidity",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetSupportedIdentityProviders(val *[]*string) {
	if err := j.validateSetSupportedIdentityProvidersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedIdentityProviders",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetUserPoolId(val *string) {
	if err := j.validateSetUserPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

func (j *jsiiProxy_AwsManagedUserPoolClient)SetWriteAttributes(val *[]*string) {
	if err := j.validateSetWriteAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeAttributes",
		val,
	)
}

// Generates CDKTN code for importing a AwsManagedUserPoolClient resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsManagedUserPoolClient_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsManagedUserPoolClient_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-cognito-idp.AwsManagedUserPoolClient",
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
func AwsManagedUserPoolClient_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsManagedUserPoolClient_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cognito-idp.AwsManagedUserPoolClient",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsManagedUserPoolClient_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsManagedUserPoolClient_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cognito-idp.AwsManagedUserPoolClient",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsManagedUserPoolClient_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsManagedUserPoolClient_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cognito-idp.AwsManagedUserPoolClient",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsManagedUserPoolClient_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-cognito-idp.AwsManagedUserPoolClient",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsManagedUserPoolClient) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsManagedUserPoolClient) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsManagedUserPoolClient) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsManagedUserPoolClient) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsManagedUserPoolClient) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsManagedUserPoolClient) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsManagedUserPoolClient) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsManagedUserPoolClient) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsManagedUserPoolClient) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsManagedUserPoolClient) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsManagedUserPoolClient) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsManagedUserPoolClient) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) PutAnalyticsConfiguration(value interface{}) {
	if err := a.validatePutAnalyticsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAnalyticsConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) PutRefreshTokenRotation(value interface{}) {
	if err := a.validatePutRefreshTokenRotationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRefreshTokenRotation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) PutTokenValidityUnits(value interface{}) {
	if err := a.validatePutTokenValidityUnitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTokenValidityUnits",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetAccessTokenValidity() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessTokenValidity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetAllowedOauthFlows() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedOauthFlows",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetAllowedOauthFlowsUserPoolClient() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedOauthFlowsUserPoolClient",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetAllowedOauthScopes() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedOauthScopes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetAnalyticsConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAnalyticsConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetAuthSessionValidity() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthSessionValidity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetCallbackUrls() {
	_jsii_.InvokeVoid(
		a,
		"resetCallbackUrls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetDefaultRedirectUri() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultRedirectUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetEnablePropagateAdditionalUserContextData() {
	_jsii_.InvokeVoid(
		a,
		"resetEnablePropagateAdditionalUserContextData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetEnableTokenRevocation() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableTokenRevocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetExplicitAuthFlows() {
	_jsii_.InvokeVoid(
		a,
		"resetExplicitAuthFlows",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetIdTokenValidity() {
	_jsii_.InvokeVoid(
		a,
		"resetIdTokenValidity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetLogoutUrls() {
	_jsii_.InvokeVoid(
		a,
		"resetLogoutUrls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetNamePattern() {
	_jsii_.InvokeVoid(
		a,
		"resetNamePattern",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetPreventUserExistenceErrors() {
	_jsii_.InvokeVoid(
		a,
		"resetPreventUserExistenceErrors",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetReadAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetReadAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetRefreshTokenRotation() {
	_jsii_.InvokeVoid(
		a,
		"resetRefreshTokenRotation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetRefreshTokenValidity() {
	_jsii_.InvokeVoid(
		a,
		"resetRefreshTokenValidity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetSupportedIdentityProviders() {
	_jsii_.InvokeVoid(
		a,
		"resetSupportedIdentityProviders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetTokenValidityUnits() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenValidityUnits",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ResetWriteAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetWriteAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsManagedUserPoolClient) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsManagedUserPoolClient) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsManagedUserPoolClient) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsManagedUserPoolClient) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

