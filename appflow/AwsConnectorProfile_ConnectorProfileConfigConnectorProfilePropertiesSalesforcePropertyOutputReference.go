package appflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference interface {
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
	InternalValue() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty
	// Experimental.
	SetInternalValue(val *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty)
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

// The jsii proxy struct for AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference
type jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) InstanceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) InstanceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) InternalValue() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) IsSandboxEnvironment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSandboxEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) IsSandboxEnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSandboxEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) UsePrivatelinkForMetadataAndAuthorization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePrivatelinkForMetadataAndAuthorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) UsePrivatelinkForMetadataAndAuthorizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePrivatelinkForMetadataAndAuthorizationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsConnectorProfile.ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference_Override(a AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsConnectorProfile.ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference)SetInstanceUrl(val *string) {
	if err := j.validateSetInstanceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceUrl",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference)SetInternalValue(val *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference)SetIsSandboxEnvironment(val interface{}) {
	if err := j.validateSetIsSandboxEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSandboxEnvironment",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference)SetUsePrivatelinkForMetadataAndAuthorization(val interface{}) {
	if err := j.validateSetUsePrivatelinkForMetadataAndAuthorizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usePrivatelinkForMetadataAndAuthorization",
		val,
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) ResetInstanceUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) ResetIsSandboxEnvironment() {
	_jsii_.InvokeVoid(
		a,
		"resetIsSandboxEnvironment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) ResetUsePrivatelinkForMetadataAndAuthorization() {
	_jsii_.InvokeVoid(
		a,
		"resetUsePrivatelinkForMetadataAndAuthorization",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

