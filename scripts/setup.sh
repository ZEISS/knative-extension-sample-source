#!/bin/bash

set -x

# Get the current directory
dir=$(pwd)

# Create kind cluster
kind create cluster --config $dir/.cluster.yaml

kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.17.0/serving-crds.yaml
kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.17.0/serving-core.yaml
kubectl apply -f https://github.com/knative/eventing/releases/download/knative-v1.17.2/eventing-crds.yaml
kubectl apply -f https://github.com/knative/eventing/releases/download/knative-v1.17.2/eventing-core.yaml
kubectl apply -f https://github.com/knative/eventing/releases/download/knative-v1.17.2/in-memory-channel.yaml