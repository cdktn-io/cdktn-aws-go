package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodebuild/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscodebuild/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCodebuildProject_SecondarySourcesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Auth() AwsCodebuildProject_SecondarySourcesAuthPropertyOutputReference
	// Experimental.
	AuthInput() *AwsCodebuildProject_SecondarySourcesAuthProperty
	// Experimental.
	Buildspec() *string
	// Experimental.
	SetBuildspec(val *string)
	// Experimental.
	BuildspecInput() *string
	// Experimental.
	BuildStatusConfig() AwsCodebuildProject_SecondarySourcesBuildStatusConfigPropertyOutputReference
	// Experimental.
	BuildStatusConfigInput() *AwsCodebuildProject_SecondarySourcesBuildStatusConfigProperty
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
	GitSubmodulesConfig() AwsCodebuildProject_SecondarySourcesGitSubmodulesConfigPropertyOutputReference
	// Experimental.
	GitSubmodulesConfigInput() *AwsCodebuildProject_SecondarySourcesGitSubmodulesConfigProperty
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
	PutAuth(value *AwsCodebuildProject_SecondarySourcesAuthProperty)
	// Experimental.
	PutBuildStatusConfig(value *AwsCodebuildProject_SecondarySourcesBuildStatusConfigProperty)
	// Experimental.
	PutGitSubmodulesConfig(value *AwsCodebuildProject_SecondarySourcesGitSubmodulesConfigProperty)
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

// The jsii proxy struct for AwsCodebuildProject_SecondarySourcesPropertyOutputReference
type jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) Auth() AwsCodebuildProject_SecondarySourcesAuthPropertyOutputReference {
	var returns AwsCodebuildProject_SecondarySourcesAuthPropertyOutputReference
	_jsii_.Get(
		j,
		"auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) AuthInput() *AwsCodebuildProject_SecondarySourcesAuthProperty {
	var returns *AwsCodebuildProject_SecondarySourcesAuthProperty
	_jsii_.Get(
		j,
		"authInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) Buildspec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildspec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) BuildspecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildspecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) BuildStatusConfig() AwsCodebuildProject_SecondarySourcesBuildStatusConfigPropertyOutputReference {
	var returns AwsCodebuildProject_SecondarySourcesBuildStatusConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"buildStatusConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) BuildStatusConfigInput() *AwsCodebuildProject_SecondarySourcesBuildStatusConfigProperty {
	var returns *AwsCodebuildProject_SecondarySourcesBuildStatusConfigProperty
	_jsii_.Get(
		j,
		"buildStatusConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) GitCloneDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitCloneDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) GitCloneDepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitCloneDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) GitSubmodulesConfig() AwsCodebuildProject_SecondarySourcesGitSubmodulesConfigPropertyOutputReference {
	var returns AwsCodebuildProject_SecondarySourcesGitSubmodulesConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"gitSubmodulesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) GitSubmodulesConfigInput() *AwsCodebuildProject_SecondarySourcesGitSubmodulesConfigProperty {
	var returns *AwsCodebuildProject_SecondarySourcesGitSubmodulesConfigProperty
	_jsii_.Get(
		j,
		"gitSubmodulesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) InsecureSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) InsecureSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureSslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) ReportBuildStatus() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reportBuildStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) ReportBuildStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reportBuildStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) SourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) SourceIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCodebuildProject_SecondarySourcesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsCodebuildProject_SecondarySourcesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCodebuildProject_SecondarySourcesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codebuild.AwsCodebuildProject.SecondarySourcesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCodebuildProject_SecondarySourcesPropertyOutputReference_Override(a AwsCodebuildProject_SecondarySourcesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codebuild.AwsCodebuildProject.SecondarySourcesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference)SetBuildspec(val *string) {
	if err := j.validateSetBuildspecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildspec",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference)SetGitCloneDepth(val *float64) {
	if err := j.validateSetGitCloneDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitCloneDepth",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference)SetInsecureSsl(val interface{}) {
	if err := j.validateSetInsecureSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureSsl",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference)SetReportBuildStatus(val interface{}) {
	if err := j.validateSetReportBuildStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportBuildStatus",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference)SetSourceIdentifier(val *string) {
	if err := j.validateSetSourceIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) PutAuth(value *AwsCodebuildProject_SecondarySourcesAuthProperty) {
	if err := a.validatePutAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuth",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) PutBuildStatusConfig(value *AwsCodebuildProject_SecondarySourcesBuildStatusConfigProperty) {
	if err := a.validatePutBuildStatusConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBuildStatusConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) PutGitSubmodulesConfig(value *AwsCodebuildProject_SecondarySourcesGitSubmodulesConfigProperty) {
	if err := a.validatePutGitSubmodulesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGitSubmodulesConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) ResetAuth() {
	_jsii_.InvokeVoid(
		a,
		"resetAuth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) ResetBuildspec() {
	_jsii_.InvokeVoid(
		a,
		"resetBuildspec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) ResetBuildStatusConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetBuildStatusConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) ResetGitCloneDepth() {
	_jsii_.InvokeVoid(
		a,
		"resetGitCloneDepth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) ResetGitSubmodulesConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetGitSubmodulesConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) ResetInsecureSsl() {
	_jsii_.InvokeVoid(
		a,
		"resetInsecureSsl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) ResetReportBuildStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetReportBuildStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCodebuildProject_SecondarySourcesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

