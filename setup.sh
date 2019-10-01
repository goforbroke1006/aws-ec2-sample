#!/usr/bin/env bash

aws s3 cp s3://aws-codedeploy-eu-central-1/samples/latest/SampleApp_Linux.zip . --region eu-central-1
unzip SampleApp_Linux.zip
rm -f SampleApp_Linux.zip