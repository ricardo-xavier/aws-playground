# aws-playground

## Purpose

The purpose of aws-playground is to enable the creation an use of a local aws-like infrastructure.
This infrastructure can be used for studies, experiments and tests.

The micro-services that use the aws java sdk will be able to access the aws-playground without any changes in the code. Just change the dependency in the pom.xml.

Example:
    aws dynamodb sdk
        <dependency>
            <groupId>software.amazon.awssdk</groupId>
            <artifactId>dynamodb-enhanced</artifactId>
            <version>2.18.24</version>
        </dependency>

    aws-playground sdk
        <dependency>
            <groupId>awsplayground.sdk</groupId>
            <artifactId>nosql</artifactId>
            <version>0.2.0-SNAPSHOT</version>
        </dependency>

## Modules

- services
    - nosql (dynamodb)
    - messaging (SQS and SNS)
    - storage (S3)
    - secrets (parameter store)
    - cache (memcached)
- cli - terminal access with aws cli sintaxe
- iac - infrastructure maintenance with terraform sintax
- sdk - java sdk
- demo - examples and POCs
