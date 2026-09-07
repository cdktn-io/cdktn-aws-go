package cleanrooms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cleanrooms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/cleanrooms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCollaboration_DataEncryptionMetadataPropertyOutputReference interface {
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
	InternalValue() *AwsCollaboration_DataEncryptionMetadataProperty
	// Experimental.
	SetInternalValue(val *AwsCollaboration_DataEncryptionMetadataProperty)
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

// The jsii proxy struct for AwsCollaboration_DataEncryptionMetadataPropertyOutputReference
type jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) AllowClearText() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClearText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) AllowClearTextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClearTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) AllowDuplicates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowDuplicates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) AllowDuplicatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowDuplicatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) AllowJoinsOnColumnsWithDifferentNames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowJoinsOnColumnsWithDifferentNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) AllowJoinsOnColumnsWithDifferentNamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowJoinsOnColumnsWithDifferentNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) InternalValue() *AwsCollaboration_DataEncryptionMetadataProperty {
	var returns *AwsCollaboration_DataEncryptionMetadataProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) PreserveNulls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveNulls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) PreserveNullsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveNullsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCollaboration_DataEncryptionMetadataPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCollaboration_DataEncryptionMetadataPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCollaboration_DataEncryptionMetadataPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-clean-rooms.AwsCollaboration.DataEncryptionMetadataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCollaboration_DataEncryptionMetadataPropertyOutputReference_Override(a AwsCollaboration_DataEncryptionMetadataPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-clean-rooms.AwsCollaboration.DataEncryptionMetadataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference)SetAllowClearText(val interface{}) {
	if err := j.validateSetAllowClearTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowClearText",
		val,
	)
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference)SetAllowDuplicates(val interface{}) {
	if err := j.validateSetAllowDuplicatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowDuplicates",
		val,
	)
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference)SetAllowJoinsOnColumnsWithDifferentNames(val interface{}) {
	if err := j.validateSetAllowJoinsOnColumnsWithDifferentNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowJoinsOnColumnsWithDifferentNames",
		val,
	)
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference)SetInternalValue(val *AwsCollaboration_DataEncryptionMetadataProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference)SetPreserveNulls(val interface{}) {
	if err := j.validateSetPreserveNullsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveNulls",
		val,
	)
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCollaboration_DataEncryptionMetadataPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

