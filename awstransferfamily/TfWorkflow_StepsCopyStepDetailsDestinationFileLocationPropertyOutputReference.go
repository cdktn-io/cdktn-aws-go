package awstransferfamily

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awstransferfamily/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awstransferfamily/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference interface {
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
	EfsFileLocation() TfWorkflow_StepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference
	// Experimental.
	EfsFileLocationInput() *TfWorkflow_StepsCopyStepDetailsDestinationFileLocationEfsFileLocationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfWorkflow_StepsCopyStepDetailsDestinationFileLocationProperty
	// Experimental.
	SetInternalValue(val *TfWorkflow_StepsCopyStepDetailsDestinationFileLocationProperty)
	// Experimental.
	S3FileLocation() TfWorkflow_StepsCopyStepDetailsDestinationFileLocationS3FileLocationPropertyOutputReference
	// Experimental.
	S3FileLocationInput() *TfWorkflow_StepsCopyStepDetailsDestinationFileLocationS3FileLocationProperty
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
	PutEfsFileLocation(value *TfWorkflow_StepsCopyStepDetailsDestinationFileLocationEfsFileLocationProperty)
	// Experimental.
	PutS3FileLocation(value *TfWorkflow_StepsCopyStepDetailsDestinationFileLocationS3FileLocationProperty)
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

// The jsii proxy struct for TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference
type jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) EfsFileLocation() TfWorkflow_StepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference {
	var returns TfWorkflow_StepsCopyStepDetailsDestinationFileLocationEfsFileLocationPropertyOutputReference
	_jsii_.Get(
		j,
		"efsFileLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) EfsFileLocationInput() *TfWorkflow_StepsCopyStepDetailsDestinationFileLocationEfsFileLocationProperty {
	var returns *TfWorkflow_StepsCopyStepDetailsDestinationFileLocationEfsFileLocationProperty
	_jsii_.Get(
		j,
		"efsFileLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) InternalValue() *TfWorkflow_StepsCopyStepDetailsDestinationFileLocationProperty {
	var returns *TfWorkflow_StepsCopyStepDetailsDestinationFileLocationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) S3FileLocation() TfWorkflow_StepsCopyStepDetailsDestinationFileLocationS3FileLocationPropertyOutputReference {
	var returns TfWorkflow_StepsCopyStepDetailsDestinationFileLocationS3FileLocationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3FileLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) S3FileLocationInput() *TfWorkflow_StepsCopyStepDetailsDestinationFileLocationS3FileLocationProperty {
	var returns *TfWorkflow_StepsCopyStepDetailsDestinationFileLocationS3FileLocationProperty
	_jsii_.Get(
		j,
		"s3FileLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-transfer-family.TfWorkflow.StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference_Override(t TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-transfer-family.TfWorkflow.StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference)SetInternalValue(val *TfWorkflow_StepsCopyStepDetailsDestinationFileLocationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) PutEfsFileLocation(value *TfWorkflow_StepsCopyStepDetailsDestinationFileLocationEfsFileLocationProperty) {
	if err := t.validatePutEfsFileLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEfsFileLocation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) PutS3FileLocation(value *TfWorkflow_StepsCopyStepDetailsDestinationFileLocationS3FileLocationProperty) {
	if err := t.validatePutS3FileLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3FileLocation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) ResetEfsFileLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetEfsFileLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) ResetS3FileLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetS3FileLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWorkflow_StepsCopyStepDetailsDestinationFileLocationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

