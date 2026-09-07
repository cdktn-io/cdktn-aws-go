package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AppLifecycleManagement() AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference
	// Experimental.
	AppLifecycleManagementInput() *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementProperty
	// Experimental.
	BuiltInLifecycleConfigArn() *string
	// Experimental.
	SetBuiltInLifecycleConfigArn(val *string)
	// Experimental.
	BuiltInLifecycleConfigArnInput() *string
	// Experimental.
	CodeRepository() AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsCodeRepositoryPropertyList
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
	CustomImage() AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsCustomImagePropertyList
	// Experimental.
	CustomImageInput() interface{}
	// Experimental.
	DefaultResourceSpec() AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference
	// Experimental.
	DefaultResourceSpecInput() *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecProperty
	// Experimental.
	EmrSettings() AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference
	// Experimental.
	EmrSettingsInput() *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsEmrSettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsProperty)
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
	PutAppLifecycleManagement(value *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementProperty)
	// Experimental.
	PutCodeRepository(value interface{})
	// Experimental.
	PutCustomImage(value interface{})
	// Experimental.
	PutDefaultResourceSpec(value *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecProperty)
	// Experimental.
	PutEmrSettings(value *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsEmrSettingsProperty)
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

// The jsii proxy struct for AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference
type jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) AppLifecycleManagement() AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference {
	var returns AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference
	_jsii_.Get(
		j,
		"appLifecycleManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) AppLifecycleManagementInput() *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementProperty {
	var returns *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementProperty
	_jsii_.Get(
		j,
		"appLifecycleManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) BuiltInLifecycleConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInLifecycleConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) BuiltInLifecycleConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInLifecycleConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) CodeRepository() AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsCodeRepositoryPropertyList {
	var returns AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsCodeRepositoryPropertyList
	_jsii_.Get(
		j,
		"codeRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) CodeRepositoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) CustomImage() AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsCustomImagePropertyList {
	var returns AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsCustomImagePropertyList
	_jsii_.Get(
		j,
		"customImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) CustomImageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) DefaultResourceSpec() AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference {
	var returns AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference
	_jsii_.Get(
		j,
		"defaultResourceSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) DefaultResourceSpecInput() *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecProperty {
	var returns *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecProperty
	_jsii_.Get(
		j,
		"defaultResourceSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) EmrSettings() AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference {
	var returns AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"emrSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) EmrSettingsInput() *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsEmrSettingsProperty {
	var returns *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsEmrSettingsProperty
	_jsii_.Get(
		j,
		"emrSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) InternalValue() *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsProperty {
	var returns *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) LifecycleConfigArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) LifecycleConfigArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsDomain.DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference_Override(a AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsDomain.DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference)SetBuiltInLifecycleConfigArn(val *string) {
	if err := j.validateSetBuiltInLifecycleConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"builtInLifecycleConfigArn",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference)SetInternalValue(val *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference)SetLifecycleConfigArns(val *[]*string) {
	if err := j.validateSetLifecycleConfigArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleConfigArns",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) PutAppLifecycleManagement(value *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementProperty) {
	if err := a.validatePutAppLifecycleManagementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAppLifecycleManagement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) PutCodeRepository(value interface{}) {
	if err := a.validatePutCodeRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeRepository",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) PutCustomImage(value interface{}) {
	if err := a.validatePutCustomImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomImage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) PutDefaultResourceSpec(value *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecProperty) {
	if err := a.validatePutDefaultResourceSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultResourceSpec",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) PutEmrSettings(value *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsEmrSettingsProperty) {
	if err := a.validatePutEmrSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEmrSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) ResetAppLifecycleManagement() {
	_jsii_.InvokeVoid(
		a,
		"resetAppLifecycleManagement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) ResetBuiltInLifecycleConfigArn() {
	_jsii_.InvokeVoid(
		a,
		"resetBuiltInLifecycleConfigArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) ResetCodeRepository() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeRepository",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) ResetCustomImage() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomImage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) ResetDefaultResourceSpec() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultResourceSpec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) ResetEmrSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetEmrSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) ResetLifecycleConfigArns() {
	_jsii_.InvokeVoid(
		a,
		"resetLifecycleConfigArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

