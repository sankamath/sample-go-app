# sample-go-app


oc new-project dev --display-name="Dev"
oc new-project stage --display-name="Stage"
oc new-project cicd --display-name="CI/CD"
oc new-project build --display-name="Builds"


oc new-app jenkins-ephemeral -n cicd 


kustomize build | oc apply -f- -n st