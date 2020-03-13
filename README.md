# Kubernetes Chaos Server  
An chaos server implementation that terminates Kubernetes pods via an API.

## What is a Chaos Server ?
Chaos server hosts an API that terminantes instances. It is used in [April](https://github.com/barbosaigor/april) CLI, which run its algorithm and asks the Chaos server to finish 
the selected instances. The API implementation lives in april/destroyer package, so chaos servers must include that package and
implement the Destroyer interface, which contain the business logic for terminate instances.  

## Installation  
```bash 
go install github.com/barbosaigor/kubernetescs
```  

Kubernetescs hosts API    
-u username for chaos server auth  
-s password for chaos server auth  
-p port number (Default is 7071)  
-n Kubernetes namespace  
--host especify Kubernetes host API  
```bash 
kubernetescs -u bob -s mysecret -n dev
```  
