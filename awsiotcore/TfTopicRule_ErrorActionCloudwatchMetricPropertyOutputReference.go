package awsiotcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsiotcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsiotcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference interface {
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
	InternalValue() *TfTopicRule_ErrorActionCloudwatchMetricProperty
	// Experimental.
	SetInternalValue(val *TfTopicRule_ErrorActionCloudwatchMetricProperty)
	// Experimental.
	MetricName() *string
	// Experimental.
	SetMetricName(val *string)
	// Experimental.
	MetricNameInput() *string
	// Experimental.
	MetricNamespace() *string
	// Experimental.
	SetMetricNamespace(val *string)
	// Experimental.
	MetricNamespaceInput() *string
	// Experimental.
	MetricTimestamp() *string
	// Experimental.
	SetMetricTimestamp(val *string)
	// Experimental.
	MetricTimestampInput() *string
	// Experimental.
	MetricUnit() *string
	// Experimental.
	SetMetricUnit(val *string)
	// Experimental.
	MetricUnitInput() *string
	// Experimental.
	MetricValue() *string
	// Experimental.
	SetMetricValue(val *string)
	// Experimental.
	MetricValueInput() *string
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
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
	ResetMetricTimestamp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference
type jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) InternalValue() *TfTopicRule_ErrorActionCloudwatchMetricProperty {
	var returns *TfTopicRule_ErrorActionCloudwatchMetricProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) MetricNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) MetricNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) MetricNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) MetricTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) MetricTimestampInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricTimestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) MetricUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) MetricUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) MetricValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) MetricValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-iot-core.TfTopicRule.ErrorActionCloudwatchMetricPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference_Override(t TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-iot-core.TfTopicRule.ErrorActionCloudwatchMetricPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetInternalValue(val *TfTopicRule_ErrorActionCloudwatchMetricProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetMetricName(val *string) {
	if err := j.validateSetMetricNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetMetricNamespace(val *string) {
	if err := j.validateSetMetricNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricNamespace",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetMetricTimestamp(val *string) {
	if err := j.validateSetMetricTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricTimestamp",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetMetricUnit(val *string) {
	if err := j.validateSetMetricUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricUnit",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetMetricValue(val *string) {
	if err := j.validateSetMetricValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricValue",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) ResetMetricTimestamp() {
	_jsii_.InvokeVoid(
		t,
		"resetMetricTimestamp",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

