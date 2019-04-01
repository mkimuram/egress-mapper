#! /bin/bash

yamls() {
cat << EOF  
service_account.yaml
role.yaml
role_binding.yaml
crds/egress_v1alpha1_egressmapper_crd.yaml 
crds/egress_v1alpha1_egressmapper_cr.yaml
crds/egress_v1alpha1_egress_crd.yaml
crds/egress_v1alpha1_vip_crd.yaml 
keepalived/service_account.yaml
keepalived/clusterrole.yaml
keepalived/clusterrolebinding.yaml
operator.yaml
EOF
}

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

do_create() {
        yamls | while read yaml;do 
                kubectl create -f ${DIR}/${yaml}
				sleep 1
        done
}

do_delete() {
        yamls | tac | while read yaml;do 
                kubectl delete -f ${DIR}/${yaml}
        done
}

if [ $# -eq 0 ];then
        do_create
elif [ x"$1" == x"-d" ];then
        do_delete
else
        echo "Usage: $0 [-d]"
fi
