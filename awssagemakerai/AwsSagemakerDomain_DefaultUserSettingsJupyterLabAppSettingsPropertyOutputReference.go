package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AppLifecycleManagement() AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference
	// Experimental.
	AppLifecycleManagementInput() *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty
	// Experimental.
	BuiltInLifecycleConfigArn() *string
	// Experimental.
	SetBuiltInLifecycleConfigArn(val *string)
	// Experimental.
	BuiltInLifecycleConfigArnInput() *string
	// Experimental.
	CodeRepository() AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsCodeRepositoryPropertyList
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
	CustomImage() AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsCustomImagePropertyList
	// Experimental.
	CustomImageInput() interface{}
	// Experimental.
	DefaultResourceSpec() AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference
	// Experimental.
	DefaultResourceSpecInput() *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpecProperty
	// Experimental.
	EmrSettings() AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference
	// Experimental.
	EmrSettingsInput() *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsProperty)
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
	PutAppLifecycleManagement(value *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty)
	// Experimental.
	PutCodeRepository(value interface{})
	// Experimental.
	PutCustomImage(value interface{})
	// Experimental.
	PutDefaultResourceSpec(value *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpecProperty)
	// Experimental.
	PutEmrSettings(value *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsProperty)
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

// The jsii proxy struct for AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference
type jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) AppLifecycleManagement() AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference {
	var returns AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference
	_jsii_.Get(
		j,
		"appLifecycleManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) AppLifecycleManagementInput() *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty {
	var returns *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty
	_jsii_.Get(
		j,
		"appLifecycleManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) BuiltInLifecycleConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInLifecycleConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) BuiltInLifecycleConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInLifecycleConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) CodeRepository() AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsCodeRepositoryPropertyList {
	var returns AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsCodeRepositoryPropertyList
	_jsii_.Get(
		j,
		"codeRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) CodeRepositoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) CustomImage() AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsCustomImagePropertyList {
	var returns AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsCustomImagePropertyList
	_jsii_.Get(
		j,
		"customImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) CustomImageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) DefaultResourceSpec() AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference {
	var returns AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference
	_jsii_.Get(
		j,
		"defaultResourceSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) DefaultResourceSpecInput() *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpecProperty {
	var returns *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpecProperty
	_jsii_.Get(
		j,
		"defaultResourceSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) EmrSettings() AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference {
	var returns AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"emrSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) EmrSettingsInput() *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsProperty {
	var returns *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsProperty
	_jsii_.Get(
		j,
		"emrSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) InternalValue() *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsProperty {
	var returns *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) LifecycleConfigArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) LifecycleConfigArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerDomain.DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference_Override(a AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerDomain.DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference)SetBuiltInLifecycleConfigArn(val *string) {
	if err := j.validateSetBuiltInLifecycleConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"builtInLifecycleConfigArn",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference)SetInternalValue(val *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference)SetLifecycleConfigArns(val *[]*string) {
	if err := j.validateSetLifecycleConfigArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleConfigArns",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) PutAppLifecycleManagement(value *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty) {
	if err := a.validatePutAppLifecycleManagementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAppLifecycleManagement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) PutCodeRepository(value interface{}) {
	if err := a.validatePutCodeRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeRepository",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) PutCustomImage(value interface{}) {
	if err := a.validatePutCustomImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomImage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) PutDefaultResourceSpec(value *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpecProperty) {
	if err := a.validatePutDefaultResourceSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultResourceSpec",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) PutEmrSettings(value *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsProperty) {
	if err := a.validatePutEmrSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEmrSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ResetAppLifecycleManagement() {
	_jsii_.InvokeVoid(
		a,
		"resetAppLifecycleManagement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ResetBuiltInLifecycleConfigArn() {
	_jsii_.InvokeVoid(
		a,
		"resetBuiltInLifecycleConfigArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ResetCodeRepository() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeRepository",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ResetCustomImage() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomImage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ResetDefaultResourceSpec() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultResourceSpec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ResetEmrSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetEmrSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ResetLifecycleConfigArns() {
	_jsii_.InvokeVoid(
		a,
		"resetLifecycleConfigArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

