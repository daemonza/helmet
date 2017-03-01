# Helmet

Helmet is simple and easy to use Helm(https://helm.sh) repository, useful
when you want to setup a private Helm repository and be able to easy
upload new `helm charts`. Typically deployed to the same Kubernetes cluster in which you 
want to deploy helm charts into. 

Example of uploading a chart to a locally running Helmet.

```
curl -T testapi-chart-0.1.0.tgz -X PUT http://127.0.0.1:1323/upload/
```

After the update you can confirm that the Helm index.yaml got created
by running 

```
curl http://127.0.0.1:1323/charts/index.yaml
```

Output should look similar to :

```
apiVersion: v1
entries:
  testapi-chart:
  - apiVersion: v1
    created: 2017-02-24T12:15:09.995448981+01:00
    description: Test API
    digest: bb1291bb38cf19f583892789e233c6b94ca832853845749c1bbcbd4d92eeb844
    name: testapi-chart
    urls:
    - http://localhost:1323/testapi-chart-0.1.0.tgz
    version: 0.1.0
generated: 2017-02-24T12:15:09.994657561+01:00
```
