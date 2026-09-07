package appflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference interface {
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
	InternalValue() *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty
	// Experimental.
	SetInternalValue(val *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty)
	// Experimental.
	ObjectPath() *string
	// Experimental.
	SetObjectPath(val *string)
	// Experimental.
	ObjectPathInput() *string
	// Experimental.
	PaginationConfig() AwsFlow_PaginationConfigPropertyOutputReference
	// Experimental.
	PaginationConfigInput() *AwsFlow_PaginationConfigProperty
	// Experimental.
	ParallelismConfig() AwsFlow_ParallelismConfigPropertyOutputReference
	// Experimental.
	ParallelismConfigInput() *AwsFlow_ParallelismConfigProperty
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
	PutPaginationConfig(value *AwsFlow_PaginationConfigProperty)
	// Experimental.
	PutParallelismConfig(value *AwsFlow_ParallelismConfigProperty)
	// Experimental.
	ResetPaginationConfig()
	// Experimental.
	ResetParallelismConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference
type jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) InternalValue() *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty {
	var returns *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) ObjectPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) ObjectPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) PaginationConfig() AwsFlow_PaginationConfigPropertyOutputReference {
	var returns AwsFlow_PaginationConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"paginationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) PaginationConfigInput() *AwsFlow_PaginationConfigProperty {
	var returns *AwsFlow_PaginationConfigProperty
	_jsii_.Get(
		j,
		"paginationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) ParallelismConfig() AwsFlow_ParallelismConfigPropertyOutputReference {
	var returns AwsFlow_ParallelismConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"parallelismConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) ParallelismConfigInput() *AwsFlow_ParallelismConfigProperty {
	var returns *AwsFlow_ParallelismConfigProperty
	_jsii_.Get(
		j,
		"parallelismConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsFlow.SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference_Override(a AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsFlow.SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference)SetInternalValue(val *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference)SetObjectPath(val *string) {
	if err := j.validateSetObjectPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectPath",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) PutPaginationConfig(value *AwsFlow_PaginationConfigProperty) {
	if err := a.validatePutPaginationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPaginationConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) PutParallelismConfig(value *AwsFlow_ParallelismConfigProperty) {
	if err := a.validatePutParallelismConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParallelismConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) ResetPaginationConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetPaginationConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) ResetParallelismConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetParallelismConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

