package cognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cognitoidp/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/cognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool aws_cognito_user_pool}.
// Experimental.
type AwsUserPool interface {
	cdktn.TerraformResource
	// Experimental.
	AccountRecoverySetting() AwsUserPool_AccountRecoverySettingPropertyOutputReference
	// Experimental.
	AccountRecoverySettingInput() *AwsUserPool_AccountRecoverySettingProperty
	// Experimental.
	AdminCreateUserConfig() AwsUserPool_AdminCreateUserConfigPropertyOutputReference
	// Experimental.
	AdminCreateUserConfigInput() *AwsUserPool_AdminCreateUserConfigProperty
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
	DeviceConfiguration() AwsUserPool_DeviceConfigurationPropertyOutputReference
	// Experimental.
	DeviceConfigurationInput() *AwsUserPool_DeviceConfigurationProperty
	// Experimental.
	Domain() *string
	// Experimental.
	EmailConfiguration() AwsUserPool_EmailConfigurationPropertyOutputReference
	// Experimental.
	EmailConfigurationInput() *AwsUserPool_EmailConfigurationProperty
	// Experimental.
	EmailMfaConfiguration() AwsUserPool_EmailMfaConfigurationPropertyOutputReference
	// Experimental.
	EmailMfaConfigurationInput() *AwsUserPool_EmailMfaConfigurationProperty
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
	LambdaConfig() AwsUserPool_LambdaConfigPropertyOutputReference
	// Experimental.
	LambdaConfigInput() *AwsUserPool_LambdaConfigProperty
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
	PasswordPolicy() AwsUserPool_PasswordPolicyPropertyOutputReference
	// Experimental.
	PasswordPolicyInput() *AwsUserPool_PasswordPolicyProperty
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
	Schema() AwsUserPool_SchemaPropertyList
	// Experimental.
	SchemaInput() interface{}
	// Experimental.
	SignInPolicy() AwsUserPool_SignInPolicyPropertyOutputReference
	// Experimental.
	SignInPolicyInput() *AwsUserPool_SignInPolicyProperty
	// Experimental.
	SmsAuthenticationMessage() *string
	// Experimental.
	SetSmsAuthenticationMessage(val *string)
	// Experimental.
	SmsAuthenticationMessageInput() *string
	// Experimental.
	SmsConfiguration() AwsUserPool_SmsConfigurationPropertyOutputReference
	// Experimental.
	SmsConfigurationInput() *AwsUserPool_SmsConfigurationProperty
	// Experimental.
	SmsVerificationMessage() *string
	// Experimental.
	SetSmsVerificationMessage(val *string)
	// Experimental.
	SmsVerificationMessageInput() *string
	// Experimental.
	SoftwareTokenMfaConfiguration() AwsUserPool_SoftwareTokenMfaConfigurationPropertyOutputReference
	// Experimental.
	SoftwareTokenMfaConfigurationInput() *AwsUserPool_SoftwareTokenMfaConfigurationProperty
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
	UserAttributeUpdateSettings() AwsUserPool_UserAttributeUpdateSettingsPropertyOutputReference
	// Experimental.
	UserAttributeUpdateSettingsInput() *AwsUserPool_UserAttributeUpdateSettingsProperty
	// Experimental.
	UsernameAttributes() *[]*string
	// Experimental.
	SetUsernameAttributes(val *[]*string)
	// Experimental.
	UsernameAttributesInput() *[]*string
	// Experimental.
	UsernameConfiguration() AwsUserPool_UsernameConfigurationPropertyOutputReference
	// Experimental.
	UsernameConfigurationInput() *AwsUserPool_UsernameConfigurationProperty
	// Experimental.
	UserPoolAddOns() AwsUserPool_UserPoolAddOnsPropertyOutputReference
	// Experimental.
	UserPoolAddOnsInput() *AwsUserPool_UserPoolAddOnsProperty
	// Experimental.
	UserPoolTier() *string
	// Experimental.
	SetUserPoolTier(val *string)
	// Experimental.
	UserPoolTierInput() *string
	// Experimental.
	VerificationMessageTemplate() AwsUserPool_VerificationMessageTemplatePropertyOutputReference
	// Experimental.
	VerificationMessageTemplateInput() *AwsUserPool_VerificationMessageTemplateProperty
	// Experimental.
	WebAuthnConfiguration() AwsUserPool_WebAuthnConfigurationPropertyOutputReference
	// Experimental.
	WebAuthnConfigurationInput() *AwsUserPool_WebAuthnConfigurationProperty
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
	PutAccountRecoverySetting(value *AwsUserPool_AccountRecoverySettingProperty)
	// Experimental.
	PutAdminCreateUserConfig(value *AwsUserPool_AdminCreateUserConfigProperty)
	// Experimental.
	PutDeviceConfiguration(value *AwsUserPool_DeviceConfigurationProperty)
	// Experimental.
	PutEmailConfiguration(value *AwsUserPool_EmailConfigurationProperty)
	// Experimental.
	PutEmailMfaConfiguration(value *AwsUserPool_EmailMfaConfigurationProperty)
	// Experimental.
	PutLambdaConfig(value *AwsUserPool_LambdaConfigProperty)
	// Experimental.
	PutPasswordPolicy(value *AwsUserPool_PasswordPolicyProperty)
	// Experimental.
	PutSchema(value interface{})
	// Experimental.
	PutSignInPolicy(value *AwsUserPool_SignInPolicyProperty)
	// Experimental.
	PutSmsConfiguration(value *AwsUserPool_SmsConfigurationProperty)
	// Experimental.
	PutSoftwareTokenMfaConfiguration(value *AwsUserPool_SoftwareTokenMfaConfigurationProperty)
	// Experimental.
	PutUserAttributeUpdateSettings(value *AwsUserPool_UserAttributeUpdateSettingsProperty)
	// Experimental.
	PutUsernameConfiguration(value *AwsUserPool_UsernameConfigurationProperty)
	// Experimental.
	PutUserPoolAddOns(value *AwsUserPool_UserPoolAddOnsProperty)
	// Experimental.
	PutVerificationMessageTemplate(value *AwsUserPool_VerificationMessageTemplateProperty)
	// Experimental.
	PutWebAuthnConfiguration(value *AwsUserPool_WebAuthnConfigurationProperty)
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

// The jsii proxy struct for AwsUserPool
type jsiiProxy_AwsUserPool struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsUserPool) AccountRecoverySetting() AwsUserPool_AccountRecoverySettingPropertyOutputReference {
	var returns AwsUserPool_AccountRecoverySettingPropertyOutputReference
	_jsii_.Get(
		j,
		"accountRecoverySetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) AccountRecoverySettingInput() *AwsUserPool_AccountRecoverySettingProperty {
	var returns *AwsUserPool_AccountRecoverySettingProperty
	_jsii_.Get(
		j,
		"accountRecoverySettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) AdminCreateUserConfig() AwsUserPool_AdminCreateUserConfigPropertyOutputReference {
	var returns AwsUserPool_AdminCreateUserConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"adminCreateUserConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) AdminCreateUserConfigInput() *AwsUserPool_AdminCreateUserConfigProperty {
	var returns *AwsUserPool_AdminCreateUserConfigProperty
	_jsii_.Get(
		j,
		"adminCreateUserConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) AliasAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) AliasAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) AutoVerifiedAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoVerifiedAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) AutoVerifiedAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoVerifiedAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) CreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) CustomDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) DeletionProtection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) DeletionProtectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) DeviceConfiguration() AwsUserPool_DeviceConfigurationPropertyOutputReference {
	var returns AwsUserPool_DeviceConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"deviceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) DeviceConfigurationInput() *AwsUserPool_DeviceConfigurationProperty {
	var returns *AwsUserPool_DeviceConfigurationProperty
	_jsii_.Get(
		j,
		"deviceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) EmailConfiguration() AwsUserPool_EmailConfigurationPropertyOutputReference {
	var returns AwsUserPool_EmailConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"emailConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) EmailConfigurationInput() *AwsUserPool_EmailConfigurationProperty {
	var returns *AwsUserPool_EmailConfigurationProperty
	_jsii_.Get(
		j,
		"emailConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) EmailMfaConfiguration() AwsUserPool_EmailMfaConfigurationPropertyOutputReference {
	var returns AwsUserPool_EmailMfaConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"emailMfaConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) EmailMfaConfigurationInput() *AwsUserPool_EmailMfaConfigurationProperty {
	var returns *AwsUserPool_EmailMfaConfigurationProperty
	_jsii_.Get(
		j,
		"emailMfaConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) EmailVerificationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) EmailVerificationMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) EmailVerificationSubject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationSubject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) EmailVerificationSubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationSubjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) EstimatedNumberOfUsers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"estimatedNumberOfUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) LambdaConfig() AwsUserPool_LambdaConfigPropertyOutputReference {
	var returns AwsUserPool_LambdaConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"lambdaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) LambdaConfigInput() *AwsUserPool_LambdaConfigProperty {
	var returns *AwsUserPool_LambdaConfigProperty
	_jsii_.Get(
		j,
		"lambdaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) LastModifiedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) MfaConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mfaConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) MfaConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mfaConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) PasswordPolicy() AwsUserPool_PasswordPolicyPropertyOutputReference {
	var returns AwsUserPool_PasswordPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"passwordPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) PasswordPolicyInput() *AwsUserPool_PasswordPolicyProperty {
	var returns *AwsUserPool_PasswordPolicyProperty
	_jsii_.Get(
		j,
		"passwordPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) Schema() AwsUserPool_SchemaPropertyList {
	var returns AwsUserPool_SchemaPropertyList
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) SchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) SignInPolicy() AwsUserPool_SignInPolicyPropertyOutputReference {
	var returns AwsUserPool_SignInPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"signInPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) SignInPolicyInput() *AwsUserPool_SignInPolicyProperty {
	var returns *AwsUserPool_SignInPolicyProperty
	_jsii_.Get(
		j,
		"signInPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) SmsAuthenticationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsAuthenticationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) SmsAuthenticationMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsAuthenticationMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) SmsConfiguration() AwsUserPool_SmsConfigurationPropertyOutputReference {
	var returns AwsUserPool_SmsConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"smsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) SmsConfigurationInput() *AwsUserPool_SmsConfigurationProperty {
	var returns *AwsUserPool_SmsConfigurationProperty
	_jsii_.Get(
		j,
		"smsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) SmsVerificationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsVerificationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) SmsVerificationMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsVerificationMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) SoftwareTokenMfaConfiguration() AwsUserPool_SoftwareTokenMfaConfigurationPropertyOutputReference {
	var returns AwsUserPool_SoftwareTokenMfaConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"softwareTokenMfaConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) SoftwareTokenMfaConfigurationInput() *AwsUserPool_SoftwareTokenMfaConfigurationProperty {
	var returns *AwsUserPool_SoftwareTokenMfaConfigurationProperty
	_jsii_.Get(
		j,
		"softwareTokenMfaConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) UserAttributeUpdateSettings() AwsUserPool_UserAttributeUpdateSettingsPropertyOutputReference {
	var returns AwsUserPool_UserAttributeUpdateSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"userAttributeUpdateSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) UserAttributeUpdateSettingsInput() *AwsUserPool_UserAttributeUpdateSettingsProperty {
	var returns *AwsUserPool_UserAttributeUpdateSettingsProperty
	_jsii_.Get(
		j,
		"userAttributeUpdateSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) UsernameAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usernameAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) UsernameAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usernameAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) UsernameConfiguration() AwsUserPool_UsernameConfigurationPropertyOutputReference {
	var returns AwsUserPool_UsernameConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"usernameConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) UsernameConfigurationInput() *AwsUserPool_UsernameConfigurationProperty {
	var returns *AwsUserPool_UsernameConfigurationProperty
	_jsii_.Get(
		j,
		"usernameConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) UserPoolAddOns() AwsUserPool_UserPoolAddOnsPropertyOutputReference {
	var returns AwsUserPool_UserPoolAddOnsPropertyOutputReference
	_jsii_.Get(
		j,
		"userPoolAddOns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) UserPoolAddOnsInput() *AwsUserPool_UserPoolAddOnsProperty {
	var returns *AwsUserPool_UserPoolAddOnsProperty
	_jsii_.Get(
		j,
		"userPoolAddOnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) UserPoolTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) UserPoolTierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) VerificationMessageTemplate() AwsUserPool_VerificationMessageTemplatePropertyOutputReference {
	var returns AwsUserPool_VerificationMessageTemplatePropertyOutputReference
	_jsii_.Get(
		j,
		"verificationMessageTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) VerificationMessageTemplateInput() *AwsUserPool_VerificationMessageTemplateProperty {
	var returns *AwsUserPool_VerificationMessageTemplateProperty
	_jsii_.Get(
		j,
		"verificationMessageTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) WebAuthnConfiguration() AwsUserPool_WebAuthnConfigurationPropertyOutputReference {
	var returns AwsUserPool_WebAuthnConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"webAuthnConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool) WebAuthnConfigurationInput() *AwsUserPool_WebAuthnConfigurationProperty {
	var returns *AwsUserPool_WebAuthnConfigurationProperty
	_jsii_.Get(
		j,
		"webAuthnConfigurationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool aws_cognito_user_pool} Resource.
// Experimental.
func NewAwsUserPool(scope constructs.Construct, id *string, config *AwsUserPoolConfig) AwsUserPool {
	_init_.Initialize()

	if err := validateNewAwsUserPoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsUserPool{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsUserPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool aws_cognito_user_pool} Resource.
// Experimental.
func NewAwsUserPool_Override(a AwsUserPool, scope constructs.Construct, id *string, config *AwsUserPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsUserPool",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsUserPool)SetAliasAttributes(val *[]*string) {
	if err := j.validateSetAliasAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliasAttributes",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetAutoVerifiedAttributes(val *[]*string) {
	if err := j.validateSetAutoVerifiedAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoVerifiedAttributes",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetDeletionProtection(val *string) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetEmailVerificationMessage(val *string) {
	if err := j.validateSetEmailVerificationMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailVerificationMessage",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetEmailVerificationSubject(val *string) {
	if err := j.validateSetEmailVerificationSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailVerificationSubject",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetMfaConfiguration(val *string) {
	if err := j.validateSetMfaConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mfaConfiguration",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetSmsAuthenticationMessage(val *string) {
	if err := j.validateSetSmsAuthenticationMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smsAuthenticationMessage",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetSmsVerificationMessage(val *string) {
	if err := j.validateSetSmsVerificationMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smsVerificationMessage",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetUsernameAttributes(val *[]*string) {
	if err := j.validateSetUsernameAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameAttributes",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool)SetUserPoolTier(val *string) {
	if err := j.validateSetUserPoolTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPoolTier",
		val,
	)
}

// Generates CDKTN code for importing a AwsUserPool resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsUserPool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsUserPool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-cognito-idp.AwsUserPool",
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
func AwsUserPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsUserPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cognito-idp.AwsUserPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsUserPool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsUserPool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cognito-idp.AwsUserPool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsUserPool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsUserPool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cognito-idp.AwsUserPool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsUserPool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-cognito-idp.AwsUserPool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsUserPool) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsUserPool) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsUserPool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsUserPool) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsUserPool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsUserPool) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsUserPool) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsUserPool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsUserPool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsUserPool) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsUserPool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsUserPool) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUserPool) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsUserPool) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsUserPool) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsUserPool) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsUserPool) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsUserPool) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsUserPool) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsUserPool) PutAccountRecoverySetting(value *AwsUserPool_AccountRecoverySettingProperty) {
	if err := a.validatePutAccountRecoverySettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccountRecoverySetting",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserPool) PutAdminCreateUserConfig(value *AwsUserPool_AdminCreateUserConfigProperty) {
	if err := a.validatePutAdminCreateUserConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdminCreateUserConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserPool) PutDeviceConfiguration(value *AwsUserPool_DeviceConfigurationProperty) {
	if err := a.validatePutDeviceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeviceConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserPool) PutEmailConfiguration(value *AwsUserPool_EmailConfigurationProperty) {
	if err := a.validatePutEmailConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEmailConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserPool) PutEmailMfaConfiguration(value *AwsUserPool_EmailMfaConfigurationProperty) {
	if err := a.validatePutEmailMfaConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEmailMfaConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserPool) PutLambdaConfig(value *AwsUserPool_LambdaConfigProperty) {
	if err := a.validatePutLambdaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserPool) PutPasswordPolicy(value *AwsUserPool_PasswordPolicyProperty) {
	if err := a.validatePutPasswordPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPasswordPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserPool) PutSchema(value interface{}) {
	if err := a.validatePutSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSchema",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserPool) PutSignInPolicy(value *AwsUserPool_SignInPolicyProperty) {
	if err := a.validatePutSignInPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSignInPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserPool) PutSmsConfiguration(value *AwsUserPool_SmsConfigurationProperty) {
	if err := a.validatePutSmsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSmsConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserPool) PutSoftwareTokenMfaConfiguration(value *AwsUserPool_SoftwareTokenMfaConfigurationProperty) {
	if err := a.validatePutSoftwareTokenMfaConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSoftwareTokenMfaConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserPool) PutUserAttributeUpdateSettings(value *AwsUserPool_UserAttributeUpdateSettingsProperty) {
	if err := a.validatePutUserAttributeUpdateSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUserAttributeUpdateSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserPool) PutUsernameConfiguration(value *AwsUserPool_UsernameConfigurationProperty) {
	if err := a.validatePutUsernameConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUsernameConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserPool) PutUserPoolAddOns(value *AwsUserPool_UserPoolAddOnsProperty) {
	if err := a.validatePutUserPoolAddOnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUserPoolAddOns",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserPool) PutVerificationMessageTemplate(value *AwsUserPool_VerificationMessageTemplateProperty) {
	if err := a.validatePutVerificationMessageTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVerificationMessageTemplate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserPool) PutWebAuthnConfiguration(value *AwsUserPool_WebAuthnConfigurationProperty) {
	if err := a.validatePutWebAuthnConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWebAuthnConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserPool) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsUserPool) ResetAccountRecoverySetting() {
	_jsii_.InvokeVoid(
		a,
		"resetAccountRecoverySetting",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetAdminCreateUserConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAdminCreateUserConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetAliasAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetAliasAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetAutoVerifiedAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoVerifiedAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetDeviceConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetDeviceConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetEmailConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetEmailMfaConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailMfaConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetEmailVerificationMessage() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailVerificationMessage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetEmailVerificationSubject() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailVerificationSubject",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetLambdaConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetMfaConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetMfaConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetPasswordPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetPasswordPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetSchema() {
	_jsii_.InvokeVoid(
		a,
		"resetSchema",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetSignInPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetSignInPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetSmsAuthenticationMessage() {
	_jsii_.InvokeVoid(
		a,
		"resetSmsAuthenticationMessage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetSmsConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSmsConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetSmsVerificationMessage() {
	_jsii_.InvokeVoid(
		a,
		"resetSmsVerificationMessage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetSoftwareTokenMfaConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSoftwareTokenMfaConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetUserAttributeUpdateSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetUserAttributeUpdateSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetUsernameAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetUsernameAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetUsernameConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetUsernameConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetUserPoolAddOns() {
	_jsii_.InvokeVoid(
		a,
		"resetUserPoolAddOns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetUserPoolTier() {
	_jsii_.InvokeVoid(
		a,
		"resetUserPoolTier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetVerificationMessageTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetVerificationMessageTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) ResetWebAuthnConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetWebAuthnConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUserPool) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUserPool) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUserPool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUserPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUserPool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUserPool) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

