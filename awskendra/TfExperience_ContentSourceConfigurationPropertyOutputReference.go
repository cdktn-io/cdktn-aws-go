package awskendra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskendra/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskendra/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfExperience_ContentSourceConfigurationPropertyOutputReference interface {
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
	DataSourceIds() *[]*string
	// Experimental.
	SetDataSourceIds(val *[]*string)
	// Experimental.
	DataSourceIdsInput() *[]*string
	// Experimental.
	DirectPutContent() interface{}
	// Experimental.
	SetDirectPutContent(val interface{})
	// Experimental.
	DirectPutContentInput() interface{}
	// Experimental.
	FaqIds() *[]*string
	// Experimental.
	SetFaqIds(val *[]*string)
	// Experimental.
	FaqIdsInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfExperience_ContentSourceConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfExperience_ContentSourceConfigurationProperty)
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
	ResetDataSourceIds()
	// Experimental.
	ResetDirectPutContent()
	// Experimental.
	ResetFaqIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfExperience_ContentSourceConfigurationPropertyOutputReference
type jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) DataSourceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataSourceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) DataSourceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataSourceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) DirectPutContent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"directPutContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) DirectPutContentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"directPutContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) FaqIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"faqIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) FaqIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"faqIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) InternalValue() *TfExperience_ContentSourceConfigurationProperty {
	var returns *TfExperience_ContentSourceConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfExperience_ContentSourceConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfExperience_ContentSourceConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfExperience_ContentSourceConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kendra.TfExperience.ContentSourceConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfExperience_ContentSourceConfigurationPropertyOutputReference_Override(t TfExperience_ContentSourceConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kendra.TfExperience.ContentSourceConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference)SetDataSourceIds(val *[]*string) {
	if err := j.validateSetDataSourceIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceIds",
		val,
	)
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference)SetDirectPutContent(val interface{}) {
	if err := j.validateSetDirectPutContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directPutContent",
		val,
	)
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference)SetFaqIds(val *[]*string) {
	if err := j.validateSetFaqIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"faqIds",
		val,
	)
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference)SetInternalValue(val *TfExperience_ContentSourceConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) ResetDataSourceIds() {
	_jsii_.InvokeVoid(
		t,
		"resetDataSourceIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) ResetDirectPutContent() {
	_jsii_.InvokeVoid(
		t,
		"resetDirectPutContent",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) ResetFaqIds() {
	_jsii_.InvokeVoid(
		t,
		"resetFaqIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfExperience_ContentSourceConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

