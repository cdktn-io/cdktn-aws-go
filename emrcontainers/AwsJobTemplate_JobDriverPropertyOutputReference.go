package emrcontainers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/emrcontainers/jsii"

	"github.com/cdktn-io/cdktn-aws-go/emrcontainers/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsJobTemplate_JobDriverPropertyOutputReference interface {
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
	InternalValue() *AwsJobTemplate_JobDriverProperty
	// Experimental.
	SetInternalValue(val *AwsJobTemplate_JobDriverProperty)
	// Experimental.
	SparkSqlJobDriver() AwsJobTemplate_SparkSqlJobDriverPropertyOutputReference
	// Experimental.
	SparkSqlJobDriverInput() *AwsJobTemplate_SparkSqlJobDriverProperty
	// Experimental.
	SparkSubmitJobDriver() AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference
	// Experimental.
	SparkSubmitJobDriverInput() *AwsJobTemplate_SparkSubmitJobDriverProperty
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
	PutSparkSqlJobDriver(value *AwsJobTemplate_SparkSqlJobDriverProperty)
	// Experimental.
	PutSparkSubmitJobDriver(value *AwsJobTemplate_SparkSubmitJobDriverProperty)
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

// The jsii proxy struct for AwsJobTemplate_JobDriverPropertyOutputReference
type jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) InternalValue() *AwsJobTemplate_JobDriverProperty {
	var returns *AwsJobTemplate_JobDriverProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) SparkSqlJobDriver() AwsJobTemplate_SparkSqlJobDriverPropertyOutputReference {
	var returns AwsJobTemplate_SparkSqlJobDriverPropertyOutputReference
	_jsii_.Get(
		j,
		"sparkSqlJobDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) SparkSqlJobDriverInput() *AwsJobTemplate_SparkSqlJobDriverProperty {
	var returns *AwsJobTemplate_SparkSqlJobDriverProperty
	_jsii_.Get(
		j,
		"sparkSqlJobDriverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) SparkSubmitJobDriver() AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference {
	var returns AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference
	_jsii_.Get(
		j,
		"sparkSubmitJobDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) SparkSubmitJobDriverInput() *AwsJobTemplate_SparkSubmitJobDriverProperty {
	var returns *AwsJobTemplate_SparkSubmitJobDriverProperty
	_jsii_.Get(
		j,
		"sparkSubmitJobDriverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsJobTemplate_JobDriverPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsJobTemplate_JobDriverPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsJobTemplate_JobDriverPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr-containers.AwsJobTemplate.JobDriverPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsJobTemplate_JobDriverPropertyOutputReference_Override(a AwsJobTemplate_JobDriverPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr-containers.AwsJobTemplate.JobDriverPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference)SetInternalValue(val *AwsJobTemplate_JobDriverProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) PutSparkSqlJobDriver(value *AwsJobTemplate_SparkSqlJobDriverProperty) {
	if err := a.validatePutSparkSqlJobDriverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSparkSqlJobDriver",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) PutSparkSubmitJobDriver(value *AwsJobTemplate_SparkSubmitJobDriverProperty) {
	if err := a.validatePutSparkSubmitJobDriverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSparkSubmitJobDriver",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) ResetSparkSqlJobDriver() {
	_jsii_.InvokeVoid(
		a,
		"resetSparkSqlJobDriver",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) ResetSparkSubmitJobDriver() {
	_jsii_.InvokeVoid(
		a,
		"resetSparkSubmitJobDriver",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsJobTemplate_JobDriverPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

