package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference interface {
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
	// Experimental.
	ConstraintsResource() TfDataQualityJobDefinition_ConstraintsResourcePropertyOutputReference
	// Experimental.
	ConstraintsResourceInput() *TfDataQualityJobDefinition_ConstraintsResourceProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDataQualityJobDefinition_DataQualityBaselineConfigProperty
	// Experimental.
	SetInternalValue(val *TfDataQualityJobDefinition_DataQualityBaselineConfigProperty)
	// Experimental.
	StatisticsResource() TfDataQualityJobDefinition_StatisticsResourcePropertyOutputReference
	// Experimental.
	StatisticsResourceInput() *TfDataQualityJobDefinition_StatisticsResourceProperty
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
	PutConstraintsResource(value *TfDataQualityJobDefinition_ConstraintsResourceProperty)
	// Experimental.
	PutStatisticsResource(value *TfDataQualityJobDefinition_StatisticsResourceProperty)
	// Experimental.
	ResetConstraintsResource()
	// Experimental.
	ResetStatisticsResource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference
type jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) ConstraintsResource() TfDataQualityJobDefinition_ConstraintsResourcePropertyOutputReference {
	var returns TfDataQualityJobDefinition_ConstraintsResourcePropertyOutputReference
	_jsii_.Get(
		j,
		"constraintsResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) ConstraintsResourceInput() *TfDataQualityJobDefinition_ConstraintsResourceProperty {
	var returns *TfDataQualityJobDefinition_ConstraintsResourceProperty
	_jsii_.Get(
		j,
		"constraintsResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) InternalValue() *TfDataQualityJobDefinition_DataQualityBaselineConfigProperty {
	var returns *TfDataQualityJobDefinition_DataQualityBaselineConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) StatisticsResource() TfDataQualityJobDefinition_StatisticsResourcePropertyOutputReference {
	var returns TfDataQualityJobDefinition_StatisticsResourcePropertyOutputReference
	_jsii_.Get(
		j,
		"statisticsResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) StatisticsResourceInput() *TfDataQualityJobDefinition_StatisticsResourceProperty {
	var returns *TfDataQualityJobDefinition_StatisticsResourceProperty
	_jsii_.Get(
		j,
		"statisticsResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDataQualityJobDefinition.DataQualityBaselineConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference_Override(t TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDataQualityJobDefinition.DataQualityBaselineConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference)SetInternalValue(val *TfDataQualityJobDefinition_DataQualityBaselineConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) PutConstraintsResource(value *TfDataQualityJobDefinition_ConstraintsResourceProperty) {
	if err := t.validatePutConstraintsResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConstraintsResource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) PutStatisticsResource(value *TfDataQualityJobDefinition_StatisticsResourceProperty) {
	if err := t.validatePutStatisticsResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStatisticsResource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) ResetConstraintsResource() {
	_jsii_.InvokeVoid(
		t,
		"resetConstraintsResource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) ResetStatisticsResource() {
	_jsii_.InvokeVoid(
		t,
		"resetStatisticsResource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityBaselineConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

