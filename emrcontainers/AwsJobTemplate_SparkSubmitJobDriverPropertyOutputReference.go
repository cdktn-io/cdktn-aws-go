package emrcontainers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/emrcontainers/jsii"

	"github.com/cdktn-io/cdktn-aws-go/emrcontainers/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference interface {
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
	InternalValue() *AwsJobTemplate_SparkSubmitJobDriverProperty
	// Experimental.
	SetInternalValue(val *AwsJobTemplate_SparkSubmitJobDriverProperty)
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

// The jsii proxy struct for AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference
type jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) EntryPoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) EntryPointArguments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryPointArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) EntryPointArgumentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryPointArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) EntryPointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) InternalValue() *AwsJobTemplate_SparkSubmitJobDriverProperty {
	var returns *AwsJobTemplate_SparkSubmitJobDriverProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) SparkSubmitParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkSubmitParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) SparkSubmitParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkSubmitParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsJobTemplate_SparkSubmitJobDriverPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr-containers.AwsJobTemplate.SparkSubmitJobDriverPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference_Override(a AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr-containers.AwsJobTemplate.SparkSubmitJobDriverPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference)SetEntryPoint(val *string) {
	if err := j.validateSetEntryPointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryPoint",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference)SetEntryPointArguments(val *[]*string) {
	if err := j.validateSetEntryPointArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryPointArguments",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference)SetInternalValue(val *AwsJobTemplate_SparkSubmitJobDriverProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference)SetSparkSubmitParameters(val *string) {
	if err := j.validateSetSparkSubmitParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkSubmitParameters",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) ResetEntryPointArguments() {
	_jsii_.InvokeVoid(
		a,
		"resetEntryPointArguments",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) ResetSparkSubmitParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetSparkSubmitParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsJobTemplate_SparkSubmitJobDriverPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

