package main

import (
	"fmt"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hamms", func() {
	go SleepFor()
	time.Sleep(1 * time.Second)
	Describe("SleepFor", func() {
		Context("Http request sent to /sleep", func() {
			It("Should not be nil", func() {

				resp, err := http.Get("http://localhost:5508/sleep/5")
				defer resp.Body.Close()
				Expect(err).Should(BeNil())
				Expect(resp).ShouldNot(BeNil())

			})
			It("Should respond in 5 seconds", func() {
				start := time.Now()
				resp, _ := http.Get("http://localhost:5508/sleep/5")
				defer resp.Body.Close()
				elapsed := time.Since(start)

				fmt.Println(elapsed)

			})

		})

	})
})
