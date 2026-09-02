package awstransferfamily

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awstransferfamily/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awstransferfamily/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference interface {
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
	FileSystemId() *string
	// Experimental.
	SetFileSystemId(val *string)
	// Experimental.
	FileSystemIdInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationProperty
	// Experimental.
	SetInternalValue(val *TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationProperty)
	// Experimental.
	Path() *string
	// Experimental.
	SetPath(val *string)
	// Experimental.
	PathInput() *string
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
	ResetFileSystemId()
	// Experimental.
	ResetPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference
type jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) FileSystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) InternalValue() *TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationProperty {
	var returns *TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-transfer-family.TfWorkflow.OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference_Override(t TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-transfer-family.TfWorkflow.OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference)SetFileSystemId(val *string) {
	if err := j.validateSetFileSystemIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileSystemId",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference)SetInternalValue(val *TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) ResetFileSystemId() {
	_jsii_.InvokeVoid(
		t,
		"resetFileSystemId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		t,
		"resetPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

