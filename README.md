# go-rancher-suse

In this script, we use Ginkgo to define a test suite called "go_rancher_suse_test". We define one test case using the ginkgo.It function that logs in to the Rancher web page.

To perform the login, we create an HTTP client using bearer token for authentication. We send a POST request to the Rancher login endpoint, passing the token as the header. We expect the response status code to be 200 OK, and we read the response body into a byte slice. Assertion is further performed on the response using gomega.

Note: Folder structure is inspired from K8s. we have a seprate runner file to control hooks like before and after etc. Test cases can be written in seprate file like login_rancher.go.

## To run/ install rancher:
https://ranchermanager.docs.rancher.com/v2.5/pages-for-subheaders/rancher-on-a-single-node-with-docker#option-a-default-rancher-generated-self-signed-certificate

Terminal command:
```
docker run -d --restart=unless-stopped \
  -p 80:80 -p 443:443 \
  --privileged \
  rancher/rancher:latest
```


## To run the test
Clone the repository and 
run cmd

```
cd <root of the code directory>
go mod download
ginkgo -v
```

Sample Test result for reference: 
```
% ginkgo -v
Running Suite: GoRancherSuse Suite - /Volumes/D drive/prsnl_git/go-rancher-suse
===============================================================================
Random Seed: 1681456932

Will run 1 of 1 specs
------------------------------
Given: Rancher API test When: Login API executes with valid token Then: It should be able to login successfully
/Volumes/D drive/prsnl_git/go-rancher-suse/login_rancher.go:21
2023/04/14 12:52:14 Validate that the response status code is 200 OK
2023/04/14 12:52:14 Validate keys in the response
• [0.008 seconds]
------------------------------

Ran 1 of 1 Specs in 0.008 seconds
SUCCESS! -- 1 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS

Ginkgo ran 1 suite in 1.725807215s
Test Suite Passed
```
