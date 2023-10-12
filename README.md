# DX-test

## Stage 1

This is a Golang API to manage products. We would like you to add a route that will return only a single product based on an ID provided 

e.g. ```curl http://localhost:10000/product/1``` would return
```{"Id":"1","Name":"Airpods Pro","desc":"Apple Airpods Pro","price":"350"}```

## Stage 2
Create a route that will update the price of a product.


## Stage 3
We've provided a Dockerfile to build a docker image from this code. Identify and describe any problems or inefficiencies you can find, and fix them to create your image.

## Stage 4
Create a pipeline to build and deploy the services with sensible behaviours around branches and provide a means for engineers to deploy the services to AWS.

At the end of Stage 4 an engineer should be able to deploy a new version of your service in AWS with the artefacts created during Stage 1. There is no requirement to deploy all the surrounding infrastructure. The goal is to provide a sensible deployment approach.

## Stage 5
Describe how you would scale your development, testing and deployment process to be able to deploy multiple times per day for multiple services.

As a benchmark assume that there are a number of teams which work on a total of about 100 different services. Each team would like to 
be able to create new features, test them and deploy them as quickly as possible without causing issues for the customers. Assume that there will be
about 150 deployments per day.
