package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdatasync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdatasync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDatasyncTask_TaskReportConfigPropertyOutputReference interface {
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
	InternalValue() *AwsDatasyncTask_TaskReportConfigProperty
	// Experimental.
	SetInternalValue(val *AwsDatasyncTask_TaskReportConfigProperty)
	// Experimental.
	OutputType() *string
	// Experimental.
	SetOutputType(val *string)
	// Experimental.
	OutputTypeInput() *string
	// Experimental.
	ReportLevel() *string
	// Experimental.
	SetReportLevel(val *string)
	// Experimental.
	ReportLevelInput() *string
	// Experimental.
	ReportOverrides() AwsDatasyncTask_ReportOverridesPropertyOutputReference
	// Experimental.
	ReportOverridesInput() *AwsDatasyncTask_ReportOverridesProperty
	// Experimental.
	S3Destination() AwsDatasyncTask_S3DestinationPropertyOutputReference
	// Experimental.
	S3DestinationInput() *AwsDatasyncTask_S3DestinationProperty
	// Experimental.
	S3ObjectVersioning() *string
	// Experimental.
	SetS3ObjectVersioning(val *string)
	// Experimental.
	S3ObjectVersioningInput() *string
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
	PutReportOverrides(value *AwsDatasyncTask_ReportOverridesProperty)
	// Experimental.
	PutS3Destination(value *AwsDatasyncTask_S3DestinationProperty)
	// Experimental.
	ResetOutputType()
	// Experimental.
	ResetReportLevel()
	// Experimental.
	ResetReportOverrides()
	// Experimental.
	ResetS3ObjectVersioning()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDatasyncTask_TaskReportConfigPropertyOutputReference
type jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) InternalValue() *AwsDatasyncTask_TaskReportConfigProperty {
	var returns *AwsDatasyncTask_TaskReportConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) OutputType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) OutputTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) ReportLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) ReportLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) ReportOverrides() AwsDatasyncTask_ReportOverridesPropertyOutputReference {
	var returns AwsDatasyncTask_ReportOverridesPropertyOutputReference
	_jsii_.Get(
		j,
		"reportOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) ReportOverridesInput() *AwsDatasyncTask_ReportOverridesProperty {
	var returns *AwsDatasyncTask_ReportOverridesProperty
	_jsii_.Get(
		j,
		"reportOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) S3Destination() AwsDatasyncTask_S3DestinationPropertyOutputReference {
	var returns AwsDatasyncTask_S3DestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) S3DestinationInput() *AwsDatasyncTask_S3DestinationProperty {
	var returns *AwsDatasyncTask_S3DestinationProperty
	_jsii_.Get(
		j,
		"s3DestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) S3ObjectVersioning() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectVersioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) S3ObjectVersioningInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectVersioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDatasyncTask_TaskReportConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDatasyncTask_TaskReportConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDatasyncTask_TaskReportConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-datasync.AwsDatasyncTask.TaskReportConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDatasyncTask_TaskReportConfigPropertyOutputReference_Override(a AwsDatasyncTask_TaskReportConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-datasync.AwsDatasyncTask.TaskReportConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference)SetInternalValue(val *AwsDatasyncTask_TaskReportConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference)SetOutputType(val *string) {
	if err := j.validateSetOutputTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputType",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference)SetReportLevel(val *string) {
	if err := j.validateSetReportLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportLevel",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference)SetS3ObjectVersioning(val *string) {
	if err := j.validateSetS3ObjectVersioningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3ObjectVersioning",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) PutReportOverrides(value *AwsDatasyncTask_ReportOverridesProperty) {
	if err := a.validatePutReportOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putReportOverrides",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) PutS3Destination(value *AwsDatasyncTask_S3DestinationProperty) {
	if err := a.validatePutS3DestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3Destination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) ResetOutputType() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) ResetReportLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetReportLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) ResetReportOverrides() {
	_jsii_.InvokeVoid(
		a,
		"resetReportOverrides",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) ResetS3ObjectVersioning() {
	_jsii_.InvokeVoid(
		a,
		"resetS3ObjectVersioning",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDatasyncTask_TaskReportConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

