kubectl apply -f config/auth.casbin.org_casbinmodels.yaml 
kubectl apply -f config/auth.casbin.org_casbinpolicies.yaml
kubectl apply -f config/webhook_external.yaml 

kubectl apply -f model.yaml  -f policy.yaml 



kubectl delete -f config/webhook_external.yaml 
kubectl delete -f config/auth.casbin.org_casbinmodels.yaml 
kubectl delete -f config/auth.casbin.org_casbinpolicies.yaml