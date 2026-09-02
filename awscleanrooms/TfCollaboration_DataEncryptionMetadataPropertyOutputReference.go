package awscleanrooms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscleanrooms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscleanrooms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCollaboration_DataEncryptionMetadataPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllowClearText() interface{}
	// Experimental.
	SetAllowClearText(val interface{})
	// Experimental.
	AllowClearTextInput() interface{}
	// Experimental.
	AllowDuplicates() interface{}
	// Experimental.
	SetAllowDuplicates(val interface{})
	// Experimental.
	AllowDuplicatesInput() interface{}
	// Experimental.
	AllowJoinsOnColumnsWithDifferentNames() interface{}
	// Experimental.
	SetAllowJoinsOnColumnsWithDifferentNames(val interface{})
	// Experimental.
	AllowJoinsOnColumnsWithDifferentNamesInput() interface{}
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
	InternalValue() *TfCollaboration_DataEncryptionMetadataProperty
	// Experimental.
	SetInternalValue(val *TfCollaboration_DataEncryptionMetadataProperty)
	// Experimental.
	PreserveNulls() interface{}
	// Experimental.
	SetPreserveNulls(val interface{})
	// Experimental.
	PreserveNullsInput() interface{}
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCollaboration_DataEncryptionMetadataPropertyOutputReference
type jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) AllowClearText() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClearText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) AllowClearTextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClearTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) AllowDuplicates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowDuplicates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) AllowDuplicatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowDuplicatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) AllowJoinsOnColumnsWithDifferentNames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowJoinsOnColumnsWithDifferentNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) AllowJoinsOnColumnsWithDifferentNamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowJoinsOnColumnsWithDifferentNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) InternalValue() *TfCollaboration_DataEncryptionMetadataProperty {
	var returns *TfCollaboration_DataEncryptionMetadataProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) PreserveNulls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveNulls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) PreserveNullsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveNullsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCollaboration_DataEncryptionMetadataPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCollaboration_DataEncryptionMetadataPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCollaboration_DataEncryptionMetadataPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-clean-rooms.TfCollaboration.DataEncryptionMetadataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCollaboration_DataEncryptionMetadataPropertyOutputReference_Override(t TfCollaboration_DataEncryptionMetadataPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-clean-rooms.TfCollaboration.DataEncryptionMetadataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference)SetAllowClearText(val interface{}) {
	if err := j.validateSetAllowClearTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowClearText",
		val,
	)
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference)SetAllowDuplicates(val interface{}) {
	if err := j.validateSetAllowDuplicatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowDuplicates",
		val,
	)
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference)SetAllowJoinsOnColumnsWithDifferentNames(val interface{}) {
	if err := j.validateSetAllowJoinsOnColumnsWithDifferentNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowJoinsOnColumnsWithDifferentNames",
		val,
	)
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference)SetInternalValue(val *TfCollaboration_DataEncryptionMetadataProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference)SetPreserveNulls(val interface{}) {
	if err := j.validateSetPreserveNullsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveNulls",
		val,
	)
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCollaboration_DataEncryptionMetadataPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

