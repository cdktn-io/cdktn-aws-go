package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdatasync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdatasync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTask_TaskReportConfigPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InternalValue() *TfTask_TaskReportConfigProperty
	// Experimental.
	SetInternalValue(val *TfTask_TaskReportConfigProperty)
	// Experimental.
	OutputType() *string
	// Experimental.
	SetOutputType(val *string)
	// Experimental.
	OutputTypeInput() *string
	// Experimental.
	ReportLevel() *string
	// Experimental.
	SetReportLevel(val *string)
	// Experimental.
	ReportLevelInput() *string
	// Experimental.
	ReportOverrides() TfTask_ReportOverridesPropertyOutputReference
	// Experimental.
	ReportOverridesInput() *TfTask_ReportOverridesProperty
	// Experimental.
	S3Destination() TfTask_S3DestinationPropertyOutputReference
	// Experimental.
	S3DestinationInput() *TfTask_S3DestinationProperty
	// Experimental.
	S3ObjectVersioning() *string
	// Experimental.
	SetS3ObjectVersioning(val *string)
	// Experimental.
	S3ObjectVersioningInput() *string
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
	PutReportOverrides(value *TfTask_ReportOverridesProperty)
	// Experimental.
	PutS3Destination(value *TfTask_S3DestinationProperty)
	// Experimental.
	ResetOutputType()
	// Experimental.
	ResetReportLevel()
	// Experimental.
	ResetReportOverrides()
	// Experimental.
	ResetS3ObjectVersioning()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTask_TaskReportConfigPropertyOutputReference
type jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) InternalValue() *TfTask_TaskReportConfigProperty {
	var returns *TfTask_TaskReportConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) OutputType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) OutputTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) ReportLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) ReportLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) ReportOverrides() TfTask_ReportOverridesPropertyOutputReference {
	var returns TfTask_ReportOverridesPropertyOutputReference
	_jsii_.Get(
		j,
		"reportOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) ReportOverridesInput() *TfTask_ReportOverridesProperty {
	var returns *TfTask_ReportOverridesProperty
	_jsii_.Get(
		j,
		"reportOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) S3Destination() TfTask_S3DestinationPropertyOutputReference {
	var returns TfTask_S3DestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) S3DestinationInput() *TfTask_S3DestinationProperty {
	var returns *TfTask_S3DestinationProperty
	_jsii_.Get(
		j,
		"s3DestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) S3ObjectVersioning() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectVersioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) S3ObjectVersioningInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectVersioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTask_TaskReportConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfTask_TaskReportConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTask_TaskReportConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-datasync.TfTask.TaskReportConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTask_TaskReportConfigPropertyOutputReference_Override(t TfTask_TaskReportConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-datasync.TfTask.TaskReportConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference)SetInternalValue(val *TfTask_TaskReportConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference)SetOutputType(val *string) {
	if err := j.validateSetOutputTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputType",
		val,
	)
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference)SetReportLevel(val *string) {
	if err := j.validateSetReportLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportLevel",
		val,
	)
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference)SetS3ObjectVersioning(val *string) {
	if err := j.validateSetS3ObjectVersioningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3ObjectVersioning",
		val,
	)
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) PutReportOverrides(value *TfTask_ReportOverridesProperty) {
	if err := t.validatePutReportOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putReportOverrides",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) PutS3Destination(value *TfTask_S3DestinationProperty) {
	if err := t.validatePutS3DestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3Destination",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) ResetOutputType() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) ResetReportLevel() {
	_jsii_.InvokeVoid(
		t,
		"resetReportLevel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) ResetReportOverrides() {
	_jsii_.InvokeVoid(
		t,
		"resetReportOverrides",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) ResetS3ObjectVersioning() {
	_jsii_.InvokeVoid(
		t,
		"resetS3ObjectVersioning",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTask_TaskReportConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

