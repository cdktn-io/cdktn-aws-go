package awscognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscognitoidp/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awscognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool aws_cognito_user_pool}.
// Experimental.
type TfUserPool interface {
	cdktn.TerraformResource
	// Experimental.
	AccountRecoverySetting() TfUserPool_AccountRecoverySettingPropertyOutputReference
	// Experimental.
	AccountRecoverySettingInput() *TfUserPool_AccountRecoverySettingProperty
	// Experimental.
	AdminCreateUserConfig() TfUserPool_AdminCreateUserConfigPropertyOutputReference
	// Experimental.
	AdminCreateUserConfigInput() *TfUserPool_AdminCreateUserConfigProperty
	// Experimental.
	AliasAttributes() *[]*string
	// Experimental.
	SetAliasAttributes(val *[]*string)
	// Experimental.
	AliasAttributesInput() *[]*string
	// Experimental.
	Arn() *string
	// Experimental.
	AutoVerifiedAttributes() *[]*string
	// Experimental.
	SetAutoVerifiedAttributes(val *[]*string)
	// Experimental.
	AutoVerifiedAttributesInput() *[]*string
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
	CreationDate() *string
	// Experimental.
	CustomDomain() *string
	// Experimental.
	DeletionProtection() *string
	// Experimental.
	SetDeletionProtection(val *string)
	// Experimental.
	DeletionProtectionInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DeviceConfiguration() TfUserPool_DeviceConfigurationPropertyOutputReference
	// Experimental.
	DeviceConfigurationInput() *TfUserPool_DeviceConfigurationProperty
	// Experimental.
	Domain() *string
	// Experimental.
	EmailConfiguration() TfUserPool_EmailConfigurationPropertyOutputReference
	// Experimental.
	EmailConfigurationInput() *TfUserPool_EmailConfigurationProperty
	// Experimental.
	EmailMfaConfiguration() TfUserPool_EmailMfaConfigurationPropertyOutputReference
	// Experimental.
	EmailMfaConfigurationInput() *TfUserPool_EmailMfaConfigurationProperty
	// Experimental.
	EmailVerificationMessage() *string
	// Experimental.
	SetEmailVerificationMessage(val *string)
	// Experimental.
	EmailVerificationMessageInput() *string
	// Experimental.
	EmailVerificationSubject() *string
	// Experimental.
	SetEmailVerificationSubject(val *string)
	// Experimental.
	EmailVerificationSubjectInput() *string
	// Experimental.
	Endpoint() *string
	// Experimental.
	EstimatedNumberOfUsers() *float64
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
	LambdaConfig() TfUserPool_LambdaConfigPropertyOutputReference
	// Experimental.
	LambdaConfigInput() *TfUserPool_LambdaConfigProperty
	// Experimental.
	LastModifiedDate() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MfaConfiguration() *string
	// Experimental.
	SetMfaConfiguration(val *string)
	// Experimental.
	MfaConfigurationInput() *string
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PasswordPolicy() TfUserPool_PasswordPolicyPropertyOutputReference
	// Experimental.
	PasswordPolicyInput() *TfUserPool_PasswordPolicyProperty
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
	Schema() TfUserPool_SchemaPropertyList
	// Experimental.
	SchemaInput() interface{}
	// Experimental.
	SignInPolicy() TfUserPool_SignInPolicyPropertyOutputReference
	// Experimental.
	SignInPolicyInput() *TfUserPool_SignInPolicyProperty
	// Experimental.
	SmsAuthenticationMessage() *string
	// Experimental.
	SetSmsAuthenticationMessage(val *string)
	// Experimental.
	SmsAuthenticationMessageInput() *string
	// Experimental.
	SmsConfiguration() TfUserPool_SmsConfigurationPropertyOutputReference
	// Experimental.
	SmsConfigurationInput() *TfUserPool_SmsConfigurationProperty
	// Experimental.
	SmsVerificationMessage() *string
	// Experimental.
	SetSmsVerificationMessage(val *string)
	// Experimental.
	SmsVerificationMessageInput() *string
	// Experimental.
	SoftwareTokenMfaConfiguration() TfUserPool_SoftwareTokenMfaConfigurationPropertyOutputReference
	// Experimental.
	SoftwareTokenMfaConfigurationInput() *TfUserPool_SoftwareTokenMfaConfigurationProperty
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
	UserAttributeUpdateSettings() TfUserPool_UserAttributeUpdateSettingsPropertyOutputReference
	// Experimental.
	UserAttributeUpdateSettingsInput() *TfUserPool_UserAttributeUpdateSettingsProperty
	// Experimental.
	UsernameAttributes() *[]*string
	// Experimental.
	SetUsernameAttributes(val *[]*string)
	// Experimental.
	UsernameAttributesInput() *[]*string
	// Experimental.
	UsernameConfiguration() TfUserPool_UsernameConfigurationPropertyOutputReference
	// Experimental.
	UsernameConfigurationInput() *TfUserPool_UsernameConfigurationProperty
	// Experimental.
	UserPoolAddOns() TfUserPool_UserPoolAddOnsPropertyOutputReference
	// Experimental.
	UserPoolAddOnsInput() *TfUserPool_UserPoolAddOnsProperty
	// Experimental.
	UserPoolTier() *string
	// Experimental.
	SetUserPoolTier(val *string)
	// Experimental.
	UserPoolTierInput() *string
	// Experimental.
	VerificationMessageTemplate() TfUserPool_VerificationMessageTemplatePropertyOutputReference
	// Experimental.
	VerificationMessageTemplateInput() *TfUserPool_VerificationMessageTemplateProperty
	// Experimental.
	WebAuthnConfiguration() TfUserPool_WebAuthnConfigurationPropertyOutputReference
	// Experimental.
	WebAuthnConfigurationInput() *TfUserPool_WebAuthnConfigurationProperty
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
	PutAccountRecoverySetting(value *TfUserPool_AccountRecoverySettingProperty)
	// Experimental.
	PutAdminCreateUserConfig(value *TfUserPool_AdminCreateUserConfigProperty)
	// Experimental.
	PutDeviceConfiguration(value *TfUserPool_DeviceConfigurationProperty)
	// Experimental.
	PutEmailConfiguration(value *TfUserPool_EmailConfigurationProperty)
	// Experimental.
	PutEmailMfaConfiguration(value *TfUserPool_EmailMfaConfigurationProperty)
	// Experimental.
	PutLambdaConfig(value *TfUserPool_LambdaConfigProperty)
	// Experimental.
	PutPasswordPolicy(value *TfUserPool_PasswordPolicyProperty)
	// Experimental.
	PutSchema(value interface{})
	// Experimental.
	PutSignInPolicy(value *TfUserPool_SignInPolicyProperty)
	// Experimental.
	PutSmsConfiguration(value *TfUserPool_SmsConfigurationProperty)
	// Experimental.
	PutSoftwareTokenMfaConfiguration(value *TfUserPool_SoftwareTokenMfaConfigurationProperty)
	// Experimental.
	PutUserAttributeUpdateSettings(value *TfUserPool_UserAttributeUpdateSettingsProperty)
	// Experimental.
	PutUsernameConfiguration(value *TfUserPool_UsernameConfigurationProperty)
	// Experimental.
	PutUserPoolAddOns(value *TfUserPool_UserPoolAddOnsProperty)
	// Experimental.
	PutVerificationMessageTemplate(value *TfUserPool_VerificationMessageTemplateProperty)
	// Experimental.
	PutWebAuthnConfiguration(value *TfUserPool_WebAuthnConfigurationProperty)
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
	ResetAccountRecoverySetting()
	// Experimental.
	ResetAdminCreateUserConfig()
	// Experimental.
	ResetAliasAttributes()
	// Experimental.
	ResetAutoVerifiedAttributes()
	// Experimental.
	ResetDeletionProtection()
	// Experimental.
	ResetDeviceConfiguration()
	// Experimental.
	ResetEmailConfiguration()
	// Experimental.
	ResetEmailMfaConfiguration()
	// Experimental.
	ResetEmailVerificationMessage()
	// Experimental.
	ResetEmailVerificationSubject()
	// Experimental.
	ResetId()
	// Experimental.
	ResetLambdaConfig()
	// Experimental.
	ResetMfaConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPasswordPolicy()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSchema()
	// Experimental.
	ResetSignInPolicy()
	// Experimental.
	ResetSmsAuthenticationMessage()
	// Experimental.
	ResetSmsConfiguration()
	// Experimental.
	ResetSmsVerificationMessage()
	// Experimental.
	ResetSoftwareTokenMfaConfiguration()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetUserAttributeUpdateSettings()
	// Experimental.
	ResetUsernameAttributes()
	// Experimental.
	ResetUsernameConfiguration()
	// Experimental.
	ResetUserPoolAddOns()
	// Experimental.
	ResetUserPoolTier()
	// Experimental.
	ResetVerificationMessageTemplate()
	// Experimental.
	ResetWebAuthnConfiguration()
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

// The jsii proxy struct for TfUserPool
type jsiiProxy_TfUserPool struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfUserPool) AccountRecoverySetting() TfUserPool_AccountRecoverySettingPropertyOutputReference {
	var returns TfUserPool_AccountRecoverySettingPropertyOutputReference
	_jsii_.Get(
		j,
		"accountRecoverySetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) AccountRecoverySettingInput() *TfUserPool_AccountRecoverySettingProperty {
	var returns *TfUserPool_AccountRecoverySettingProperty
	_jsii_.Get(
		j,
		"accountRecoverySettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) AdminCreateUserConfig() TfUserPool_AdminCreateUserConfigPropertyOutputReference {
	var returns TfUserPool_AdminCreateUserConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"adminCreateUserConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) AdminCreateUserConfigInput() *TfUserPool_AdminCreateUserConfigProperty {
	var returns *TfUserPool_AdminCreateUserConfigProperty
	_jsii_.Get(
		j,
		"adminCreateUserConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) AliasAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) AliasAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) AutoVerifiedAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoVerifiedAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) AutoVerifiedAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoVerifiedAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) CreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) CustomDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) DeletionProtection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) DeletionProtectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) DeviceConfiguration() TfUserPool_DeviceConfigurationPropertyOutputReference {
	var returns TfUserPool_DeviceConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"deviceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) DeviceConfigurationInput() *TfUserPool_DeviceConfigurationProperty {
	var returns *TfUserPool_DeviceConfigurationProperty
	_jsii_.Get(
		j,
		"deviceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) EmailConfiguration() TfUserPool_EmailConfigurationPropertyOutputReference {
	var returns TfUserPool_EmailConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"emailConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) EmailConfigurationInput() *TfUserPool_EmailConfigurationProperty {
	var returns *TfUserPool_EmailConfigurationProperty
	_jsii_.Get(
		j,
		"emailConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) EmailMfaConfiguration() TfUserPool_EmailMfaConfigurationPropertyOutputReference {
	var returns TfUserPool_EmailMfaConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"emailMfaConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) EmailMfaConfigurationInput() *TfUserPool_EmailMfaConfigurationProperty {
	var returns *TfUserPool_EmailMfaConfigurationProperty
	_jsii_.Get(
		j,
		"emailMfaConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) EmailVerificationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) EmailVerificationMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) EmailVerificationSubject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationSubject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) EmailVerificationSubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationSubjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) EstimatedNumberOfUsers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"estimatedNumberOfUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) LambdaConfig() TfUserPool_LambdaConfigPropertyOutputReference {
	var returns TfUserPool_LambdaConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"lambdaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) LambdaConfigInput() *TfUserPool_LambdaConfigProperty {
	var returns *TfUserPool_LambdaConfigProperty
	_jsii_.Get(
		j,
		"lambdaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) LastModifiedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) MfaConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mfaConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) MfaConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mfaConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) PasswordPolicy() TfUserPool_PasswordPolicyPropertyOutputReference {
	var returns TfUserPool_PasswordPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"passwordPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) PasswordPolicyInput() *TfUserPool_PasswordPolicyProperty {
	var returns *TfUserPool_PasswordPolicyProperty
	_jsii_.Get(
		j,
		"passwordPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) Schema() TfUserPool_SchemaPropertyList {
	var returns TfUserPool_SchemaPropertyList
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) SchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) SignInPolicy() TfUserPool_SignInPolicyPropertyOutputReference {
	var returns TfUserPool_SignInPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"signInPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) SignInPolicyInput() *TfUserPool_SignInPolicyProperty {
	var returns *TfUserPool_SignInPolicyProperty
	_jsii_.Get(
		j,
		"signInPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) SmsAuthenticationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsAuthenticationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) SmsAuthenticationMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsAuthenticationMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) SmsConfiguration() TfUserPool_SmsConfigurationPropertyOutputReference {
	var returns TfUserPool_SmsConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"smsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) SmsConfigurationInput() *TfUserPool_SmsConfigurationProperty {
	var returns *TfUserPool_SmsConfigurationProperty
	_jsii_.Get(
		j,
		"smsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) SmsVerificationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsVerificationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) SmsVerificationMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsVerificationMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) SoftwareTokenMfaConfiguration() TfUserPool_SoftwareTokenMfaConfigurationPropertyOutputReference {
	var returns TfUserPool_SoftwareTokenMfaConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"softwareTokenMfaConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) SoftwareTokenMfaConfigurationInput() *TfUserPool_SoftwareTokenMfaConfigurationProperty {
	var returns *TfUserPool_SoftwareTokenMfaConfigurationProperty
	_jsii_.Get(
		j,
		"softwareTokenMfaConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) UserAttributeUpdateSettings() TfUserPool_UserAttributeUpdateSettingsPropertyOutputReference {
	var returns TfUserPool_UserAttributeUpdateSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"userAttributeUpdateSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) UserAttributeUpdateSettingsInput() *TfUserPool_UserAttributeUpdateSettingsProperty {
	var returns *TfUserPool_UserAttributeUpdateSettingsProperty
	_jsii_.Get(
		j,
		"userAttributeUpdateSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) UsernameAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usernameAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) UsernameAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usernameAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) UsernameConfiguration() TfUserPool_UsernameConfigurationPropertyOutputReference {
	var returns TfUserPool_UsernameConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"usernameConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) UsernameConfigurationInput() *TfUserPool_UsernameConfigurationProperty {
	var returns *TfUserPool_UsernameConfigurationProperty
	_jsii_.Get(
		j,
		"usernameConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) UserPoolAddOns() TfUserPool_UserPoolAddOnsPropertyOutputReference {
	var returns TfUserPool_UserPoolAddOnsPropertyOutputReference
	_jsii_.Get(
		j,
		"userPoolAddOns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) UserPoolAddOnsInput() *TfUserPool_UserPoolAddOnsProperty {
	var returns *TfUserPool_UserPoolAddOnsProperty
	_jsii_.Get(
		j,
		"userPoolAddOnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) UserPoolTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) UserPoolTierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) VerificationMessageTemplate() TfUserPool_VerificationMessageTemplatePropertyOutputReference {
	var returns TfUserPool_VerificationMessageTemplatePropertyOutputReference
	_jsii_.Get(
		j,
		"verificationMessageTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) VerificationMessageTemplateInput() *TfUserPool_VerificationMessageTemplateProperty {
	var returns *TfUserPool_VerificationMessageTemplateProperty
	_jsii_.Get(
		j,
		"verificationMessageTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) WebAuthnConfiguration() TfUserPool_WebAuthnConfigurationPropertyOutputReference {
	var returns TfUserPool_WebAuthnConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"webAuthnConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool) WebAuthnConfigurationInput() *TfUserPool_WebAuthnConfigurationProperty {
	var returns *TfUserPool_WebAuthnConfigurationProperty
	_jsii_.Get(
		j,
		"webAuthnConfigurationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool aws_cognito_user_pool} Resource.
// Experimental.
func NewTfUserPool(scope constructs.Construct, id *string, config *TfUserPoolConfig) TfUserPool {
	_init_.Initialize()

	if err := validateNewTfUserPoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfUserPool{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.TfUserPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool aws_cognito_user_pool} Resource.
// Experimental.
func NewTfUserPool_Override(t TfUserPool, scope constructs.Construct, id *string, config *TfUserPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.TfUserPool",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfUserPool)SetAliasAttributes(val *[]*string) {
	if err := j.validateSetAliasAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliasAttributes",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetAutoVerifiedAttributes(val *[]*string) {
	if err := j.validateSetAutoVerifiedAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoVerifiedAttributes",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetDeletionProtection(val *string) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetEmailVerificationMessage(val *string) {
	if err := j.validateSetEmailVerificationMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailVerificationMessage",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetEmailVerificationSubject(val *string) {
	if err := j.validateSetEmailVerificationSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailVerificationSubject",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetMfaConfiguration(val *string) {
	if err := j.validateSetMfaConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mfaConfiguration",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetSmsAuthenticationMessage(val *string) {
	if err := j.validateSetSmsAuthenticationMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smsAuthenticationMessage",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetSmsVerificationMessage(val *string) {
	if err := j.validateSetSmsVerificationMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smsVerificationMessage",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetUsernameAttributes(val *[]*string) {
	if err := j.validateSetUsernameAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameAttributes",
		val,
	)
}

func (j *jsiiProxy_TfUserPool)SetUserPoolTier(val *string) {
	if err := j.validateSetUserPoolTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPoolTier",
		val,
	)
}

// Generates CDKTN code for importing a TfUserPool resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfUserPool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfUserPool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-cognito-idp.TfUserPool",
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
func TfUserPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfUserPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cognito-idp.TfUserPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfUserPool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfUserPool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cognito-idp.TfUserPool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfUserPool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfUserPool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cognito-idp.TfUserPool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfUserPool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-cognito-idp.TfUserPool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfUserPool) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfUserPool) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfUserPool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfUserPool) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserPool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfUserPool) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfUserPool) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfUserPool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfUserPool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfUserPool) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfUserPool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfUserPool) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserPool) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfUserPool) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserPool) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfUserPool) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfUserPool) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfUserPool) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfUserPool) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfUserPool) PutAccountRecoverySetting(value *TfUserPool_AccountRecoverySettingProperty) {
	if err := t.validatePutAccountRecoverySettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAccountRecoverySetting",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool) PutAdminCreateUserConfig(value *TfUserPool_AdminCreateUserConfigProperty) {
	if err := t.validatePutAdminCreateUserConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAdminCreateUserConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool) PutDeviceConfiguration(value *TfUserPool_DeviceConfigurationProperty) {
	if err := t.validatePutDeviceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeviceConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool) PutEmailConfiguration(value *TfUserPool_EmailConfigurationProperty) {
	if err := t.validatePutEmailConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEmailConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool) PutEmailMfaConfiguration(value *TfUserPool_EmailMfaConfigurationProperty) {
	if err := t.validatePutEmailMfaConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEmailMfaConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool) PutLambdaConfig(value *TfUserPool_LambdaConfigProperty) {
	if err := t.validatePutLambdaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLambdaConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool) PutPasswordPolicy(value *TfUserPool_PasswordPolicyProperty) {
	if err := t.validatePutPasswordPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPasswordPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool) PutSchema(value interface{}) {
	if err := t.validatePutSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSchema",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool) PutSignInPolicy(value *TfUserPool_SignInPolicyProperty) {
	if err := t.validatePutSignInPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSignInPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool) PutSmsConfiguration(value *TfUserPool_SmsConfigurationProperty) {
	if err := t.validatePutSmsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSmsConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool) PutSoftwareTokenMfaConfiguration(value *TfUserPool_SoftwareTokenMfaConfigurationProperty) {
	if err := t.validatePutSoftwareTokenMfaConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSoftwareTokenMfaConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool) PutUserAttributeUpdateSettings(value *TfUserPool_UserAttributeUpdateSettingsProperty) {
	if err := t.validatePutUserAttributeUpdateSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUserAttributeUpdateSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool) PutUsernameConfiguration(value *TfUserPool_UsernameConfigurationProperty) {
	if err := t.validatePutUsernameConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUsernameConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool) PutUserPoolAddOns(value *TfUserPool_UserPoolAddOnsProperty) {
	if err := t.validatePutUserPoolAddOnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUserPoolAddOns",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool) PutVerificationMessageTemplate(value *TfUserPool_VerificationMessageTemplateProperty) {
	if err := t.validatePutVerificationMessageTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVerificationMessageTemplate",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool) PutWebAuthnConfiguration(value *TfUserPool_WebAuthnConfigurationProperty) {
	if err := t.validatePutWebAuthnConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWebAuthnConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfUserPool) ResetAccountRecoverySetting() {
	_jsii_.InvokeVoid(
		t,
		"resetAccountRecoverySetting",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetAdminCreateUserConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetAdminCreateUserConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetAliasAttributes() {
	_jsii_.InvokeVoid(
		t,
		"resetAliasAttributes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetAutoVerifiedAttributes() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoVerifiedAttributes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		t,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetDeviceConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetDeviceConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetEmailConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetEmailConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetEmailMfaConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetEmailMfaConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetEmailVerificationMessage() {
	_jsii_.InvokeVoid(
		t,
		"resetEmailVerificationMessage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetEmailVerificationSubject() {
	_jsii_.InvokeVoid(
		t,
		"resetEmailVerificationSubject",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetLambdaConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetMfaConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetMfaConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetPasswordPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetPasswordPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetSchema() {
	_jsii_.InvokeVoid(
		t,
		"resetSchema",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetSignInPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetSignInPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetSmsAuthenticationMessage() {
	_jsii_.InvokeVoid(
		t,
		"resetSmsAuthenticationMessage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetSmsConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSmsConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetSmsVerificationMessage() {
	_jsii_.InvokeVoid(
		t,
		"resetSmsVerificationMessage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetSoftwareTokenMfaConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSoftwareTokenMfaConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetUserAttributeUpdateSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetUserAttributeUpdateSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetUsernameAttributes() {
	_jsii_.InvokeVoid(
		t,
		"resetUsernameAttributes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetUsernameConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetUsernameConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetUserPoolAddOns() {
	_jsii_.InvokeVoid(
		t,
		"resetUserPoolAddOns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetUserPoolTier() {
	_jsii_.InvokeVoid(
		t,
		"resetUserPoolTier",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetVerificationMessageTemplate() {
	_jsii_.InvokeVoid(
		t,
		"resetVerificationMessageTemplate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) ResetWebAuthnConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetWebAuthnConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserPool) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserPool) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserPool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserPool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserPool) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

