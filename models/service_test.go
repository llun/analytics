package models_test

import (
	"encoding/json"
	. "github.com/llun/analytics/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Models/Service", func() {

	Context("#NewService", func() {

		It("should convert configuration to json and store it as string", func() {

			configuration := map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			}
			service := NewService(nil, "mockservice", SERVICE_OTHER, configuration)

			bytes, _ := json.Marshal(configuration)
			Expect(service.Configuration).To(Equal(string(bytes)))
		})

	})

	Describe("Service", func() {

		Context("#ToMap", func() {

			It("should export all public properties", func() {

				service := NewService(nil, "mockservice", SERVICE_OTHER, map[string]interface{}{
					"key": "value",
				})
				Expect(service.ToMap()).To(Equal(map[string]interface{}{
					"name": "mockservice",
					"type": "other",
				}))

			})

		})

		Context("#GetConfiguration", func() {
			It("should unmarshal configuration to map object", func() {

				service := Service{
					Name:          "MockService",
					Configuration: `{"key1":"value1","key2":"value2","key3":{"key4":"value4"}}`,
				}
				configuration := service.GetConfiguration()
				Expect(configuration).To(Equal(map[string]interface{}{
					"key1": "value1",
					"key2": "value2",
					"key3": map[string]interface{}{
						"key4": "value4",
					},
				}))

			})
		})

	})

})
