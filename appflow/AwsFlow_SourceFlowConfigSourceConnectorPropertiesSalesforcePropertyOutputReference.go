package appflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference interface {
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
	DataTransferApi() *string
	// Experimental.
	SetDataTransferApi(val *string)
	// Experimental.
	DataTransferApiInput() *string
	// Experimental.
	EnableDynamicFieldUpdate() interface{}
	// Experimental.
	SetEnableDynamicFieldUpdate(val interface{})
	// Experimental.
	EnableDynamicFieldUpdateInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	IncludeDeletedRecords() interface{}
	// Experimental.
	SetIncludeDeletedRecords(val interface{})
	// Experimental.
	IncludeDeletedRecordsInput() interface{}
	// Experimental.
	InternalValue() *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty
	// Experimental.
	SetInternalValue(val *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty)
	// Experimental.
	Object() *string
	// Experimental.
	SetObject(val *string)
	// Experimental.
	ObjectInput() *string
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
	ResetDataTransferApi()
	// Experimental.
	ResetEnableDynamicFieldUpdate()
	// Experimental.
	ResetIncludeDeletedRecords()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference
type jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) DataTransferApi() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTransferApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) DataTransferApiInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTransferApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) EnableDynamicFieldUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDynamicFieldUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) EnableDynamicFieldUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDynamicFieldUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) IncludeDeletedRecords() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeDeletedRecords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) IncludeDeletedRecordsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeDeletedRecordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) InternalValue() *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty {
	var returns *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) Object() *string {
	var returns *string
	_jsii_.Get(
		j,
		"object",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) ObjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsFlow.SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference_Override(a AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsFlow.SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference)SetDataTransferApi(val *string) {
	if err := j.validateSetDataTransferApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataTransferApi",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference)SetEnableDynamicFieldUpdate(val interface{}) {
	if err := j.validateSetEnableDynamicFieldUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDynamicFieldUpdate",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference)SetIncludeDeletedRecords(val interface{}) {
	if err := j.validateSetIncludeDeletedRecordsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeDeletedRecords",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference)SetInternalValue(val *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference)SetObject(val *string) {
	if err := j.validateSetObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"object",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) ResetDataTransferApi() {
	_jsii_.InvokeVoid(
		a,
		"resetDataTransferApi",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) ResetEnableDynamicFieldUpdate() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableDynamicFieldUpdate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) ResetIncludeDeletedRecords() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeDeletedRecords",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

