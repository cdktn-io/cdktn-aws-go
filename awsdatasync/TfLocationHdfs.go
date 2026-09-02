package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdatasync/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsdatasync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs aws_datasync_location_hdfs}.
// Experimental.
type TfLocationHdfs interface {
	cdktn.TerraformResource
	// Experimental.
	AgentArns() *[]*string
	// Experimental.
	SetAgentArns(val *[]*string)
	// Experimental.
	AgentArnsInput() *[]*string
	// Experimental.
	Arn() *string
	// Experimental.
	AuthenticationType() *string
	// Experimental.
	SetAuthenticationType(val *string)
	// Experimental.
	AuthenticationTypeInput() *string
	// Experimental.
	BlockSize() *float64
	// Experimental.
	SetBlockSize(val *float64)
	// Experimental.
	BlockSizeInput() *float64
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
	KerberosKeytab() *string
	// Experimental.
	SetKerberosKeytab(val *string)
	// Experimental.
	KerberosKeytabBase64() *string
	// Experimental.
	SetKerberosKeytabBase64(val *string)
	// Experimental.
	KerberosKeytabBase64Input() *string
	// Experimental.
	KerberosKeytabInput() *string
	// Experimental.
	KerberosKrb5Conf() *string
	// Experimental.
	SetKerberosKrb5Conf(val *string)
	// Experimental.
	KerberosKrb5ConfBase64() *string
	// Experimental.
	SetKerberosKrb5ConfBase64(val *string)
	// Experimental.
	KerberosKrb5ConfBase64Input() *string
	// Experimental.
	KerberosKrb5ConfInput() *string
	// Experimental.
	KerberosPrincipal() *string
	// Experimental.
	SetKerberosPrincipal(val *string)
	// Experimental.
	KerberosPrincipalInput() *string
	// Experimental.
	KmsKeyProviderUri() *string
	// Experimental.
	SetKmsKeyProviderUri(val *string)
	// Experimental.
	KmsKeyProviderUriInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	NameNode() TfLocationHdfs_NameNodePropertyList
	// Experimental.
	NameNodeInput() interface{}
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
	QopConfiguration() TfLocationHdfs_QopConfigurationPropertyOutputReference
	// Experimental.
	QopConfigurationInput() *TfLocationHdfs_QopConfigurationProperty
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	ReplicationFactor() *float64
	// Experimental.
	SetReplicationFactor(val *float64)
	// Experimental.
	ReplicationFactorInput() *float64
	// Experimental.
	SimpleUser() *string
	// Experimental.
	SetSimpleUser(val *string)
	// Experimental.
	SimpleUserInput() *string
	// Experimental.
	Subdirectory() *string
	// Experimental.
	SetSubdirectory(val *string)
	// Experimental.
	SubdirectoryInput() *string
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
	Uri() *string
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
	PutNameNode(value interface{})
	// Experimental.
	PutQopConfiguration(value *TfLocationHdfs_QopConfigurationProperty)
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
	ResetAuthenticationType()
	// Experimental.
	ResetBlockSize()
	// Experimental.
	ResetId()
	// Experimental.
	ResetKerberosKeytab()
	// Experimental.
	ResetKerberosKeytabBase64()
	// Experimental.
	ResetKerberosKrb5Conf()
	// Experimental.
	ResetKerberosKrb5ConfBase64()
	// Experimental.
	ResetKerberosPrincipal()
	// Experimental.
	ResetKmsKeyProviderUri()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetQopConfiguration()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetReplicationFactor()
	// Experimental.
	ResetSimpleUser()
	// Experimental.
	ResetSubdirectory()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
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

// The jsii proxy struct for TfLocationHdfs
type jsiiProxy_TfLocationHdfs struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfLocationHdfs) AgentArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"agentArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) AgentArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"agentArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) BlockSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) BlockSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) KerberosKeytab() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosKeytab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) KerberosKeytabBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosKeytabBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) KerberosKeytabBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosKeytabBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) KerberosKeytabInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosKeytabInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) KerberosKrb5Conf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosKrb5Conf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) KerberosKrb5ConfBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosKrb5ConfBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) KerberosKrb5ConfBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosKrb5ConfBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) KerberosKrb5ConfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosKrb5ConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) KerberosPrincipal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) KerberosPrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosPrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) KmsKeyProviderUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyProviderUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) KmsKeyProviderUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyProviderUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) NameNode() TfLocationHdfs_NameNodePropertyList {
	var returns TfLocationHdfs_NameNodePropertyList
	_jsii_.Get(
		j,
		"nameNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) NameNodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nameNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) QopConfiguration() TfLocationHdfs_QopConfigurationPropertyOutputReference {
	var returns TfLocationHdfs_QopConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"qopConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) QopConfigurationInput() *TfLocationHdfs_QopConfigurationProperty {
	var returns *TfLocationHdfs_QopConfigurationProperty
	_jsii_.Get(
		j,
		"qopConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) ReplicationFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) ReplicationFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicationFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) SimpleUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"simpleUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) SimpleUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"simpleUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) Subdirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) SubdirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationHdfs) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs aws_datasync_location_hdfs} Resource.
// Experimental.
func NewTfLocationHdfs(scope constructs.Construct, id *string, config *TfLocationHdfsConfig) TfLocationHdfs {
	_init_.Initialize()

	if err := validateNewTfLocationHdfsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfLocationHdfs{}

	_jsii_.Create(
		"@cdktn/aws-datasync.TfLocationHdfs",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_hdfs aws_datasync_location_hdfs} Resource.
// Experimental.
func NewTfLocationHdfs_Override(t TfLocationHdfs, scope constructs.Construct, id *string, config *TfLocationHdfsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-datasync.TfLocationHdfs",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetAgentArns(val *[]*string) {
	if err := j.validateSetAgentArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentArns",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetBlockSize(val *float64) {
	if err := j.validateSetBlockSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockSize",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetKerberosKeytab(val *string) {
	if err := j.validateSetKerberosKeytabParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberosKeytab",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetKerberosKeytabBase64(val *string) {
	if err := j.validateSetKerberosKeytabBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberosKeytabBase64",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetKerberosKrb5Conf(val *string) {
	if err := j.validateSetKerberosKrb5ConfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberosKrb5Conf",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetKerberosKrb5ConfBase64(val *string) {
	if err := j.validateSetKerberosKrb5ConfBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberosKrb5ConfBase64",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetKerberosPrincipal(val *string) {
	if err := j.validateSetKerberosPrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberosPrincipal",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetKmsKeyProviderUri(val *string) {
	if err := j.validateSetKmsKeyProviderUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyProviderUri",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetReplicationFactor(val *float64) {
	if err := j.validateSetReplicationFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationFactor",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetSimpleUser(val *string) {
	if err := j.validateSetSimpleUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"simpleUser",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetSubdirectory(val *string) {
	if err := j.validateSetSubdirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subdirectory",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfLocationHdfs)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a TfLocationHdfs resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfLocationHdfs_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfLocationHdfs_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-datasync.TfLocationHdfs",
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
func TfLocationHdfs_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfLocationHdfs_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-datasync.TfLocationHdfs",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfLocationHdfs_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfLocationHdfs_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-datasync.TfLocationHdfs",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfLocationHdfs_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfLocationHdfs_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-datasync.TfLocationHdfs",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfLocationHdfs_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-datasync.TfLocationHdfs",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfLocationHdfs) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfLocationHdfs) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfLocationHdfs) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfLocationHdfs) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLocationHdfs) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfLocationHdfs) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfLocationHdfs) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfLocationHdfs) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfLocationHdfs) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfLocationHdfs) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfLocationHdfs) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfLocationHdfs) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLocationHdfs) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfLocationHdfs) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLocationHdfs) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfLocationHdfs) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfLocationHdfs) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfLocationHdfs) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfLocationHdfs) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfLocationHdfs) PutNameNode(value interface{}) {
	if err := t.validatePutNameNodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNameNode",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLocationHdfs) PutQopConfiguration(value *TfLocationHdfs_QopConfigurationProperty) {
	if err := t.validatePutQopConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putQopConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLocationHdfs) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfLocationHdfs) ResetAuthenticationType() {
	_jsii_.InvokeVoid(
		t,
		"resetAuthenticationType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLocationHdfs) ResetBlockSize() {
	_jsii_.InvokeVoid(
		t,
		"resetBlockSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLocationHdfs) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLocationHdfs) ResetKerberosKeytab() {
	_jsii_.InvokeVoid(
		t,
		"resetKerberosKeytab",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLocationHdfs) ResetKerberosKeytabBase64() {
	_jsii_.InvokeVoid(
		t,
		"resetKerberosKeytabBase64",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLocationHdfs) ResetKerberosKrb5Conf() {
	_jsii_.InvokeVoid(
		t,
		"resetKerberosKrb5Conf",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLocationHdfs) ResetKerberosKrb5ConfBase64() {
	_jsii_.InvokeVoid(
		t,
		"resetKerberosKrb5ConfBase64",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLocationHdfs) ResetKerberosPrincipal() {
	_jsii_.InvokeVoid(
		t,
		"resetKerberosPrincipal",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLocationHdfs) ResetKmsKeyProviderUri() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeyProviderUri",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLocationHdfs) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLocationHdfs) ResetQopConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetQopConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLocationHdfs) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLocationHdfs) ResetReplicationFactor() {
	_jsii_.InvokeVoid(
		t,
		"resetReplicationFactor",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLocationHdfs) ResetSimpleUser() {
	_jsii_.InvokeVoid(
		t,
		"resetSimpleUser",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLocationHdfs) ResetSubdirectory() {
	_jsii_.InvokeVoid(
		t,
		"resetSubdirectory",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLocationHdfs) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLocationHdfs) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLocationHdfs) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLocationHdfs) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLocationHdfs) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLocationHdfs) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLocationHdfs) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLocationHdfs) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLocationHdfs) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

