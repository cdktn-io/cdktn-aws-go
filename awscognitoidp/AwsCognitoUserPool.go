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
type AwsCognitoUserPool interface {
	cdktn.TerraformResource
	// Experimental.
	AccountRecoverySetting() AwsCognitoUserPool_AccountRecoverySettingPropertyOutputReference
	// Experimental.
	AccountRecoverySettingInput() *AwsCognitoUserPool_AccountRecoverySettingProperty
	// Experimental.
	AdminCreateUserConfig() AwsCognitoUserPool_AdminCreateUserConfigPropertyOutputReference
	// Experimental.
	AdminCreateUserConfigInput() *AwsCognitoUserPool_AdminCreateUserConfigProperty
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
	DeviceConfiguration() AwsCognitoUserPool_DeviceConfigurationPropertyOutputReference
	// Experimental.
	DeviceConfigurationInput() *AwsCognitoUserPool_DeviceConfigurationProperty
	// Experimental.
	Domain() *string
	// Experimental.
	EmailConfiguration() AwsCognitoUserPool_EmailConfigurationPropertyOutputReference
	// Experimental.
	EmailConfigurationInput() *AwsCognitoUserPool_EmailConfigurationProperty
	// Experimental.
	EmailMfaConfiguration() AwsCognitoUserPool_EmailMfaConfigurationPropertyOutputReference
	// Experimental.
	EmailMfaConfigurationInput() *AwsCognitoUserPool_EmailMfaConfigurationProperty
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
	LambdaConfig() AwsCognitoUserPool_LambdaConfigPropertyOutputReference
	// Experimental.
	LambdaConfigInput() *AwsCognitoUserPool_LambdaConfigProperty
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
	PasswordPolicy() AwsCognitoUserPool_PasswordPolicyPropertyOutputReference
	// Experimental.
	PasswordPolicyInput() *AwsCognitoUserPool_PasswordPolicyProperty
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
	Schema() AwsCognitoUserPool_SchemaPropertyList
	// Experimental.
	SchemaInput() interface{}
	// Experimental.
	SignInPolicy() AwsCognitoUserPool_SignInPolicyPropertyOutputReference
	// Experimental.
	SignInPolicyInput() *AwsCognitoUserPool_SignInPolicyProperty
	// Experimental.
	SmsAuthenticationMessage() *string
	// Experimental.
	SetSmsAuthenticationMessage(val *string)
	// Experimental.
	SmsAuthenticationMessageInput() *string
	// Experimental.
	SmsConfiguration() AwsCognitoUserPool_SmsConfigurationPropertyOutputReference
	// Experimental.
	SmsConfigurationInput() *AwsCognitoUserPool_SmsConfigurationProperty
	// Experimental.
	SmsVerificationMessage() *string
	// Experimental.
	SetSmsVerificationMessage(val *string)
	// Experimental.
	SmsVerificationMessageInput() *string
	// Experimental.
	SoftwareTokenMfaConfiguration() AwsCognitoUserPool_SoftwareTokenMfaConfigurationPropertyOutputReference
	// Experimental.
	SoftwareTokenMfaConfigurationInput() *AwsCognitoUserPool_SoftwareTokenMfaConfigurationProperty
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
	UserAttributeUpdateSettings() AwsCognitoUserPool_UserAttributeUpdateSettingsPropertyOutputReference
	// Experimental.
	UserAttributeUpdateSettingsInput() *AwsCognitoUserPool_UserAttributeUpdateSettingsProperty
	// Experimental.
	UsernameAttributes() *[]*string
	// Experimental.
	SetUsernameAttributes(val *[]*string)
	// Experimental.
	UsernameAttributesInput() *[]*string
	// Experimental.
	UsernameConfiguration() AwsCognitoUserPool_UsernameConfigurationPropertyOutputReference
	// Experimental.
	UsernameConfigurationInput() *AwsCognitoUserPool_UsernameConfigurationProperty
	// Experimental.
	UserPoolAddOns() AwsCognitoUserPool_UserPoolAddOnsPropertyOutputReference
	// Experimental.
	UserPoolAddOnsInput() *AwsCognitoUserPool_UserPoolAddOnsProperty
	// Experimental.
	UserPoolTier() *string
	// Experimental.
	SetUserPoolTier(val *string)
	// Experimental.
	UserPoolTierInput() *string
	// Experimental.
	VerificationMessageTemplate() AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference
	// Experimental.
	VerificationMessageTemplateInput() *AwsCognitoUserPool_VerificationMessageTemplateProperty
	// Experimental.
	WebAuthnConfiguration() AwsCognitoUserPool_WebAuthnConfigurationPropertyOutputReference
	// Experimental.
	WebAuthnConfigurationInput() *AwsCognitoUserPool_WebAuthnConfigurationProperty
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
	PutAccountRecoverySetting(value *AwsCognitoUserPool_AccountRecoverySettingProperty)
	// Experimental.
	PutAdminCreateUserConfig(value *AwsCognitoUserPool_AdminCreateUserConfigProperty)
	// Experimental.
	PutDeviceConfiguration(value *AwsCognitoUserPool_DeviceConfigurationProperty)
	// Experimental.
	PutEmailConfiguration(value *AwsCognitoUserPool_EmailConfigurationProperty)
	// Experimental.
	PutEmailMfaConfiguration(value *AwsCognitoUserPool_EmailMfaConfigurationProperty)
	// Experimental.
	PutLambdaConfig(value *AwsCognitoUserPool_LambdaConfigProperty)
	// Experimental.
	PutPasswordPolicy(value *AwsCognitoUserPool_PasswordPolicyProperty)
	// Experimental.
	PutSchema(value interface{})
	// Experimental.
	PutSignInPolicy(value *AwsCognitoUserPool_SignInPolicyProperty)
	// Experimental.
	PutSmsConfiguration(value *AwsCognitoUserPool_SmsConfigurationProperty)
	// Experimental.
	PutSoftwareTokenMfaConfiguration(value *AwsCognitoUserPool_SoftwareTokenMfaConfigurationProperty)
	// Experimental.
	PutUserAttributeUpdateSettings(value *AwsCognitoUserPool_UserAttributeUpdateSettingsProperty)
	// Experimental.
	PutUsernameConfiguration(value *AwsCognitoUserPool_UsernameConfigurationProperty)
	// Experimental.
	PutUserPoolAddOns(value *AwsCognitoUserPool_UserPoolAddOnsProperty)
	// Experimental.
	PutVerificationMessageTemplate(value *AwsCognitoUserPool_VerificationMessageTemplateProperty)
	// Experimental.
	PutWebAuthnConfiguration(value *AwsCognitoUserPool_WebAuthnConfigurationProperty)
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

// The jsii proxy struct for AwsCognitoUserPool
type jsiiProxy_AwsCognitoUserPool struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsCognitoUserPool) AccountRecoverySetting() AwsCognitoUserPool_AccountRecoverySettingPropertyOutputReference {
	var returns AwsCognitoUserPool_AccountRecoverySettingPropertyOutputReference
	_jsii_.Get(
		j,
		"accountRecoverySetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) AccountRecoverySettingInput() *AwsCognitoUserPool_AccountRecoverySettingProperty {
	var returns *AwsCognitoUserPool_AccountRecoverySettingProperty
	_jsii_.Get(
		j,
		"accountRecoverySettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) AdminCreateUserConfig() AwsCognitoUserPool_AdminCreateUserConfigPropertyOutputReference {
	var returns AwsCognitoUserPool_AdminCreateUserConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"adminCreateUserConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) AdminCreateUserConfigInput() *AwsCognitoUserPool_AdminCreateUserConfigProperty {
	var returns *AwsCognitoUserPool_AdminCreateUserConfigProperty
	_jsii_.Get(
		j,
		"adminCreateUserConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) AliasAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) AliasAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) AutoVerifiedAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoVerifiedAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) AutoVerifiedAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoVerifiedAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) CreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) CustomDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) DeletionProtection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) DeletionProtectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) DeviceConfiguration() AwsCognitoUserPool_DeviceConfigurationPropertyOutputReference {
	var returns AwsCognitoUserPool_DeviceConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"deviceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) DeviceConfigurationInput() *AwsCognitoUserPool_DeviceConfigurationProperty {
	var returns *AwsCognitoUserPool_DeviceConfigurationProperty
	_jsii_.Get(
		j,
		"deviceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) EmailConfiguration() AwsCognitoUserPool_EmailConfigurationPropertyOutputReference {
	var returns AwsCognitoUserPool_EmailConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"emailConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) EmailConfigurationInput() *AwsCognitoUserPool_EmailConfigurationProperty {
	var returns *AwsCognitoUserPool_EmailConfigurationProperty
	_jsii_.Get(
		j,
		"emailConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) EmailMfaConfiguration() AwsCognitoUserPool_EmailMfaConfigurationPropertyOutputReference {
	var returns AwsCognitoUserPool_EmailMfaConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"emailMfaConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) EmailMfaConfigurationInput() *AwsCognitoUserPool_EmailMfaConfigurationProperty {
	var returns *AwsCognitoUserPool_EmailMfaConfigurationProperty
	_jsii_.Get(
		j,
		"emailMfaConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) EmailVerificationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) EmailVerificationMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) EmailVerificationSubject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationSubject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) EmailVerificationSubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationSubjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) EstimatedNumberOfUsers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"estimatedNumberOfUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) LambdaConfig() AwsCognitoUserPool_LambdaConfigPropertyOutputReference {
	var returns AwsCognitoUserPool_LambdaConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"lambdaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) LambdaConfigInput() *AwsCognitoUserPool_LambdaConfigProperty {
	var returns *AwsCognitoUserPool_LambdaConfigProperty
	_jsii_.Get(
		j,
		"lambdaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) LastModifiedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) MfaConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mfaConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) MfaConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mfaConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) PasswordPolicy() AwsCognitoUserPool_PasswordPolicyPropertyOutputReference {
	var returns AwsCognitoUserPool_PasswordPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"passwordPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) PasswordPolicyInput() *AwsCognitoUserPool_PasswordPolicyProperty {
	var returns *AwsCognitoUserPool_PasswordPolicyProperty
	_jsii_.Get(
		j,
		"passwordPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) Schema() AwsCognitoUserPool_SchemaPropertyList {
	var returns AwsCognitoUserPool_SchemaPropertyList
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) SchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) SignInPolicy() AwsCognitoUserPool_SignInPolicyPropertyOutputReference {
	var returns AwsCognitoUserPool_SignInPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"signInPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) SignInPolicyInput() *AwsCognitoUserPool_SignInPolicyProperty {
	var returns *AwsCognitoUserPool_SignInPolicyProperty
	_jsii_.Get(
		j,
		"signInPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) SmsAuthenticationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsAuthenticationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) SmsAuthenticationMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsAuthenticationMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) SmsConfiguration() AwsCognitoUserPool_SmsConfigurationPropertyOutputReference {
	var returns AwsCognitoUserPool_SmsConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"smsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) SmsConfigurationInput() *AwsCognitoUserPool_SmsConfigurationProperty {
	var returns *AwsCognitoUserPool_SmsConfigurationProperty
	_jsii_.Get(
		j,
		"smsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) SmsVerificationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsVerificationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) SmsVerificationMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsVerificationMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) SoftwareTokenMfaConfiguration() AwsCognitoUserPool_SoftwareTokenMfaConfigurationPropertyOutputReference {
	var returns AwsCognitoUserPool_SoftwareTokenMfaConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"softwareTokenMfaConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) SoftwareTokenMfaConfigurationInput() *AwsCognitoUserPool_SoftwareTokenMfaConfigurationProperty {
	var returns *AwsCognitoUserPool_SoftwareTokenMfaConfigurationProperty
	_jsii_.Get(
		j,
		"softwareTokenMfaConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) UserAttributeUpdateSettings() AwsCognitoUserPool_UserAttributeUpdateSettingsPropertyOutputReference {
	var returns AwsCognitoUserPool_UserAttributeUpdateSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"userAttributeUpdateSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) UserAttributeUpdateSettingsInput() *AwsCognitoUserPool_UserAttributeUpdateSettingsProperty {
	var returns *AwsCognitoUserPool_UserAttributeUpdateSettingsProperty
	_jsii_.Get(
		j,
		"userAttributeUpdateSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) UsernameAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usernameAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) UsernameAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usernameAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) UsernameConfiguration() AwsCognitoUserPool_UsernameConfigurationPropertyOutputReference {
	var returns AwsCognitoUserPool_UsernameConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"usernameConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) UsernameConfigurationInput() *AwsCognitoUserPool_UsernameConfigurationProperty {
	var returns *AwsCognitoUserPool_UsernameConfigurationProperty
	_jsii_.Get(
		j,
		"usernameConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) UserPoolAddOns() AwsCognitoUserPool_UserPoolAddOnsPropertyOutputReference {
	var returns AwsCognitoUserPool_UserPoolAddOnsPropertyOutputReference
	_jsii_.Get(
		j,
		"userPoolAddOns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) UserPoolAddOnsInput() *AwsCognitoUserPool_UserPoolAddOnsProperty {
	var returns *AwsCognitoUserPool_UserPoolAddOnsProperty
	_jsii_.Get(
		j,
		"userPoolAddOnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) UserPoolTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) UserPoolTierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) VerificationMessageTemplate() AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference {
	var returns AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference
	_jsii_.Get(
		j,
		"verificationMessageTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) VerificationMessageTemplateInput() *AwsCognitoUserPool_VerificationMessageTemplateProperty {
	var returns *AwsCognitoUserPool_VerificationMessageTemplateProperty
	_jsii_.Get(
		j,
		"verificationMessageTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) WebAuthnConfiguration() AwsCognitoUserPool_WebAuthnConfigurationPropertyOutputReference {
	var returns AwsCognitoUserPool_WebAuthnConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"webAuthnConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool) WebAuthnConfigurationInput() *AwsCognitoUserPool_WebAuthnConfigurationProperty {
	var returns *AwsCognitoUserPool_WebAuthnConfigurationProperty
	_jsii_.Get(
		j,
		"webAuthnConfigurationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool aws_cognito_user_pool} Resource.
// Experimental.
func NewAwsCognitoUserPool(scope constructs.Construct, id *string, config *AwsCognitoUserPoolConfig) AwsCognitoUserPool {
	_init_.Initialize()

	if err := validateNewAwsCognitoUserPoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCognitoUserPool{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsCognitoUserPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool aws_cognito_user_pool} Resource.
// Experimental.
func NewAwsCognitoUserPool_Override(a AwsCognitoUserPool, scope constructs.Construct, id *string, config *AwsCognitoUserPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsCognitoUserPool",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetAliasAttributes(val *[]*string) {
	if err := j.validateSetAliasAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliasAttributes",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetAutoVerifiedAttributes(val *[]*string) {
	if err := j.validateSetAutoVerifiedAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoVerifiedAttributes",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetDeletionProtection(val *string) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetEmailVerificationMessage(val *string) {
	if err := j.validateSetEmailVerificationMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailVerificationMessage",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetEmailVerificationSubject(val *string) {
	if err := j.validateSetEmailVerificationSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailVerificationSubject",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetMfaConfiguration(val *string) {
	if err := j.validateSetMfaConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mfaConfiguration",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetSmsAuthenticationMessage(val *string) {
	if err := j.validateSetSmsAuthenticationMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smsAuthenticationMessage",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetSmsVerificationMessage(val *string) {
	if err := j.validateSetSmsVerificationMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smsVerificationMessage",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetUsernameAttributes(val *[]*string) {
	if err := j.validateSetUsernameAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameAttributes",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool)SetUserPoolTier(val *string) {
	if err := j.validateSetUserPoolTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPoolTier",
		val,
	)
}

// Generates CDKTN code for importing a AwsCognitoUserPool resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsCognitoUserPool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsCognitoUserPool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-cognito-idp.AwsCognitoUserPool",
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
func AwsCognitoUserPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCognitoUserPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cognito-idp.AwsCognitoUserPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsCognitoUserPool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCognitoUserPool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cognito-idp.AwsCognitoUserPool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsCognitoUserPool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCognitoUserPool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cognito-idp.AwsCognitoUserPool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsCognitoUserPool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-cognito-idp.AwsCognitoUserPool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsCognitoUserPool) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCognitoUserPool) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCognitoUserPool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCognitoUserPool) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCognitoUserPool) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCognitoUserPool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCognitoUserPool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCognitoUserPool) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCognitoUserPool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCognitoUserPool) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoUserPool) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCognitoUserPool) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsCognitoUserPool) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) PutAccountRecoverySetting(value *AwsCognitoUserPool_AccountRecoverySettingProperty) {
	if err := a.validatePutAccountRecoverySettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccountRecoverySetting",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) PutAdminCreateUserConfig(value *AwsCognitoUserPool_AdminCreateUserConfigProperty) {
	if err := a.validatePutAdminCreateUserConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdminCreateUserConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) PutDeviceConfiguration(value *AwsCognitoUserPool_DeviceConfigurationProperty) {
	if err := a.validatePutDeviceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeviceConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) PutEmailConfiguration(value *AwsCognitoUserPool_EmailConfigurationProperty) {
	if err := a.validatePutEmailConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEmailConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) PutEmailMfaConfiguration(value *AwsCognitoUserPool_EmailMfaConfigurationProperty) {
	if err := a.validatePutEmailMfaConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEmailMfaConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) PutLambdaConfig(value *AwsCognitoUserPool_LambdaConfigProperty) {
	if err := a.validatePutLambdaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) PutPasswordPolicy(value *AwsCognitoUserPool_PasswordPolicyProperty) {
	if err := a.validatePutPasswordPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPasswordPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) PutSchema(value interface{}) {
	if err := a.validatePutSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSchema",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) PutSignInPolicy(value *AwsCognitoUserPool_SignInPolicyProperty) {
	if err := a.validatePutSignInPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSignInPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) PutSmsConfiguration(value *AwsCognitoUserPool_SmsConfigurationProperty) {
	if err := a.validatePutSmsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSmsConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) PutSoftwareTokenMfaConfiguration(value *AwsCognitoUserPool_SoftwareTokenMfaConfigurationProperty) {
	if err := a.validatePutSoftwareTokenMfaConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSoftwareTokenMfaConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) PutUserAttributeUpdateSettings(value *AwsCognitoUserPool_UserAttributeUpdateSettingsProperty) {
	if err := a.validatePutUserAttributeUpdateSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUserAttributeUpdateSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) PutUsernameConfiguration(value *AwsCognitoUserPool_UsernameConfigurationProperty) {
	if err := a.validatePutUsernameConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUsernameConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) PutUserPoolAddOns(value *AwsCognitoUserPool_UserPoolAddOnsProperty) {
	if err := a.validatePutUserPoolAddOnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUserPoolAddOns",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) PutVerificationMessageTemplate(value *AwsCognitoUserPool_VerificationMessageTemplateProperty) {
	if err := a.validatePutVerificationMessageTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVerificationMessageTemplate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) PutWebAuthnConfiguration(value *AwsCognitoUserPool_WebAuthnConfigurationProperty) {
	if err := a.validatePutWebAuthnConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWebAuthnConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetAccountRecoverySetting() {
	_jsii_.InvokeVoid(
		a,
		"resetAccountRecoverySetting",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetAdminCreateUserConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAdminCreateUserConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetAliasAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetAliasAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetAutoVerifiedAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoVerifiedAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetDeviceConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetDeviceConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetEmailConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetEmailMfaConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailMfaConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetEmailVerificationMessage() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailVerificationMessage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetEmailVerificationSubject() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailVerificationSubject",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetLambdaConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetMfaConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetMfaConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetPasswordPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetPasswordPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetSchema() {
	_jsii_.InvokeVoid(
		a,
		"resetSchema",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetSignInPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetSignInPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetSmsAuthenticationMessage() {
	_jsii_.InvokeVoid(
		a,
		"resetSmsAuthenticationMessage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetSmsConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSmsConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetSmsVerificationMessage() {
	_jsii_.InvokeVoid(
		a,
		"resetSmsVerificationMessage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetSoftwareTokenMfaConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSoftwareTokenMfaConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetUserAttributeUpdateSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetUserAttributeUpdateSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetUsernameAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetUsernameAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetUsernameConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetUsernameConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetUserPoolAddOns() {
	_jsii_.InvokeVoid(
		a,
		"resetUserPoolAddOns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetUserPoolTier() {
	_jsii_.InvokeVoid(
		a,
		"resetUserPoolTier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetVerificationMessageTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetVerificationMessageTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) ResetWebAuthnConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetWebAuthnConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoUserPool) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoUserPool) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoUserPool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoUserPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoUserPool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoUserPool) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

