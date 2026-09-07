package codebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/codebuild/jsii"

	"github.com/cdktn-io/cdktn-aws-go/codebuild/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsProject_SourcePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Auth() AwsProject_SourceAuthPropertyOutputReference
	// Experimental.
	AuthInput() *AwsProject_SourceAuthProperty
	// Experimental.
	Buildspec() *string
	// Experimental.
	SetBuildspec(val *string)
	// Experimental.
	BuildspecInput() *string
	// Experimental.
	BuildStatusConfig() AwsProject_SourceBuildStatusConfigPropertyOutputReference
	// Experimental.
	BuildStatusConfigInput() *AwsProject_SourceBuildStatusConfigProperty
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
	GitSubmodulesConfig() AwsProject_SourceGitSubmodulesConfigPropertyOutputReference
	// Experimental.
	GitSubmodulesConfigInput() *AwsProject_SourceGitSubmodulesConfigProperty
	// Experimental.
	InsecureSsl() interface{}
	// Experimental.
	SetInsecureSsl(val interface{})
	// Experimental.
	InsecureSslInput() interface{}
	// Experimental.
	InternalValue() *AwsProject_SourceProperty
	// Experimental.
	SetInternalValue(val *AwsProject_SourceProperty)
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
	PutAuth(value *AwsProject_SourceAuthProperty)
	// Experimental.
	PutBuildStatusConfig(value *AwsProject_SourceBuildStatusConfigProperty)
	// Experimental.
	PutGitSubmodulesConfig(value *AwsProject_SourceGitSubmodulesConfigProperty)
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

// The jsii proxy struct for AwsProject_SourcePropertyOutputReference
type jsiiProxy_AwsProject_SourcePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) Auth() AwsProject_SourceAuthPropertyOutputReference {
	var returns AwsProject_SourceAuthPropertyOutputReference
	_jsii_.Get(
		j,
		"auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) AuthInput() *AwsProject_SourceAuthProperty {
	var returns *AwsProject_SourceAuthProperty
	_jsii_.Get(
		j,
		"authInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) Buildspec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildspec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) BuildspecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildspecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) BuildStatusConfig() AwsProject_SourceBuildStatusConfigPropertyOutputReference {
	var returns AwsProject_SourceBuildStatusConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"buildStatusConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) BuildStatusConfigInput() *AwsProject_SourceBuildStatusConfigProperty {
	var returns *AwsProject_SourceBuildStatusConfigProperty
	_jsii_.Get(
		j,
		"buildStatusConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) GitCloneDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitCloneDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) GitCloneDepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitCloneDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) GitSubmodulesConfig() AwsProject_SourceGitSubmodulesConfigPropertyOutputReference {
	var returns AwsProject_SourceGitSubmodulesConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"gitSubmodulesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) GitSubmodulesConfigInput() *AwsProject_SourceGitSubmodulesConfigProperty {
	var returns *AwsProject_SourceGitSubmodulesConfigProperty
	_jsii_.Get(
		j,
		"gitSubmodulesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) InsecureSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) InsecureSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureSslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) InternalValue() *AwsProject_SourceProperty {
	var returns *AwsProject_SourceProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) ReportBuildStatus() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reportBuildStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) ReportBuildStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reportBuildStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsProject_SourcePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsProject_SourcePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsProject_SourcePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsProject_SourcePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codebuild.AwsProject.SourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsProject_SourcePropertyOutputReference_Override(a AwsProject_SourcePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codebuild.AwsProject.SourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference)SetBuildspec(val *string) {
	if err := j.validateSetBuildspecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildspec",
		val,
	)
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference)SetGitCloneDepth(val *float64) {
	if err := j.validateSetGitCloneDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitCloneDepth",
		val,
	)
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference)SetInsecureSsl(val interface{}) {
	if err := j.validateSetInsecureSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureSsl",
		val,
	)
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference)SetInternalValue(val *AwsProject_SourceProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference)SetReportBuildStatus(val interface{}) {
	if err := j.validateSetReportBuildStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportBuildStatus",
		val,
	)
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsProject_SourcePropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) PutAuth(value *AwsProject_SourceAuthProperty) {
	if err := a.validatePutAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuth",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) PutBuildStatusConfig(value *AwsProject_SourceBuildStatusConfigProperty) {
	if err := a.validatePutBuildStatusConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBuildStatusConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) PutGitSubmodulesConfig(value *AwsProject_SourceGitSubmodulesConfigProperty) {
	if err := a.validatePutGitSubmodulesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGitSubmodulesConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) ResetAuth() {
	_jsii_.InvokeVoid(
		a,
		"resetAuth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) ResetBuildspec() {
	_jsii_.InvokeVoid(
		a,
		"resetBuildspec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) ResetBuildStatusConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetBuildStatusConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) ResetGitCloneDepth() {
	_jsii_.InvokeVoid(
		a,
		"resetGitCloneDepth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) ResetGitSubmodulesConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetGitSubmodulesConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) ResetInsecureSsl() {
	_jsii_.InvokeVoid(
		a,
		"resetInsecureSsl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) ResetReportBuildStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetReportBuildStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProject_SourcePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

