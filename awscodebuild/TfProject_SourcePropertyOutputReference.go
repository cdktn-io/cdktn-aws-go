package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodebuild/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscodebuild/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfProject_SourcePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Auth() TfProject_SourceAuthPropertyOutputReference
	// Experimental.
	AuthInput() *TfProject_SourceAuthProperty
	// Experimental.
	Buildspec() *string
	// Experimental.
	SetBuildspec(val *string)
	// Experimental.
	BuildspecInput() *string
	// Experimental.
	BuildStatusConfig() TfProject_SourceBuildStatusConfigPropertyOutputReference
	// Experimental.
	BuildStatusConfigInput() *TfProject_SourceBuildStatusConfigProperty
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
	GitSubmodulesConfig() TfProject_SourceGitSubmodulesConfigPropertyOutputReference
	// Experimental.
	GitSubmodulesConfigInput() *TfProject_SourceGitSubmodulesConfigProperty
	// Experimental.
	InsecureSsl() interface{}
	// Experimental.
	SetInsecureSsl(val interface{})
	// Experimental.
	InsecureSslInput() interface{}
	// Experimental.
	InternalValue() *TfProject_SourceProperty
	// Experimental.
	SetInternalValue(val *TfProject_SourceProperty)
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
	PutAuth(value *TfProject_SourceAuthProperty)
	// Experimental.
	PutBuildStatusConfig(value *TfProject_SourceBuildStatusConfigProperty)
	// Experimental.
	PutGitSubmodulesConfig(value *TfProject_SourceGitSubmodulesConfigProperty)
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

// The jsii proxy struct for TfProject_SourcePropertyOutputReference
type jsiiProxy_TfProject_SourcePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) Auth() TfProject_SourceAuthPropertyOutputReference {
	var returns TfProject_SourceAuthPropertyOutputReference
	_jsii_.Get(
		j,
		"auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) AuthInput() *TfProject_SourceAuthProperty {
	var returns *TfProject_SourceAuthProperty
	_jsii_.Get(
		j,
		"authInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) Buildspec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildspec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) BuildspecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildspecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) BuildStatusConfig() TfProject_SourceBuildStatusConfigPropertyOutputReference {
	var returns TfProject_SourceBuildStatusConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"buildStatusConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) BuildStatusConfigInput() *TfProject_SourceBuildStatusConfigProperty {
	var returns *TfProject_SourceBuildStatusConfigProperty
	_jsii_.Get(
		j,
		"buildStatusConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) GitCloneDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitCloneDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) GitCloneDepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitCloneDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) GitSubmodulesConfig() TfProject_SourceGitSubmodulesConfigPropertyOutputReference {
	var returns TfProject_SourceGitSubmodulesConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"gitSubmodulesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) GitSubmodulesConfigInput() *TfProject_SourceGitSubmodulesConfigProperty {
	var returns *TfProject_SourceGitSubmodulesConfigProperty
	_jsii_.Get(
		j,
		"gitSubmodulesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) InsecureSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) InsecureSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureSslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) InternalValue() *TfProject_SourceProperty {
	var returns *TfProject_SourceProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) ReportBuildStatus() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reportBuildStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) ReportBuildStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reportBuildStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfProject_SourcePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfProject_SourcePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfProject_SourcePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfProject_SourcePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codebuild.TfProject.SourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfProject_SourcePropertyOutputReference_Override(t TfProject_SourcePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codebuild.TfProject.SourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference)SetBuildspec(val *string) {
	if err := j.validateSetBuildspecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildspec",
		val,
	)
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference)SetGitCloneDepth(val *float64) {
	if err := j.validateSetGitCloneDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitCloneDepth",
		val,
	)
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference)SetInsecureSsl(val interface{}) {
	if err := j.validateSetInsecureSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureSsl",
		val,
	)
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference)SetInternalValue(val *TfProject_SourceProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference)SetReportBuildStatus(val interface{}) {
	if err := j.validateSetReportBuildStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportBuildStatus",
		val,
	)
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfProject_SourcePropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) PutAuth(value *TfProject_SourceAuthProperty) {
	if err := t.validatePutAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAuth",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) PutBuildStatusConfig(value *TfProject_SourceBuildStatusConfigProperty) {
	if err := t.validatePutBuildStatusConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBuildStatusConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) PutGitSubmodulesConfig(value *TfProject_SourceGitSubmodulesConfigProperty) {
	if err := t.validatePutGitSubmodulesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGitSubmodulesConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) ResetAuth() {
	_jsii_.InvokeVoid(
		t,
		"resetAuth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) ResetBuildspec() {
	_jsii_.InvokeVoid(
		t,
		"resetBuildspec",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) ResetBuildStatusConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetBuildStatusConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) ResetGitCloneDepth() {
	_jsii_.InvokeVoid(
		t,
		"resetGitCloneDepth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) ResetGitSubmodulesConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetGitSubmodulesConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) ResetInsecureSsl() {
	_jsii_.InvokeVoid(
		t,
		"resetInsecureSsl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) ResetReportBuildStatus() {
	_jsii_.InvokeVoid(
		t,
		"resetReportBuildStatus",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfProject_SourcePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

