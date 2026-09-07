package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AppLifecycleManagement() AwsUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference
	// Experimental.
	AppLifecycleManagementInput() *AwsUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementProperty
	// Experimental.
	BuiltInLifecycleConfigArn() *string
	// Experimental.
	SetBuiltInLifecycleConfigArn(val *string)
	// Experimental.
	BuiltInLifecycleConfigArnInput() *string
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
	CustomImage() AwsUserProfile_UserSettingsCodeEditorAppSettingsCustomImagePropertyList
	// Experimental.
	CustomImageInput() interface{}
	// Experimental.
	DefaultResourceSpec() AwsUserProfile_UserSettingsCodeEditorAppSettingsDefaultResourceSpecPropertyOutputReference
	// Experimental.
	DefaultResourceSpecInput() *AwsUserProfile_UserSettingsCodeEditorAppSettingsDefaultResourceSpecProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsUserProfile_CodeEditorAppSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsUserProfile_CodeEditorAppSettingsProperty)
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
	PutAppLifecycleManagement(value *AwsUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementProperty)
	// Experimental.
	PutCustomImage(value interface{})
	// Experimental.
	PutDefaultResourceSpec(value *AwsUserProfile_UserSettingsCodeEditorAppSettingsDefaultResourceSpecProperty)
	// Experimental.
	ResetAppLifecycleManagement()
	// Experimental.
	ResetBuiltInLifecycleConfigArn()
	// Experimental.
	ResetCustomImage()
	// Experimental.
	ResetDefaultResourceSpec()
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

// The jsii proxy struct for AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference
type jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) AppLifecycleManagement() AwsUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference {
	var returns AwsUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference
	_jsii_.Get(
		j,
		"appLifecycleManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) AppLifecycleManagementInput() *AwsUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementProperty {
	var returns *AwsUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementProperty
	_jsii_.Get(
		j,
		"appLifecycleManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) BuiltInLifecycleConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInLifecycleConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) BuiltInLifecycleConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInLifecycleConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) CustomImage() AwsUserProfile_UserSettingsCodeEditorAppSettingsCustomImagePropertyList {
	var returns AwsUserProfile_UserSettingsCodeEditorAppSettingsCustomImagePropertyList
	_jsii_.Get(
		j,
		"customImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) CustomImageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) DefaultResourceSpec() AwsUserProfile_UserSettingsCodeEditorAppSettingsDefaultResourceSpecPropertyOutputReference {
	var returns AwsUserProfile_UserSettingsCodeEditorAppSettingsDefaultResourceSpecPropertyOutputReference
	_jsii_.Get(
		j,
		"defaultResourceSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) DefaultResourceSpecInput() *AwsUserProfile_UserSettingsCodeEditorAppSettingsDefaultResourceSpecProperty {
	var returns *AwsUserProfile_UserSettingsCodeEditorAppSettingsDefaultResourceSpecProperty
	_jsii_.Get(
		j,
		"defaultResourceSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) InternalValue() *AwsUserProfile_CodeEditorAppSettingsProperty {
	var returns *AwsUserProfile_CodeEditorAppSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) LifecycleConfigArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) LifecycleConfigArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsUserProfile_CodeEditorAppSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsUserProfile_CodeEditorAppSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsUserProfile.CodeEditorAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsUserProfile_CodeEditorAppSettingsPropertyOutputReference_Override(a AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsUserProfile.CodeEditorAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference)SetBuiltInLifecycleConfigArn(val *string) {
	if err := j.validateSetBuiltInLifecycleConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"builtInLifecycleConfigArn",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference)SetInternalValue(val *AwsUserProfile_CodeEditorAppSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference)SetLifecycleConfigArns(val *[]*string) {
	if err := j.validateSetLifecycleConfigArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleConfigArns",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) PutAppLifecycleManagement(value *AwsUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementProperty) {
	if err := a.validatePutAppLifecycleManagementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAppLifecycleManagement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) PutCustomImage(value interface{}) {
	if err := a.validatePutCustomImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomImage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) PutDefaultResourceSpec(value *AwsUserProfile_UserSettingsCodeEditorAppSettingsDefaultResourceSpecProperty) {
	if err := a.validatePutDefaultResourceSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultResourceSpec",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) ResetAppLifecycleManagement() {
	_jsii_.InvokeVoid(
		a,
		"resetAppLifecycleManagement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) ResetBuiltInLifecycleConfigArn() {
	_jsii_.InvokeVoid(
		a,
		"resetBuiltInLifecycleConfigArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) ResetCustomImage() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomImage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) ResetDefaultResourceSpec() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultResourceSpec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) ResetLifecycleConfigArns() {
	_jsii_.InvokeVoid(
		a,
		"resetLifecycleConfigArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

