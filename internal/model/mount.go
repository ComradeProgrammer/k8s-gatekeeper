package model

import (
	"encoding/json"

	admission "k8s.io/api/admission/v1"
	app "k8s.io/api/apps/v1"
	networking "k8s.io/api/networking/v1"
)

func MountDeploymentObject(admissionReview *admission.AdmissionReview) error {
	admissionReview.Request.Object.Object = nil
	if len(admissionReview.Request.Object.Raw) != 0 {
		var deploymentObject app.Service
		err := json.Unmarshal(admissionReview.Request.Object.Raw, &deploymentObject)
		if err != nil {
			return err
		}
		admissionReview.Request.Object.Object = &deploymentObject
	}

	admissionReview.Request.OldObject.Object = nil
	if len(admissionReview.Request.OldObject.Raw) != 0 {
		var deploymentOldObject app.Service
		err := json.Unmarshal(admissionReview.Request.OldObject.Raw, &deploymentOldObject)
		if err != nil {
			return err
		}
		admissionReview.Request.OldObject.Object = &deploymentOldObject
	}
	return nil
}

func MountPodObject(admissionReview *admission.AdmissionReview) error {
	admissionReview.Request.Object.Object = nil
	if len(admissionReview.Request.Object.Raw) != 0 {
		var podObject app.Service
		err := json.Unmarshal(admissionReview.Request.Object.Raw, &podObject)
		if err != nil {
			return err
		}
		admissionReview.Request.Object.Object = &podObject
	}

	admissionReview.Request.OldObject.Object = nil
	if len(admissionReview.Request.OldObject.Raw) != 0 {
		var podOldObject app.Service
		err := json.Unmarshal(admissionReview.Request.OldObject.Raw, &podOldObject)
		if err != nil {
			return err
		}
		admissionReview.Request.OldObject.Object = &podOldObject
	}
	return nil
}

func MountServiceObject(admissionReview *admission.AdmissionReview) error {
	admissionReview.Request.Object.Object = nil
	if len(admissionReview.Request.Object.Raw) != 0 {
		var serviceObject app.Service
		err := json.Unmarshal(admissionReview.Request.Object.Raw, &serviceObject)
		if err != nil {
			return err
		}
		admissionReview.Request.Object.Object = &serviceObject
	}

	admissionReview.Request.OldObject.Object = nil
	if len(admissionReview.Request.OldObject.Raw) != 0 {
		var serviceOldObject app.Service
		err := json.Unmarshal(admissionReview.Request.OldObject.Raw, &serviceOldObject)
		if err != nil {
			return err
		}
		admissionReview.Request.OldObject.Object = &serviceOldObject
	}
	return nil
}

func MountIngressObject(admissionReview *admission.AdmissionReview) error {
	admissionReview.Request.Object.Object = nil
	if len(admissionReview.Request.Object.Raw) != 0 {
		var ingressObject networking.Ingress
		err := json.Unmarshal(admissionReview.Request.Object.Raw, &ingressObject)
		if err != nil {
			return err
		}
		admissionReview.Request.Object.Object = &ingressObject
	}

	admissionReview.Request.OldObject.Object = nil
	if len(admissionReview.Request.OldObject.Raw) != 0 {
		var ingressOldObject networking.Ingress
		err := json.Unmarshal(admissionReview.Request.OldObject.Raw, &ingressOldObject)
		if err != nil {
			return err
		}
		admissionReview.Request.OldObject.Object = &ingressOldObject
	}
	return nil
}

