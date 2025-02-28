#!/bin/bash

set -x

# Get the current directory
dir=$(pwd)

# Create kind cluster
kind create cluster --config $dir/.cluster.yaml