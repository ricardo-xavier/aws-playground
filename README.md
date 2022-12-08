# Mycloud

## Objetivo

O objetivo do mycloud é permitir a criação e utilização de uma infraestrutura local semelhante a AWS.
Essa infraestrutura poderá ser utilizada para estudos, experimentos e testes.

Os micro-serviços que utilizam o sdk java da aws poderão acessar a infraestrutura do mycloud sem nenhuma alteração nos programas, basta alterar a dependência no pom.

Exemplo:

    sdk do dynamodb da aws:
        <dependency>
            <groupId>software.amazon.awssdk</groupId>
            <artifactId>dynamodb-enhanced</artifactId>
            <version>2.18.24</version>
        </dependency>

    sdk do nosql do mycloud:
        <dependency>
            <groupId>mycloud.sdk</groupId>
            <artifactId>nosql</artifactId>
            <version>0.1.0-SNAPSHOT</version>
        </dependency>

## Módulos

- servicos
    - nosql (dynamodb)
    - messaging (SQS e SNS)
    - storage (S3)
    - secrets (parameter store)
    - cache (memcached)
- cli - acesso via terminal com a mesma sintaxe do cli da aws
- iac - manutenção da infraestrutura com a sintaxe do terraform
- sdk - sdk java com a mesma interface do sdk da aws
- demo - programas de exemplo e POCs

O sdk java será escrito em JAVA, o restante será escrito em GO.

RELEASE NOTES:
0.1.0 - nosql scan e filter
