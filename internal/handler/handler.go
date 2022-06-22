package handler

import (
	"io/ioutil"
	"log"

	admission "k8s.io/api/admission/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"

	"github.com/casbin/k8s-gatekeeper/internal/model"
	"github.com/gin-gonic/gin"
)

//Main Handler
func Handler(c *gin.Context) {

	data, _ := ioutil.ReadAll(c.Request.Body)
	var admissionReview admission.AdmissionReview
	var decoder runtime.Decoder = serializer.NewCodecFactory(runtime.NewScheme()).UniversalDeserializer()
	decoder.Decode(data, nil, &admissionReview)

	if admissionReview.Request.Resource.Resource == "casbinmodels" || admissionReview.Request.Resource.Resource == "casbinpolicies" {
		approveResponse(c, string(admissionReview.Request.UID))
		return
	}

	//for debug only. Todo:remove this block of code
	if admissionReview.Request.Namespace != "default" {
		approveResponse(c, string(admissionReview.Request.UID))
		return
	}

	//currently we are going to handle these resources:
	uid := admissionReview.Request.UID
	resource := admissionReview.Request.Resource.Resource

	switch resource {
	case "deployments":
		model.MountDeploymentObject(&admissionReview)
	case "pods":
		model.MountPodObject(&admissionReview)
	case "services":
		model.MountServiceObject(&admissionReview)
	case "ingresses":
		model.MountIngressObject(&admissionReview)
	}

	err := model.EnforcerList.Enforce(&admissionReview)
	if err != nil {
		log.Printf("%s  rejected\n", admissionReview.Request.Resource.String())
		rejectResponse(c, string(uid), err.Error())
		return
	}

	log.Printf("%s  approved\n", admissionReview.Request.Resource.String())
	approveResponse(c, string(uid))

}

func rejectResponse(c *gin.Context, uid string, rejectReason string) {
	c.JSON(200, gin.H{
		"apiVersion": "admission.k8s.io/v1",
		"kind":       "AdmissionReview",
		"response": map[string]interface{}{
			"uid":     uid,
			"allowed": false,
			"status": map[string]interface{}{
				"code":    403,
				"message": rejectReason,
			},
		},
	})
}

func approveResponse(c *gin.Context, uid string) {
	c.JSON(200, gin.H{
		"apiVersion": "admission.k8s.io/v1",
		"kind":       "AdmissionReview",
		"response": map[string]interface{}{
			"uid":     uid,
			"allowed": true,
		},
	})
}
