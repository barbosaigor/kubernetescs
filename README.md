# Kubernetes Chaos Server  
A Chaos Server implementation that terminates Kubernetes pods, used by [April](https://github.com/barbosaigor/april) tool.

## What is a Chaos Server ?
Chaos server hosts an API which terminates instances. It is used by [April](https://github.com/barbosaigor/april) CLI, that runs its algorithm and asks the Chaos Server to finish 
any selected instances. All Chaos Servers implementations must implement the interface defined in april/destroyer package, so CSs must include that package and
implement the Destroyer interface, where the business logic to terminate instances should be defined.  

## Installation  
```bash 
go get github.com/barbosaigor/kubernetescs/cmd/...
```   

Kubernetescs hosts API    
-u username for chaos server auth  
-s password for chaos server auth  
-p port number (Default is 7071)  
-n Kubernetes namespace  
-e especify Kubernetes endpoint API  (Default is 127.0.0.1:8001)
```bash 
kubernetescs -u bob -s mysecret
```  
