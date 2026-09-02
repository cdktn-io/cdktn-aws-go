package awsemrcontainers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsemrcontainers/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsemrcontainers/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfJobTemplate_JobDriverPropertyOutputReference interface {
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
	InternalValue() *TfJobTemplate_JobDriverProperty
	// Experimental.
	SetInternalValue(val *TfJobTemplate_JobDriverProperty)
	// Experimental.
	SparkSqlJobDriver() TfJobTemplate_SparkSqlJobDriverPropertyOutputReference
	// Experimental.
	SparkSqlJobDriverInput() *TfJobTemplate_SparkSqlJobDriverProperty
	// Experimental.
	SparkSubmitJobDriver() TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference
	// Experimental.
	SparkSubmitJobDriverInput() *TfJobTemplate_SparkSubmitJobDriverProperty
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
	PutSparkSqlJobDriver(value *TfJobTemplate_SparkSqlJobDriverProperty)
	// Experimental.
	PutSparkSubmitJobDriver(value *TfJobTemplate_SparkSubmitJobDriverProperty)
	// Experimental.
	ResetSparkSqlJobDriver()
	// Experimental.
	ResetSparkSubmitJobDriver()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfJobTemplate_JobDriverPropertyOutputReference
type jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) InternalValue() *TfJobTemplate_JobDriverProperty {
	var returns *TfJobTemplate_JobDriverProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) SparkSqlJobDriver() TfJobTemplate_SparkSqlJobDriverPropertyOutputReference {
	var returns TfJobTemplate_SparkSqlJobDriverPropertyOutputReference
	_jsii_.Get(
		j,
		"sparkSqlJobDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) SparkSqlJobDriverInput() *TfJobTemplate_SparkSqlJobDriverProperty {
	var returns *TfJobTemplate_SparkSqlJobDriverProperty
	_jsii_.Get(
		j,
		"sparkSqlJobDriverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) SparkSubmitJobDriver() TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference {
	var returns TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference
	_jsii_.Get(
		j,
		"sparkSubmitJobDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) SparkSubmitJobDriverInput() *TfJobTemplate_SparkSubmitJobDriverProperty {
	var returns *TfJobTemplate_SparkSubmitJobDriverProperty
	_jsii_.Get(
		j,
		"sparkSubmitJobDriverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfJobTemplate_JobDriverPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfJobTemplate_JobDriverPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfJobTemplate_JobDriverPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr-containers.TfJobTemplate.JobDriverPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfJobTemplate_JobDriverPropertyOutputReference_Override(t TfJobTemplate_JobDriverPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr-containers.TfJobTemplate.JobDriverPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference)SetInternalValue(val *TfJobTemplate_JobDriverProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) PutSparkSqlJobDriver(value *TfJobTemplate_SparkSqlJobDriverProperty) {
	if err := t.validatePutSparkSqlJobDriverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSparkSqlJobDriver",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) PutSparkSubmitJobDriver(value *TfJobTemplate_SparkSubmitJobDriverProperty) {
	if err := t.validatePutSparkSubmitJobDriverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSparkSubmitJobDriver",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) ResetSparkSqlJobDriver() {
	_jsii_.InvokeVoid(
		t,
		"resetSparkSqlJobDriver",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) ResetSparkSubmitJobDriver() {
	_jsii_.InvokeVoid(
		t,
		"resetSparkSubmitJobDriver",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfJobTemplate_JobDriverPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

