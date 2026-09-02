package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference interface {
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
	InternalValue() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty
	// Experimental.
	SetInternalValue(val *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty)
	// Experimental.
	IsSandboxEnvironment() interface{}
	// Experimental.
	SetIsSandboxEnvironment(val interface{})
	// Experimental.
	IsSandboxEnvironmentInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UsePrivatelinkForMetadataAndAuthorization() interface{}
	// Experimental.
	SetUsePrivatelinkForMetadataAndAuthorization(val interface{})
	// Experimental.
	UsePrivatelinkForMetadataAndAuthorizationInput() interface{}
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
	ResetInstanceUrl()
	// Experimental.
	ResetIsSandboxEnvironment()
	// Experimental.
	ResetUsePrivatelinkForMetadataAndAuthorization()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference
type jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) InstanceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) InstanceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) InternalValue() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) IsSandboxEnvironment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSandboxEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) IsSandboxEnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSandboxEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) UsePrivatelinkForMetadataAndAuthorization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePrivatelinkForMetadataAndAuthorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) UsePrivatelinkForMetadataAndAuthorizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePrivatelinkForMetadataAndAuthorizationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.TfConnectorProfile.ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference_Override(t TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.TfConnectorProfile.ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference)SetInstanceUrl(val *string) {
	if err := j.validateSetInstanceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceUrl",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference)SetInternalValue(val *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference)SetIsSandboxEnvironment(val interface{}) {
	if err := j.validateSetIsSandboxEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSandboxEnvironment",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference)SetUsePrivatelinkForMetadataAndAuthorization(val interface{}) {
	if err := j.validateSetUsePrivatelinkForMetadataAndAuthorizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usePrivatelinkForMetadataAndAuthorization",
		val,
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) ResetInstanceUrl() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceUrl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) ResetIsSandboxEnvironment() {
	_jsii_.InvokeVoid(
		t,
		"resetIsSandboxEnvironment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) ResetUsePrivatelinkForMetadataAndAuthorization() {
	_jsii_.InvokeVoid(
		t,
		"resetUsePrivatelinkForMetadataAndAuthorization",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

