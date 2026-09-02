package awstransferfamily

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awstransferfamily/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awstransferfamily/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference interface {
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
	EfsFileLocation() TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference
	// Experimental.
	EfsFileLocationInput() *TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationEfsFileLocationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationProperty
	// Experimental.
	SetInternalValue(val *TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationProperty)
	// Experimental.
	S3FileLocation() TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationS3FileLocationPropertyOutputReference
	// Experimental.
	S3FileLocationInput() *TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationS3FileLocationProperty
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
	PutEfsFileLocation(value *TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationEfsFileLocationProperty)
	// Experimental.
	PutS3FileLocation(value *TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationS3FileLocationProperty)
	// Experimental.
	ResetEfsFileLocation()
	// Experimental.
	ResetS3FileLocation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference
type jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) EfsFileLocation() TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference {
	var returns TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference
	_jsii_.Get(
		j,
		"efsFileLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) EfsFileLocationInput() *TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationEfsFileLocationProperty {
	var returns *TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationEfsFileLocationProperty
	_jsii_.Get(
		j,
		"efsFileLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) InternalValue() *TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationProperty {
	var returns *TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) S3FileLocation() TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationS3FileLocationPropertyOutputReference {
	var returns TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationS3FileLocationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3FileLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) S3FileLocationInput() *TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationS3FileLocationProperty {
	var returns *TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationS3FileLocationProperty
	_jsii_.Get(
		j,
		"s3FileLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-transfer-family.TfWorkflow.OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference_Override(t TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-transfer-family.TfWorkflow.OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference)SetInternalValue(val *TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) PutEfsFileLocation(value *TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationEfsFileLocationProperty) {
	if err := t.validatePutEfsFileLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEfsFileLocation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) PutS3FileLocation(value *TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationS3FileLocationProperty) {
	if err := t.validatePutS3FileLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3FileLocation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) ResetEfsFileLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetEfsFileLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) ResetS3FileLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetS3FileLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

