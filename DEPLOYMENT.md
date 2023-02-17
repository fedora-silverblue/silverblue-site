# Website deployment in Fedora Infra

The website is deployed as an OpenShift application in Fedora's OpenShift
cluster. The configuration itself is stored in the
[fedora/ansible](https://pagure.io/fedora-infra/ansible) repository in
[playbooks/openshift-apps/silverblue.yml](https://pagure.io/fedora-infra/ansible/blob/main/f/playbooks/openshift-apps/silverblue.yml)
and
[roles/openshift-apps/silverblue](https://pagure.io/fedora-infra/ansible/blob/main/f/roles/openshift-apps/silverblue).

The website is updated after each new commit via a webhook that triggers a
source to image build in the OpenShift cluster. See the
[ImageStream](https://pagure.io/fedora-infra/ansible/blob/main/f/roles/openshift-apps/silverblue/templates/imagestream.yml).

To get access to the project in the OpenShift console, you should add your name
in the list of application owners in
[silverblue.yaml](https://pagure.io/fedora-infra/ansible/blob/main/f/playbooks/openshift-apps/silverblue.yml#_15).

You will then be able to access the
[silverblue project](https://console-openshift-console.apps.ocp.fedoraproject.org/k8s/cluster/projects/silverblue).
Look for the
[BuildConfig](https://pagure.io/fedora-infra/ansible/blob/main/f/roles/openshift-apps/silverblue/templates/buildconfig.yml)
to update the Webhook.
