package awsamplify

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsamplify/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsamplify/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfApp_AutoBranchCreationConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BasicAuthCredentials() *string
	// Experimental.
	SetBasicAuthCredentials(val *string)
	// Experimental.
	BasicAuthCredentialsInput() *string
	// Experimental.
	BuildSpec() *string
	// Experimental.
	SetBuildSpec(val *string)
	// Experimental.
	BuildSpecInput() *string
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
	EnableAutoBuild() interface{}
	// Experimental.
	SetEnableAutoBuild(val interface{})
	// Experimental.
	EnableAutoBuildInput() interface{}
	// Experimental.
	EnableBasicAuth() interface{}
	// Experimental.
	SetEnableBasicAuth(val interface{})
	// Experimental.
	EnableBasicAuthInput() interface{}
	// Experimental.
	EnablePerformanceMode() interface{}
	// Experimental.
	SetEnablePerformanceMode(val interface{})
	// Experimental.
	EnablePerformanceModeInput() interface{}
	// Experimental.
	EnablePullRequestPreview() interface{}
	// Experimental.
	SetEnablePullRequestPreview(val interface{})
	// Experimental.
	EnablePullRequestPreviewInput() interface{}
	// Experimental.
	EnvironmentVariables() *map[string]*string
	// Experimental.
	SetEnvironmentVariables(val *map[string]*string)
	// Experimental.
	EnvironmentVariablesInput() *map[string]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	Framework() *string
	// Experimental.
	SetFramework(val *string)
	// Experimental.
	FrameworkInput() *string
	// Experimental.
	InternalValue() *TfApp_AutoBranchCreationConfigProperty
	// Experimental.
	SetInternalValue(val *TfApp_AutoBranchCreationConfigProperty)
	// Experimental.
	PullRequestEnvironmentName() *string
	// Experimental.
	SetPullRequestEnvironmentName(val *string)
	// Experimental.
	PullRequestEnvironmentNameInput() *string
	// Experimental.
	Stage() *string
	// Experimental.
	SetStage(val *string)
	// Experimental.
	StageInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	ResetBasicAuthCredentials()
	// Experimental.
	ResetBuildSpec()
	// Experimental.
	ResetEnableAutoBuild()
	// Experimental.
	ResetEnableBasicAuth()
	// Experimental.
	ResetEnablePerformanceMode()
	// Experimental.
	ResetEnablePullRequestPreview()
	// Experimental.
	ResetEnvironmentVariables()
	// Experimental.
	ResetFramework()
	// Experimental.
	ResetPullRequestEnvironmentName()
	// Experimental.
	ResetStage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfApp_AutoBranchCreationConfigPropertyOutputReference
type jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) BasicAuthCredentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basicAuthCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) BasicAuthCredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basicAuthCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) BuildSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) BuildSpecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) EnableAutoBuild() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) EnableAutoBuildInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoBuildInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) EnableBasicAuth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBasicAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) EnableBasicAuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBasicAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) EnablePerformanceMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePerformanceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) EnablePerformanceModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePerformanceModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) EnablePullRequestPreview() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePullRequestPreview",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) EnablePullRequestPreviewInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePullRequestPreviewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) EnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) Framework() *string {
	var returns *string
	_jsii_.Get(
		j,
		"framework",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) FrameworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) InternalValue() *TfApp_AutoBranchCreationConfigProperty {
	var returns *TfApp_AutoBranchCreationConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) PullRequestEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pullRequestEnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) PullRequestEnvironmentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pullRequestEnvironmentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) StageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfApp_AutoBranchCreationConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfApp_AutoBranchCreationConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfApp_AutoBranchCreationConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-amplify.TfApp.AutoBranchCreationConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfApp_AutoBranchCreationConfigPropertyOutputReference_Override(t TfApp_AutoBranchCreationConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-amplify.TfApp.AutoBranchCreationConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference)SetBasicAuthCredentials(val *string) {
	if err := j.validateSetBasicAuthCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"basicAuthCredentials",
		val,
	)
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference)SetBuildSpec(val *string) {
	if err := j.validateSetBuildSpecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildSpec",
		val,
	)
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference)SetEnableAutoBuild(val interface{}) {
	if err := j.validateSetEnableAutoBuildParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutoBuild",
		val,
	)
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference)SetEnableBasicAuth(val interface{}) {
	if err := j.validateSetEnableBasicAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBasicAuth",
		val,
	)
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference)SetEnablePerformanceMode(val interface{}) {
	if err := j.validateSetEnablePerformanceModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePerformanceMode",
		val,
	)
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference)SetEnablePullRequestPreview(val interface{}) {
	if err := j.validateSetEnablePullRequestPreviewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePullRequestPreview",
		val,
	)
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference)SetEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference)SetFramework(val *string) {
	if err := j.validateSetFrameworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"framework",
		val,
	)
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference)SetInternalValue(val *TfApp_AutoBranchCreationConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference)SetPullRequestEnvironmentName(val *string) {
	if err := j.validateSetPullRequestEnvironmentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pullRequestEnvironmentName",
		val,
	)
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference)SetStage(val *string) {
	if err := j.validateSetStageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stage",
		val,
	)
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) ResetBasicAuthCredentials() {
	_jsii_.InvokeVoid(
		t,
		"resetBasicAuthCredentials",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) ResetBuildSpec() {
	_jsii_.InvokeVoid(
		t,
		"resetBuildSpec",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) ResetEnableAutoBuild() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableAutoBuild",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) ResetEnableBasicAuth() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableBasicAuth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) ResetEnablePerformanceMode() {
	_jsii_.InvokeVoid(
		t,
		"resetEnablePerformanceMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) ResetEnablePullRequestPreview() {
	_jsii_.InvokeVoid(
		t,
		"resetEnablePullRequestPreview",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		t,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) ResetFramework() {
	_jsii_.InvokeVoid(
		t,
		"resetFramework",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) ResetPullRequestEnvironmentName() {
	_jsii_.InvokeVoid(
		t,
		"resetPullRequestEnvironmentName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) ResetStage() {
	_jsii_.InvokeVoid(
		t,
		"resetStage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfApp_AutoBranchCreationConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

