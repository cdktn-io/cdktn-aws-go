//go:build no_runtime_type_checking

package awscloudformation

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TfStackSetInstance_StackInstanceSummariesPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TfStackSetInstance_StackInstanceSummariesPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TfStackSetInstance_StackInstanceSummariesPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TfStackSetInstance_StackInstanceSummariesPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfStackSetInstance_StackInstanceSummariesPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TfStackSetInstance_StackInstanceSummariesPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTfStackSetInstance_StackInstanceSummariesPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

