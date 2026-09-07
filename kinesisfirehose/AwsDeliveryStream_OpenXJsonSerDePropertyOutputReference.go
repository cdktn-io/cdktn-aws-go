package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/kinesisfirehose/jsii"

	"github.com/cdktn-io/cdktn-aws-go/kinesisfirehose/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CaseInsensitive() interface{}
	// Experimental.
	SetCaseInsensitive(val interface{})
	// Experimental.
	CaseInsensitiveInput() interface{}
	// Experimental.
	ColumnToJsonKeyMappings() *map[string]*string
	// Experimental.
	SetColumnToJsonKeyMappings(val *map[string]*string)
	// Experimental.
	ColumnToJsonKeyMappingsInput() *map[string]*string
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
	ConvertDotsInJsonKeysToUnderscores() interface{}
	// Experimental.
	SetConvertDotsInJsonKeysToUnderscores(val interface{})
	// Experimental.
	ConvertDotsInJsonKeysToUnderscoresInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsDeliveryStream_OpenXJsonSerDeProperty
	// Experimental.
	SetInternalValue(val *AwsDeliveryStream_OpenXJsonSerDeProperty)
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
	ResetCaseInsensitive()
	// Experimental.
	ResetColumnToJsonKeyMappings()
	// Experimental.
	ResetConvertDotsInJsonKeysToUnderscores()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference
type jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) CaseInsensitive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseInsensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) CaseInsensitiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseInsensitiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) ColumnToJsonKeyMappings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"columnToJsonKeyMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) ColumnToJsonKeyMappingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"columnToJsonKeyMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) ConvertDotsInJsonKeysToUnderscores() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"convertDotsInJsonKeysToUnderscores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) ConvertDotsInJsonKeysToUnderscoresInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"convertDotsInJsonKeysToUnderscoresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) InternalValue() *AwsDeliveryStream_OpenXJsonSerDeProperty {
	var returns *AwsDeliveryStream_OpenXJsonSerDeProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDeliveryStream_OpenXJsonSerDePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDeliveryStream_OpenXJsonSerDePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsDeliveryStream.OpenXJsonSerDePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDeliveryStream_OpenXJsonSerDePropertyOutputReference_Override(a AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsDeliveryStream.OpenXJsonSerDePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference)SetCaseInsensitive(val interface{}) {
	if err := j.validateSetCaseInsensitiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caseInsensitive",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference)SetColumnToJsonKeyMappings(val *map[string]*string) {
	if err := j.validateSetColumnToJsonKeyMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnToJsonKeyMappings",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference)SetConvertDotsInJsonKeysToUnderscores(val interface{}) {
	if err := j.validateSetConvertDotsInJsonKeysToUnderscoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"convertDotsInJsonKeysToUnderscores",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference)SetInternalValue(val *AwsDeliveryStream_OpenXJsonSerDeProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) ResetCaseInsensitive() {
	_jsii_.InvokeVoid(
		a,
		"resetCaseInsensitive",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) ResetColumnToJsonKeyMappings() {
	_jsii_.InvokeVoid(
		a,
		"resetColumnToJsonKeyMappings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) ResetConvertDotsInJsonKeysToUnderscores() {
	_jsii_.InvokeVoid(
		a,
		"resetConvertDotsInJsonKeysToUnderscores",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDeliveryStream_OpenXJsonSerDePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

