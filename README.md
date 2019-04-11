# Egress-mapper

Egress-mapper is a kubernetes operator that manages [`keepalived-vip`](https://github.com/kubernetes/contrib/tree/master/keepalived-vip) and [`kube-egress`](https://github.com/steven-sheehy/kube-egress), and keeps mappings of pod ip and VIPs for `kube-egress` up-to-date. This project is still under development, so never use this version in production.

## Usage

1. Install [`operator-sdk`](https://github.com/operator-framework/operator-sdk)

    See [operator-sdk document](https://github.com/operator-framework/operator-sdk#quick-start)

2. Build egress-mapper container image

    1. Run below commands

        ```console
        $ git clone https://github.com/mkimuram/egress-mapper.git
        $ cd egress-mapper
        $ operator-sdk generate k8s
        $ operator-sdk build mkimuram/egress-mapper:latest
        $ docker push mkimuram/egress-mapper
        ```

3. Deploy egress-mapper

    1. Review and edit [`deploy/crds/egress_v1alpha1_egressmapper_cr.yaml`](https://github.com/mkimuram/egress-mapper/blob/master/deploy/crds/egress_v1alpha1_egressmapper_cr.yaml)

    2. Run deploy command

        ```console
        $ deploy/deploy.sh
        serviceaccount/egress-mapper created
        role.rbac.authorization.k8s.io/egress-mapper created
        rolebinding.rbac.authorization.k8s.io/egress-mapper created
        customresourcedefinition.apiextensions.k8s.io/egressmappers.egress.mapper.com created
        egressmapper.egress.mapper.com/example-egressmapper created
        customresourcedefinition.apiextensions.k8s.io/egresses.egress.mapper.com created
        customresourcedefinition.apiextensions.k8s.io/vips.egress.mapper.com created
        serviceaccount/kube-keepalived-vip created
        clusterrole.rbac.authorization.k8s.io/kube-keepalived-vip created
        clusterrolebinding.rbac.authorization.k8s.io/kube-keepalived-vip created
        deployment.apps/egress-mapper created
        ```

    3. Confirm that `egressmapper` deployment, `keepalived-vip` daemonset, and `kube-egress` daemonset are created (`keepalived-vip` pod may become `Completed` state due to empty VIP, but never mind. This will go back to `Running` state later.)

        ```console
        $ kubectl get deployment
        NAME            READY   UP-TO-DATE   AVAILABLE   AGE
        egress-mapper   1/1     1            1           6s
        $ kubectl get ds
        NAME                  DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE
        kube-egress           2         2         2       2            2           <none>          6s
        kube-keepalived-vip   2         2         2       2            2           <none>          6s
        $ kubectl get pod
        NAME                             READY   STATUS      RESTARTS   AGE
        egress-mapper-7c65b79f7d-sz2ts   1/1     Running     0          16s
        kube-egress-6qcgj                1/1     Running     0          11s
        kube-egress-glmhx                1/1     Running     0          11s
        kube-keepalived-vip-bwdcw        1/1     Running     0          11s
        kube-keepalived-vip-vms4w        0/1     Completed   0          11s
        ```

4. Create vip cr

    1. Review and edit [`deploy/crds/egress_v1alpha1_vip_cr.yaml`](https://github.com/mkimuram/egress-mapper/blob/master/deploy/crds/egress_v1alpha1_vip_cr.yaml)

    2. Run below command to create vip cr

        ```console
        $ kubectl create -f deploy/crds/egress_v1alpha1_vip_cr.yaml
        vip.egress.mapper.com/example-vip created
        ```

    3. Confirm that vip cr is created

        ```console
        $ kubectl get vip -o=custom-columns=NAME:.metadata.name,IP:.spec.ip,ROUTEID:.spec.routeid
        NAME          IP                ROUTEID
        example-vip   192.168.122.222   64
        ```

5. Create egress cr

    1. Review and edit [`deploy/crds/egress_v1alpha1_egress_cr.yaml`](https://github.com/mkimuram/egress-mapper/blob/master/deploy/crds/egress_v1alpha1_egress_cr.yaml)

    2. Run below command to create egress cr

        ```console
        $ kubectl create -f deploy/crds/egress_v1alpha1_egress_cr.yaml
        egress.egress.mapper.com/example-pod1-egress created
        ```

    3. Confirm that egress cr is created

        ```console
        $ kubectl get egress -o=custom-columns=NAME:.metadata.name,IP:.spec.ip,KIND:.spec.kind,NAMESPACE:.spec.namespace,RESOURCENAME:.spec.name
        NAME                  IP                KIND   NAMESPACE   RESOURCENAME
        example-pod1-egress   192.168.122.222   pod    default     pod1
        ```
    
6. Create pod and test
    1. Create pod

        ```console
        $ kubectl run -n default pod1 --image=centos:7 --restart=Never --command -- bash -c "trap : TERM INT; (while true;do sleep 1000;done) & wait"
        pod/pod1 created
        $ kubectl get pod pod1 -o wide
        NAME   READY   STATUS    RESTARTS   AGE   IP             NODE    NOMINATED NODE   READINESS GATES
        pod1   1/1     Running   0          5s    10.244.1.156   node1   <none>           <none>
        ```

    2. Test source ip become the ip specified in egress cr

        ```console
        $ kubectl exec -it pod1 bash
        $ IP_TO_MACHINE_OUTSIDE_CLUSTER=192.168.122.64
        $ ping -c 1 ${IP_TO_MACHINE_OUTSIDE_CLUSTER}
        PING 192.168.122.64 (192.168.122.64) 56(84) bytes of data.
        64 bytes from 192.168.122.64: icmp_seq=1 ttl=62 time=1.27 ms
        --- 192.168.122.64 ping statistics ---
        1 packets transmitted, 1 received, 0% packet loss, time 0ms
        rtt min/avg/max/mdev = 1.273/1.273/1.273/0.000 ms
        $ exit
        ```

        while running tcpdump on the machine outside cluster

        ```console
        $ tcpdump -nn icmp
        tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
        listening on eth0, link-type EN10MB (Ethernet), capture size 262144 bytes
        18:19:54.314407 IP 192.168.122.222 > 192.168.122.64: ICMP echo request, id 21, seq 1, length 64
        18:19:54.314504 IP 192.168.122.64 > 192.168.122.222: ICMP echo reply, id 21, seq 1, length 64
        ```

7. Recreate pod and test again to confirm that mapping is updated on pod ip change
    1. Delete the pod

        ```console
        $ kubectl delete pod pod1
        ```
    
    2. Do the same steps to step6

## Undeploy
1. Undeploy egress-mapper

    1. Run deploy command with `-d` option

        ```console
        $ deploy/deploy.sh -d
        deployment.apps "egress-mapper" deleted
        clusterrolebinding.rbac.authorization.k8s.io "kube-keepalived-vip" deleted
        clusterrole.rbac.authorization.k8s.io "kube-keepalived-vip" deleted
        serviceaccount "kube-keepalived-vip" deleted
        customresourcedefinition.apiextensions.k8s.io "vips.egress.mapper.com" deleted
        customresourcedefinition.apiextensions.k8s.io "egresses.egress.mapper.com" deleted
        egressmapper.egress.mapper.com "example-egressmapper" deleted
        customresourcedefinition.apiextensions.k8s.io "egressmappers.egress.mapper.com" deleted
        rolebinding.rbac.authorization.k8s.io "egress-mapper" deleted
        role.rbac.authorization.k8s.io "egress-mapper" deleted
        serviceaccount "egress-mapper" deleted
        ```

    2. Confirm that `egressmapper` deployment, `keepalived-vip` daemonset, and `kube-egress` daemonset are deleted

        ```console
        $ kubectl get deployment
        No resources found.
        $ kubectl get ds
        No resources found.
        $ kubectl get pod
        NAME   READY   STATUS    RESTARTS   AGE
        pod1   1/1     Running   0          11m
        ```

2. Delete configmaps that are no longer needed

    ```console
    $ kubectl delete configmap vip-configmap 
    configmap "vip-configmap" deleted
    $ kubectl delete configmap vip-routeid-mappings
    configmap "vip-routeid-mappings" deleted
    $ kubectl delete configmap podip-vip-mappings
    configmap "podip-vip-mappings" deleted
    ```

3. Test source ip return to the node's ip that the pod is running on

    ```console
    $ kubectl exec -it pod1 bash
    $ IP_TO_MACHINE_OUTSIDE_CLUSTER=192.168.122.64
    $ ping -c 1 ${IP_TO_MACHINE_OUTSIDE_CLUSTER}
    PING 192.168.122.64 (192.168.122.64) 56(84) bytes of data.
    64 bytes from 192.168.122.64: icmp_seq=1 ttl=63 time=0.699 ms
    --- 192.168.122.64 ping statistics ---
    1 packets transmitted, 1 received, 0% packet loss, time 0ms
    rtt min/avg/max/mdev = 0.699/0.699/0.699/0.000 ms
    $ exit
    ```

    while running tcpdump on the machine outside cluster

    ```console
    $ tcpdump -nn icmp
    tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
    listening on eth0, link-type EN10MB (Ethernet), capture size 262144 bytes
    18:33:12.287085 IP 192.168.122.12 > 192.168.122.64: ICMP echo request, id 34, seq 1, length 64
    18:33:12.287179 IP 192.168.122.64 > 192.168.122.12: ICMP echo reply, id 34, seq 1, length 64
    ```


## Limitations
- Only `pod` can be specified as `kind` of resources for `egress` cr, currently. Other resources, such as deployment and daemonset, will be implemented later.
- No check for availability of vip in egress mapping, currently. So, requesting un-available vip might cause cluster unstable.
- No check for whether vip is used by egress on vip deletion, currently. So, deleting vip in use might cause cluster unstable.
