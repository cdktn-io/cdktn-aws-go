package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AppLifecycleManagement() AwsUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference
	// Experimental.
	AppLifecycleManagementInput() *AwsUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty
	// Experimental.
	BuiltInLifecycleConfigArn() *string
	// Experimental.
	SetBuiltInLifecycleConfigArn(val *string)
	// Experimental.
	BuiltInLifecycleConfigArnInput() *string
	// Experimental.
	CodeRepository() AwsUserProfile_UserSettingsJupyterLabAppSettingsCodeRepositoryPropertyList
	// Experimental.
	CodeRepositoryInput() interface{}
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
	CustomImage() AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyList
	// Experimental.
	CustomImageInput() interface{}
	// Experimental.
	DefaultResourceSpec() AwsUserProfile_UserSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference
	// Experimental.
	DefaultResourceSpecInput() *AwsUserProfile_UserSettingsJupyterLabAppSettingsDefaultResourceSpecProperty
	// Experimental.
	EmrSettings() AwsUserProfile_EmrSettingsPropertyOutputReference
	// Experimental.
	EmrSettingsInput() *AwsUserProfile_EmrSettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsUserProfile_JupyterLabAppSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsUserProfile_JupyterLabAppSettingsProperty)
	// Experimental.
	LifecycleConfigArns() *[]*string
	// Experimental.
	SetLifecycleConfigArns(val *[]*string)
	// Experimental.
	LifecycleConfigArnsInput() *[]*string
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
	PutAppLifecycleManagement(value *AwsUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty)
	// Experimental.
	PutCodeRepository(value interface{})
	// Experimental.
	PutCustomImage(value interface{})
	// Experimental.
	PutDefaultResourceSpec(value *AwsUserProfile_UserSettingsJupyterLabAppSettingsDefaultResourceSpecProperty)
	// Experimental.
	PutEmrSettings(value *AwsUserProfile_EmrSettingsProperty)
	// Experimental.
	ResetAppLifecycleManagement()
	// Experimental.
	ResetBuiltInLifecycleConfigArn()
	// Experimental.
	ResetCodeRepository()
	// Experimental.
	ResetCustomImage()
	// Experimental.
	ResetDefaultResourceSpec()
	// Experimental.
	ResetEmrSettings()
	// Experimental.
	ResetLifecycleConfigArns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference
type jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) AppLifecycleManagement() AwsUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference {
	var returns AwsUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference
	_jsii_.Get(
		j,
		"appLifecycleManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) AppLifecycleManagementInput() *AwsUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty {
	var returns *AwsUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty
	_jsii_.Get(
		j,
		"appLifecycleManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) BuiltInLifecycleConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInLifecycleConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) BuiltInLifecycleConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInLifecycleConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) CodeRepository() AwsUserProfile_UserSettingsJupyterLabAppSettingsCodeRepositoryPropertyList {
	var returns AwsUserProfile_UserSettingsJupyterLabAppSettingsCodeRepositoryPropertyList
	_jsii_.Get(
		j,
		"codeRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) CodeRepositoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) CustomImage() AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyList {
	var returns AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyList
	_jsii_.Get(
		j,
		"customImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) CustomImageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) DefaultResourceSpec() AwsUserProfile_UserSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference {
	var returns AwsUserProfile_UserSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference
	_jsii_.Get(
		j,
		"defaultResourceSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) DefaultResourceSpecInput() *AwsUserProfile_UserSettingsJupyterLabAppSettingsDefaultResourceSpecProperty {
	var returns *AwsUserProfile_UserSettingsJupyterLabAppSettingsDefaultResourceSpecProperty
	_jsii_.Get(
		j,
		"defaultResourceSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) EmrSettings() AwsUserProfile_EmrSettingsPropertyOutputReference {
	var returns AwsUserProfile_EmrSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"emrSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) EmrSettingsInput() *AwsUserProfile_EmrSettingsProperty {
	var returns *AwsUserProfile_EmrSettingsProperty
	_jsii_.Get(
		j,
		"emrSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) InternalValue() *AwsUserProfile_JupyterLabAppSettingsProperty {
	var returns *AwsUserProfile_JupyterLabAppSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) LifecycleConfigArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) LifecycleConfigArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsUserProfile_JupyterLabAppSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsUserProfile_JupyterLabAppSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsUserProfile.JupyterLabAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsUserProfile_JupyterLabAppSettingsPropertyOutputReference_Override(a AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsUserProfile.JupyterLabAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference)SetBuiltInLifecycleConfigArn(val *string) {
	if err := j.validateSetBuiltInLifecycleConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"builtInLifecycleConfigArn",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference)SetInternalValue(val *AwsUserProfile_JupyterLabAppSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference)SetLifecycleConfigArns(val *[]*string) {
	if err := j.validateSetLifecycleConfigArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleConfigArns",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) PutAppLifecycleManagement(value *AwsUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty) {
	if err := a.validatePutAppLifecycleManagementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAppLifecycleManagement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) PutCodeRepository(value interface{}) {
	if err := a.validatePutCodeRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeRepository",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) PutCustomImage(value interface{}) {
	if err := a.validatePutCustomImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomImage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) PutDefaultResourceSpec(value *AwsUserProfile_UserSettingsJupyterLabAppSettingsDefaultResourceSpecProperty) {
	if err := a.validatePutDefaultResourceSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultResourceSpec",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) PutEmrSettings(value *AwsUserProfile_EmrSettingsProperty) {
	if err := a.validatePutEmrSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEmrSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) ResetAppLifecycleManagement() {
	_jsii_.InvokeVoid(
		a,
		"resetAppLifecycleManagement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) ResetBuiltInLifecycleConfigArn() {
	_jsii_.InvokeVoid(
		a,
		"resetBuiltInLifecycleConfigArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) ResetCodeRepository() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeRepository",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) ResetCustomImage() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomImage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) ResetDefaultResourceSpec() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultResourceSpec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) ResetEmrSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetEmrSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) ResetLifecycleConfigArns() {
	_jsii_.InvokeVoid(
		a,
		"resetLifecycleConfigArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

