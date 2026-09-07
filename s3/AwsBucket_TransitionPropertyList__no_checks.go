//go:build no_runtime_type_checking

package s3

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsBucket_TransitionPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsBucket_TransitionPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsBucket_TransitionPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsBucket_TransitionPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsBucket_TransitionPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsBucket_TransitionPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsBucket_TransitionPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsBucket_TransitionPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

