package awsemrcontainers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsemrcontainers/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsemrcontainers/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference interface {
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
	EntryPoint() *string
	// Experimental.
	SetEntryPoint(val *string)
	// Experimental.
	EntryPointArguments() *[]*string
	// Experimental.
	SetEntryPointArguments(val *[]*string)
	// Experimental.
	EntryPointArgumentsInput() *[]*string
	// Experimental.
	EntryPointInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfJobTemplate_SparkSubmitJobDriverProperty
	// Experimental.
	SetInternalValue(val *TfJobTemplate_SparkSubmitJobDriverProperty)
	// Experimental.
	SparkSubmitParameters() *string
	// Experimental.
	SetSparkSubmitParameters(val *string)
	// Experimental.
	SparkSubmitParametersInput() *string
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
	ResetEntryPointArguments()
	// Experimental.
	ResetSparkSubmitParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference
type jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) EntryPoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) EntryPointArguments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryPointArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) EntryPointArgumentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryPointArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) EntryPointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) InternalValue() *TfJobTemplate_SparkSubmitJobDriverProperty {
	var returns *TfJobTemplate_SparkSubmitJobDriverProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) SparkSubmitParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkSubmitParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) SparkSubmitParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkSubmitParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfJobTemplate_SparkSubmitJobDriverPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfJobTemplate_SparkSubmitJobDriverPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr-containers.TfJobTemplate.SparkSubmitJobDriverPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfJobTemplate_SparkSubmitJobDriverPropertyOutputReference_Override(t TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr-containers.TfJobTemplate.SparkSubmitJobDriverPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference)SetEntryPoint(val *string) {
	if err := j.validateSetEntryPointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryPoint",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference)SetEntryPointArguments(val *[]*string) {
	if err := j.validateSetEntryPointArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryPointArguments",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference)SetInternalValue(val *TfJobTemplate_SparkSubmitJobDriverProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference)SetSparkSubmitParameters(val *string) {
	if err := j.validateSetSparkSubmitParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkSubmitParameters",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) ResetEntryPointArguments() {
	_jsii_.InvokeVoid(
		t,
		"resetEntryPointArguments",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) ResetSparkSubmitParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetSparkSubmitParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfJobTemplate_SparkSubmitJobDriverPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

