package go_rancher_suse_test

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Given: Rancher API test", func() {
	ginkgo.Context("When: Login API executes with valid token", func() {
		url := "https://localhost/v3"
		//method := "POST"
		bearer_token := "Bearer token-l5gpk:g854bjdcfc4f27tpnb8gcvcxx6q59nwwpfwphgdlqvdcjvqzspttbc"

		ginkgo.It("Then: It should be able to login successfully", func() {

			// Perform Login and get response
			resp, err := ExecutePostRequest(url, bearer_token, nil)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			// Check that the response status code is 200 OK
			log.Println("Validate that the response status code is 200 OK")
			gomega.Expect(resp.StatusCode).To(gomega.Equal(http.StatusOK))

			defer resp.Body.Close()

			// Read the response body into a byte slice
			body, err := io.ReadAll(resp.Body)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			// Parse the response body into a map
			var response map[string]interface{}
			err = json.Unmarshal(body, &response)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			//keys := reflect.ValueOf(response).MapKeys()
			//fmt.Println("keys-->", keys)
			//fmt.Println("response-->", response)

			// Assert that the Rancher login was successful by checking for the presence of keys
			log.Println("Validate keys in the response")
			gomega.Expect(response).To(gomega.HaveKey("apiVersion"))
			gomega.Expect(response).To(gomega.HaveKey("baseType"))
			gomega.Expect(response).To(gomega.HaveKey("links"))

		})
	})
})

// Execute post request
func ExecutePostRequest(url string, token string, payload io.Reader) (*http.Response, error) {
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{Transport: customTransport}

	req, err := http.NewRequest(http.MethodPost, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return res, nil
}
