# AWS Well Architected Labs - Reliability - 100: Deploy Using CloudFormation

This project follows the steps provided on [AWS Well-Architected Labs](https://wellarchitectedlabs.com/reliability/100_labs/100_deploy_cloudformation/).

## Step 1 - Deploy Infrastructure

1. Log into the AWS console and choose the `us-east-2` Ohio region.
2. [Download `vpc-alb-app-db.yaml`](https://wellarchitectedlabs.com/Common/Create_VPC_Stack/Code/vpc-alb-app-db.yaml) for a standard VPC CloudFormation template.
3. Go to CloudFormation and create a new stack by uploading the template file. Name it WebApp1-VPC. Click through setup and acknowledge IAM resources with custom names.
4. Wait for stack to enter `CREATE_COMPLETE` status.


## Step 2 - Deploy Application

1 [Download `staticwebapp.yaml`](https://wellarchitectedlabs.com/Reliability/Common/Code/CloudFormation/staticwebapp.yaml). 
2. Upload it as stack like the VPC using the name `CloudFormationLab`.
3. Wait for stack to enter `CREATE_COMPLETE` status.


## Step 3 - Explore Web Application

1. Click on `CloudFormationLab` stack and go to the Outputs tab to find a `WebsiteURL`. Click it.
2. Explore the application.


## Step 4 - Explore CloudFormation

1. Go to each CloudFormation stack and click on Template.
2. Read through the templates and see how the VPC and application were deployed.

## Step 5 - Clearing up

1. Delete the CloudFormationLab stack.
2. Delete the WebApp1-VPC stack.
