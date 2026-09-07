package workspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/workspaces/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/workspaces/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory aws_workspaces_directory}.
// Experimental.
type AwsDirectory interface {
	cdktn.TerraformResource
	// Experimental.
	ActiveDirectoryConfig() AwsDirectory_ActiveDirectoryConfigPropertyOutputReference
	// Experimental.
	ActiveDirectoryConfigInput() *AwsDirectory_ActiveDirectoryConfigProperty
	// Experimental.
	Alias() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CertificateBasedAuthProperties() AwsDirectory_CertificateBasedAuthPropertiesPropertyOutputReference
	// Experimental.
	CertificateBasedAuthPropertiesInput() *AwsDirectory_CertificateBasedAuthPropertiesProperty
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
	CustomerUserName() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DirectoryId() *string
	// Experimental.
	SetDirectoryId(val *string)
	// Experimental.
	DirectoryIdInput() *string
	// Experimental.
	DirectoryName() *string
	// Experimental.
	DirectoryType() *string
	// Experimental.
	DnsIpAddresses() *[]*string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	IamRoleId() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	IpGroupIds() *[]*string
	// Experimental.
	SetIpGroupIds(val *[]*string)
	// Experimental.
	IpGroupIdsInput() *[]*string
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
	RegistrationCode() *string
	// Experimental.
	SamlProperties() AwsDirectory_SamlPropertiesPropertyOutputReference
	// Experimental.
	SamlPropertiesInput() *AwsDirectory_SamlPropertiesProperty
	// Experimental.
	SelfServicePermissions() AwsDirectory_SelfServicePermissionsPropertyOutputReference
	// Experimental.
	SelfServicePermissionsInput() *AwsDirectory_SelfServicePermissionsProperty
	// Experimental.
	SubnetIds() *[]*string
	// Experimental.
	SetSubnetIds(val *[]*string)
	// Experimental.
	SubnetIdsInput() *[]*string
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
	Tenancy() *string
	// Experimental.
	SetTenancy(val *string)
	// Experimental.
	TenancyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	UserIdentityType() *string
	// Experimental.
	SetUserIdentityType(val *string)
	// Experimental.
	UserIdentityTypeInput() *string
	// Experimental.
	WorkspaceAccessProperties() AwsDirectory_WorkspaceAccessPropertiesPropertyOutputReference
	// Experimental.
	WorkspaceAccessPropertiesInput() *AwsDirectory_WorkspaceAccessPropertiesProperty
	// Experimental.
	WorkspaceCreationProperties() AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference
	// Experimental.
	WorkspaceCreationPropertiesInput() *AwsDirectory_WorkspaceCreationPropertiesProperty
	// Experimental.
	WorkspaceDirectoryDescription() *string
	// Experimental.
	SetWorkspaceDirectoryDescription(val *string)
	// Experimental.
	WorkspaceDirectoryDescriptionInput() *string
	// Experimental.
	WorkspaceDirectoryName() *string
	// Experimental.
	SetWorkspaceDirectoryName(val *string)
	// Experimental.
	WorkspaceDirectoryNameInput() *string
	// Experimental.
	WorkspaceSecurityGroupId() *string
	// Experimental.
	WorkspaceType() *string
	// Experimental.
	SetWorkspaceType(val *string)
	// Experimental.
	WorkspaceTypeInput() *string
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
	PutActiveDirectoryConfig(value *AwsDirectory_ActiveDirectoryConfigProperty)
	// Experimental.
	PutCertificateBasedAuthProperties(value *AwsDirectory_CertificateBasedAuthPropertiesProperty)
	// Experimental.
	PutSamlProperties(value *AwsDirectory_SamlPropertiesProperty)
	// Experimental.
	PutSelfServicePermissions(value *AwsDirectory_SelfServicePermissionsProperty)
	// Experimental.
	PutWorkspaceAccessProperties(value *AwsDirectory_WorkspaceAccessPropertiesProperty)
	// Experimental.
	PutWorkspaceCreationProperties(value *AwsDirectory_WorkspaceCreationPropertiesProperty)
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
	ResetActiveDirectoryConfig()
	// Experimental.
	ResetCertificateBasedAuthProperties()
	// Experimental.
	ResetDirectoryId()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIpGroupIds()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSamlProperties()
	// Experimental.
	ResetSelfServicePermissions()
	// Experimental.
	ResetSubnetIds()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTenancy()
	// Experimental.
	ResetUserIdentityType()
	// Experimental.
	ResetWorkspaceAccessProperties()
	// Experimental.
	ResetWorkspaceCreationProperties()
	// Experimental.
	ResetWorkspaceDirectoryDescription()
	// Experimental.
	ResetWorkspaceDirectoryName()
	// Experimental.
	ResetWorkspaceType()
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

// The jsii proxy struct for AwsDirectory
type jsiiProxy_AwsDirectory struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsDirectory) ActiveDirectoryConfig() AwsDirectory_ActiveDirectoryConfigPropertyOutputReference {
	var returns AwsDirectory_ActiveDirectoryConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"activeDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) ActiveDirectoryConfigInput() *AwsDirectory_ActiveDirectoryConfigProperty {
	var returns *AwsDirectory_ActiveDirectoryConfigProperty
	_jsii_.Get(
		j,
		"activeDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) CertificateBasedAuthProperties() AwsDirectory_CertificateBasedAuthPropertiesPropertyOutputReference {
	var returns AwsDirectory_CertificateBasedAuthPropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"certificateBasedAuthProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) CertificateBasedAuthPropertiesInput() *AwsDirectory_CertificateBasedAuthPropertiesProperty {
	var returns *AwsDirectory_CertificateBasedAuthPropertiesProperty
	_jsii_.Get(
		j,
		"certificateBasedAuthPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) CustomerUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) DirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) DirectoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) DirectoryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) DnsIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) IamRoleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) IpGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) IpGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) RegistrationCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registrationCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) SamlProperties() AwsDirectory_SamlPropertiesPropertyOutputReference {
	var returns AwsDirectory_SamlPropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"samlProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) SamlPropertiesInput() *AwsDirectory_SamlPropertiesProperty {
	var returns *AwsDirectory_SamlPropertiesProperty
	_jsii_.Get(
		j,
		"samlPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) SelfServicePermissions() AwsDirectory_SelfServicePermissionsPropertyOutputReference {
	var returns AwsDirectory_SelfServicePermissionsPropertyOutputReference
	_jsii_.Get(
		j,
		"selfServicePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) SelfServicePermissionsInput() *AwsDirectory_SelfServicePermissionsProperty {
	var returns *AwsDirectory_SelfServicePermissionsProperty
	_jsii_.Get(
		j,
		"selfServicePermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) Tenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) TenancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) UserIdentityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdentityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) UserIdentityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdentityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) WorkspaceAccessProperties() AwsDirectory_WorkspaceAccessPropertiesPropertyOutputReference {
	var returns AwsDirectory_WorkspaceAccessPropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"workspaceAccessProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) WorkspaceAccessPropertiesInput() *AwsDirectory_WorkspaceAccessPropertiesProperty {
	var returns *AwsDirectory_WorkspaceAccessPropertiesProperty
	_jsii_.Get(
		j,
		"workspaceAccessPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) WorkspaceCreationProperties() AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference {
	var returns AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"workspaceCreationProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) WorkspaceCreationPropertiesInput() *AwsDirectory_WorkspaceCreationPropertiesProperty {
	var returns *AwsDirectory_WorkspaceCreationPropertiesProperty
	_jsii_.Get(
		j,
		"workspaceCreationPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) WorkspaceDirectoryDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceDirectoryDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) WorkspaceDirectoryDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceDirectoryDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) WorkspaceDirectoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceDirectoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) WorkspaceDirectoryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceDirectoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) WorkspaceSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) WorkspaceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory) WorkspaceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory aws_workspaces_directory} Resource.
// Experimental.
func NewAwsDirectory(scope constructs.Construct, id *string, config *AwsDirectoryConfig) AwsDirectory {
	_init_.Initialize()

	if err := validateNewAwsDirectoryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDirectory{}

	_jsii_.Create(
		"@cdktn/aws-workspaces.AwsDirectory",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory aws_workspaces_directory} Resource.
// Experimental.
func NewAwsDirectory_Override(a AwsDirectory, scope constructs.Construct, id *string, config *AwsDirectoryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-workspaces.AwsDirectory",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsDirectory)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory)SetDirectoryId(val *string) {
	if err := j.validateSetDirectoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryId",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory)SetIpGroupIds(val *[]*string) {
	if err := j.validateSetIpGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipGroupIds",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory)SetTenancy(val *string) {
	if err := j.validateSetTenancyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenancy",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory)SetUserIdentityType(val *string) {
	if err := j.validateSetUserIdentityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userIdentityType",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory)SetWorkspaceDirectoryDescription(val *string) {
	if err := j.validateSetWorkspaceDirectoryDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceDirectoryDescription",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory)SetWorkspaceDirectoryName(val *string) {
	if err := j.validateSetWorkspaceDirectoryNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceDirectoryName",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory)SetWorkspaceType(val *string) {
	if err := j.validateSetWorkspaceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceType",
		val,
	)
}

// Generates CDKTN code for importing a AwsDirectory resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsDirectory_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsDirectory_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-workspaces.AwsDirectory",
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
func AwsDirectory_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDirectory_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-workspaces.AwsDirectory",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsDirectory_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDirectory_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-workspaces.AwsDirectory",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsDirectory_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDirectory_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-workspaces.AwsDirectory",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsDirectory_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-workspaces.AwsDirectory",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsDirectory) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsDirectory) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsDirectory) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDirectory) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDirectory) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDirectory) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDirectory) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDirectory) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDirectory) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDirectory) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDirectory) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDirectory) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDirectory) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsDirectory) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDirectory) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsDirectory) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsDirectory) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsDirectory) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsDirectory) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsDirectory) PutActiveDirectoryConfig(value *AwsDirectory_ActiveDirectoryConfigProperty) {
	if err := a.validatePutActiveDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putActiveDirectoryConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDirectory) PutCertificateBasedAuthProperties(value *AwsDirectory_CertificateBasedAuthPropertiesProperty) {
	if err := a.validatePutCertificateBasedAuthPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCertificateBasedAuthProperties",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDirectory) PutSamlProperties(value *AwsDirectory_SamlPropertiesProperty) {
	if err := a.validatePutSamlPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSamlProperties",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDirectory) PutSelfServicePermissions(value *AwsDirectory_SelfServicePermissionsProperty) {
	if err := a.validatePutSelfServicePermissionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSelfServicePermissions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDirectory) PutWorkspaceAccessProperties(value *AwsDirectory_WorkspaceAccessPropertiesProperty) {
	if err := a.validatePutWorkspaceAccessPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWorkspaceAccessProperties",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDirectory) PutWorkspaceCreationProperties(value *AwsDirectory_WorkspaceCreationPropertiesProperty) {
	if err := a.validatePutWorkspaceCreationPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWorkspaceCreationProperties",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDirectory) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsDirectory) ResetActiveDirectoryConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetActiveDirectoryConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory) ResetCertificateBasedAuthProperties() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateBasedAuthProperties",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory) ResetDirectoryId() {
	_jsii_.InvokeVoid(
		a,
		"resetDirectoryId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory) ResetIpGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetIpGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory) ResetSamlProperties() {
	_jsii_.InvokeVoid(
		a,
		"resetSamlProperties",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory) ResetSelfServicePermissions() {
	_jsii_.InvokeVoid(
		a,
		"resetSelfServicePermissions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory) ResetTenancy() {
	_jsii_.InvokeVoid(
		a,
		"resetTenancy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory) ResetUserIdentityType() {
	_jsii_.InvokeVoid(
		a,
		"resetUserIdentityType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory) ResetWorkspaceAccessProperties() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkspaceAccessProperties",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory) ResetWorkspaceCreationProperties() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkspaceCreationProperties",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory) ResetWorkspaceDirectoryDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkspaceDirectoryDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory) ResetWorkspaceDirectoryName() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkspaceDirectoryName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory) ResetWorkspaceType() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkspaceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDirectory) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDirectory) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDirectory) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDirectory) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDirectory) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDirectory) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

