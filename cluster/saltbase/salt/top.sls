base:
  '*':
    - base

  'roles:kubernetes-pool':
    - match: grain
    - golang
    - docker
    - kubelet
    - kube-proxy
    - cadvisor
    - nsinit

  'roles:kubernetes-master':
    - match: grain
    - golang
    - etcd
    - apiserver
    - controller-manager
    - scheduler
    - build-controller
    - nginx

  'roles:kubernetes-pool-vsphere':
    - match: grain
    - static-routes
