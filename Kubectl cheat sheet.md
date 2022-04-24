Kubectl cheat sheet

mv ~/Downloads/civo-k8s_demo_1-kubeconfig $HOME/.kube/config
kubectl config use-context
kubectl config view
k get pods --all-namespaces
k get pods --namespace nginx1
k get pods
kubectl logs --follow nginx1-56c6c9f45f-9h6hz
kubectl exec --stdin --tty nginx1-56c6c9f45f-9h6hz -- /bin/bash

---

 k get nodes -o wide
 k config set-context --current --namespace=ingress-nginx
 k describe pod ingress-nginx-controller-58fdb7bb4-vmch4
 k create --filename <yaml>
 k get ns
 k get all -n <namespace>

 apply vs ...
 k create .
 k expose deploy <pod> --port 80

 k get ing(ress)
 k describe ing <ing service>

 k port-forward devops-toolkit-599895bb6b-mbks8 8080
 


