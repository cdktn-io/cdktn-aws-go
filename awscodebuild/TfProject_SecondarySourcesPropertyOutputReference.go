package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodebuild/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscodebuild/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfProject_SecondarySourcesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Auth() TfProject_SecondarySourcesAuthPropertyOutputReference
	// Experimental.
	AuthInput() *TfProject_SecondarySourcesAuthProperty
	// Experimental.
	Buildspec() *string
	// Experimental.
	SetBuildspec(val *string)
	// Experimental.
	BuildspecInput() *string
	// Experimental.
	BuildStatusConfig() TfProject_SecondarySourcesBuildStatusConfigPropertyOutputReference
	// Experimental.
	BuildStatusConfigInput() *TfProject_SecondarySourcesBuildStatusConfigProperty
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	GitCloneDepth() *float64
	// Experimental.
	SetGitCloneDepth(val *float64)
	// Experimental.
	GitCloneDepthInput() *float64
	// Experimental.
	GitSubmodulesConfig() TfProject_SecondarySourcesGitSubmodulesConfigPropertyOutputReference
	// Experimental.
	GitSubmodulesConfigInput() *TfProject_SecondarySourcesGitSubmodulesConfigProperty
	// Experimental.
	InsecureSsl() interface{}
	// Experimental.
	SetInsecureSsl(val interface{})
	// Experimental.
	InsecureSslInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Location() *string
	// Experimental.
	SetLocation(val *string)
	// Experimental.
	LocationInput() *string
	// Experimental.
	ReportBuildStatus() interface{}
	// Experimental.
	SetReportBuildStatus(val interface{})
	// Experimental.
	ReportBuildStatusInput() interface{}
	// Experimental.
	SourceIdentifier() *string
	// Experimental.
	SetSourceIdentifier(val *string)
	// Experimental.
	SourceIdentifierInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	PutAuth(value *TfProject_SecondarySourcesAuthProperty)
	// Experimental.
	PutBuildStatusConfig(value *TfProject_SecondarySourcesBuildStatusConfigProperty)
	// Experimental.
	PutGitSubmodulesConfig(value *TfProject_SecondarySourcesGitSubmodulesConfigProperty)
	// Experimental.
	ResetAuth()
	// Experimental.
	ResetBuildspec()
	// Experimental.
	ResetBuildStatusConfig()
	// Experimental.
	ResetGitCloneDepth()
	// Experimental.
	ResetGitSubmodulesConfig()
	// Experimental.
	ResetInsecureSsl()
	// Experimental.
	ResetLocation()
	// Experimental.
	ResetReportBuildStatus()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfProject_SecondarySourcesPropertyOutputReference
type jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) Auth() TfProject_SecondarySourcesAuthPropertyOutputReference {
	var returns TfProject_SecondarySourcesAuthPropertyOutputReference
	_jsii_.Get(
		j,
		"auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) AuthInput() *TfProject_SecondarySourcesAuthProperty {
	var returns *TfProject_SecondarySourcesAuthProperty
	_jsii_.Get(
		j,
		"authInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) Buildspec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildspec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) BuildspecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildspecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) BuildStatusConfig() TfProject_SecondarySourcesBuildStatusConfigPropertyOutputReference {
	var returns TfProject_SecondarySourcesBuildStatusConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"buildStatusConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) BuildStatusConfigInput() *TfProject_SecondarySourcesBuildStatusConfigProperty {
	var returns *TfProject_SecondarySourcesBuildStatusConfigProperty
	_jsii_.Get(
		j,
		"buildStatusConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) GitCloneDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitCloneDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) GitCloneDepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitCloneDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) GitSubmodulesConfig() TfProject_SecondarySourcesGitSubmodulesConfigPropertyOutputReference {
	var returns TfProject_SecondarySourcesGitSubmodulesConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"gitSubmodulesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) GitSubmodulesConfigInput() *TfProject_SecondarySourcesGitSubmodulesConfigProperty {
	var returns *TfProject_SecondarySourcesGitSubmodulesConfigProperty
	_jsii_.Get(
		j,
		"gitSubmodulesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) InsecureSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) InsecureSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureSslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) ReportBuildStatus() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reportBuildStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) ReportBuildStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reportBuildStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) SourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) SourceIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfProject_SecondarySourcesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfProject_SecondarySourcesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfProject_SecondarySourcesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codebuild.TfProject.SecondarySourcesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfProject_SecondarySourcesPropertyOutputReference_Override(t TfProject_SecondarySourcesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codebuild.TfProject.SecondarySourcesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference)SetBuildspec(val *string) {
	if err := j.validateSetBuildspecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildspec",
		val,
	)
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference)SetGitCloneDepth(val *float64) {
	if err := j.validateSetGitCloneDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitCloneDepth",
		val,
	)
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference)SetInsecureSsl(val interface{}) {
	if err := j.validateSetInsecureSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureSsl",
		val,
	)
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference)SetReportBuildStatus(val interface{}) {
	if err := j.validateSetReportBuildStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportBuildStatus",
		val,
	)
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference)SetSourceIdentifier(val *string) {
	if err := j.validateSetSourceIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIdentifier",
		val,
	)
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) PutAuth(value *TfProject_SecondarySourcesAuthProperty) {
	if err := t.validatePutAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAuth",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) PutBuildStatusConfig(value *TfProject_SecondarySourcesBuildStatusConfigProperty) {
	if err := t.validatePutBuildStatusConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBuildStatusConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) PutGitSubmodulesConfig(value *TfProject_SecondarySourcesGitSubmodulesConfigProperty) {
	if err := t.validatePutGitSubmodulesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGitSubmodulesConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) ResetAuth() {
	_jsii_.InvokeVoid(
		t,
		"resetAuth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) ResetBuildspec() {
	_jsii_.InvokeVoid(
		t,
		"resetBuildspec",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) ResetBuildStatusConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetBuildStatusConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) ResetGitCloneDepth() {
	_jsii_.InvokeVoid(
		t,
		"resetGitCloneDepth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) ResetGitSubmodulesConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetGitSubmodulesConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) ResetInsecureSsl() {
	_jsii_.InvokeVoid(
		t,
		"resetInsecureSsl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) ResetReportBuildStatus() {
	_jsii_.InvokeVoid(
		t,
		"resetReportBuildStatus",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProject_SecondarySourcesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

