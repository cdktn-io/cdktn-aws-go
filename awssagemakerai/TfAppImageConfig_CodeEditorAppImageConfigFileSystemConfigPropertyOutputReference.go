package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference interface {
	cdktn.ComplexObject
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
	DefaultGid() *float64
	// Experimental.
	SetDefaultGid(val *float64)
	// Experimental.
	DefaultGidInput() *float64
	// Experimental.
	DefaultUid() *float64
	// Experimental.
	SetDefaultUid(val *float64)
	// Experimental.
	DefaultUidInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigProperty
	// Experimental.
	SetInternalValue(val *TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigProperty)
	// Experimental.
	MountPath() *string
	// Experimental.
	SetMountPath(val *string)
	// Experimental.
	MountPathInput() *string
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
	ResetDefaultGid()
	// Experimental.
	ResetDefaultUid()
	// Experimental.
	ResetMountPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference
type jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) DefaultGid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultGid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) DefaultGidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultGidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) DefaultUid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) DefaultUidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultUidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) InternalValue() *TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigProperty {
	var returns *TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) MountPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) MountPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfAppImageConfig.CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference_Override(t TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfAppImageConfig.CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference)SetDefaultGid(val *float64) {
	if err := j.validateSetDefaultGidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultGid",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference)SetDefaultUid(val *float64) {
	if err := j.validateSetDefaultUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultUid",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference)SetInternalValue(val *TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference)SetMountPath(val *string) {
	if err := j.validateSetMountPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mountPath",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) ResetDefaultGid() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultGid",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) ResetDefaultUid() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultUid",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) ResetMountPath() {
	_jsii_.InvokeVoid(
		t,
		"resetMountPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

