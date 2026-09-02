package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference interface {
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
	InstanceUrl() *string
	// Experimental.
	SetInstanceUrl(val *string)
	// Experimental.
	InstanceUrlInput() *string
	// Experimental.
	InternalValue() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty
	// Experimental.
	SetInternalValue(val *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty)
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

// The jsii proxy struct for TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference
type jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) InstanceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) InstanceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) InternalValue() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.TfConnectorProfile.ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference_Override(t TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.TfConnectorProfile.ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference)SetInstanceUrl(val *string) {
	if err := j.validateSetInstanceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceUrl",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference)SetInternalValue(val *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

