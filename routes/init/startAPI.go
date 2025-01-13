package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/sky-fbb-uat/sky-auth/v1/user/authenticate", func(c *fiber.Ctx) error {
		println("-------Login SKY-------")

		data := map[string]interface{}{
			"token": "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9..ziAKMwAL_NiFWHecY1UiUqwI36Us0wkSHMYesv8m9ew",
		}

		log.Println("log API login SKY ==> ", (data))

		return c.JSON(data)
	})
	// app.Post("/fbb/changeproduct/conductor/v1/order/qualify", func(c *fiber.Ctx) error {
	// 	println("-------API FBB-------")
	// 	println("-------Request from AC-------")

	// 	data := map[string]interface{}{
	// 		"state": "acknowledge",
	// 		"externalId": []map[string]string{
	// 			{
	// 				"id":    "SO-58btl-240726133936340936104835",
	// 				"owner": "SKY",
	// 			},
	// 			{
	// 				"id":    "CO-qgfts-24072613393670169867",
	// 				"owner": "SFF",
	// 			},
	// 		},
	// 	}

	// 	// dataFail := map[string]interface{}{
	// 	// 	"orderNo": "1234",
	// 	// 	"message": "1234",
	// 	// 	"errors": []map[string]string{
	// 	// 		{
	// 	// 			"id":        "EB090",
	// 	// 			"statement": "Connot 1",
	// 	// 		},
	// 	// 		{
	// 	// 			"id":        "EB090",
	// 	// 			"statement": "Connot 1",
	// 	// 		},
	// 	// 	},
	// 	// }
	// 	// Parse JSON body into struct
	// 	// time.Sleep(11 * time.Second)
	// 	return c.Status(fiber.StatusOK).JSON(data)
	// })

	app.Post("/CollectionService/queryDocument", func(c *fiber.Ctx) error {

		data := map[string]interface{}{
			"Response": map[string]interface{}{
				"receiptList": []map[string]interface{}{
					{
						"receiptCoId":  "W",
						"receiptNo":    "W-OR-1004-6610-10000101",
						"receiptDate":  "11/10/2023",
						"totalMny":     356.52,
						"eReceiptFlag": "N",
						"channelDesc":  "สำ นกังำนใหญ่",
						"periodFirst":  "",
						"periodLast":   "032019",
						"docType":      "RECEIPT",
						"createDtm":    "13/10/2023 20:21:22",
						"methodList": []map[string]interface{}{
							{"methodDesc": "ภำษีหัก ณ ที่จ่ำย"},
						},
						"invoiceList": []map[string]interface{}{
							{
								"invoiceCoId":    "W",
								"invoiceNo":      "W-IN-18-6203-0000008",
								"totalMny":       356.52,
								"billStart":      "01/03/2019",
								"billEnd":        "31/03/2019",
								"paymentDueDate": "26/04/2019",
							},
						},
					},
					{
						"receiptCoId":  "W",
						"receiptNo":    "W-CS-1004-6609-10000018",
						"receiptDate":  "13/09/2023",
						"totalMny":     1.0,
						"eReceiptFlag": "N",
						"channelDesc":  "สำ นกังำนใหญ่",
						"periodFirst":  "",
						// "periodLast":  "032019",
						"docType":   "RECEIPT",
						"createDtm": "13/09/2023 13:12:11",
						"methodList": []map[string]interface{}{
							{"methodDesc": "เงินสด"},
						},
						"invoiceList": []map[string]interface{}{
							{
								"invoiceCoId":    "W",
								"invoiceNo":      "W-IN-18-6203-0000008",
								"totalMny":       1.0,
								"billStart":      "01/03/2019",
								"billEnd":        "31/03/2019",
								"paymentDueDate": "26/04/2019",
							},
						},
					},
				},
			},
			"ErrorDesc": nil,
			"ErrorMsg":  "Success",
			"ErrorCode": "000",
		}

		return c.JSON(data)
	})
	app.Post("/Resources/v1/Fulfillment/PGZInquiry/synchronous/ServiceProvisioning", func(c *fiber.Ctx) error {
		data := map[string]interface{}{
			"responseHeader": map[string]interface{}{
				"customerOrderType": "Query Bill Cycle",
				"reTransmit":        "0",
				"sourceSystem":      "BSS",
				"userSys":           "CAD",
				"resourceGroupId":   "rbmTransChangeBillCycle-20240304150200",
				"resourceOrderId":   "DBSSPHXA002G-PGZINQ-20240304150200",
				"resultCode":        "20000",
				"resultDesc":        "Success",
				"developerMessage":  "",
			},
			"resourceItemList": []map[string]interface{}{
				{
					"resourceItemId":         "rbmTransChangeBillCycle-20240304150200",
					"resourceName":           "rbmTransChangeBillCycle",
					"resourceItemStatus":     "Success",
					"resourceItemErrMessage": "Success",
					"errorFlag":              "1",
					"resourceActivatedDate":  "20240304150200+0700",
					"transStatus":            "PENDING",
					"specialErrHandling": map[string]interface{}{
						"suppCode":             []map[string]interface{}{},
						"taskKeyCondition":     []map[string]interface{}{},
						"taskDeveloperMessage": []map[string]interface{}{},
					},
				},
			},
		}

		return c.JSON(data)
	})

	app.Get("/mobile-postpaid/conductor/v1/profileChange", func(c *fiber.Ctx) error {
		// Retrieve the query string parameter
		number := c.Query("accountNumber")

		if number == "31700015619323" {
			// Response for account number 31700015619323
			data := map[string]interface{}{
				"profileChangNo":    "C0912",
				"accountId":         "323232",
				"accountNumber":     "AAAA1",
				"state":             "Waiting",
				"stateDate":         "01/01/2009 00:00:00",
				"effectiveDate":     "08/01/2010 00:00:00",
				"createdDate":       "01/10/2009 15:57:02",
				"lastUpdateDate":    "01/01/2009 15:57:02",
				"createdBy":         "sasithch",
				"lastUpdateBy":      "sasithch",
				"jobCode":           "AA",
				"jobSequench":       "2",
				"resourceOrderId":   "",
				"transStatus":       "",
				"passedPostProfChg": "1",
				"prevReq":           "8a8d35b22432c8ab0124345e25c504a7",
				"parRowId":          "8a8d35b22408851601240f6c151c0074",
				"newAttrib03":       "12",
				"newAttrib22":       "1 Old Cycle(s), New Cycle on 08/01/2010",
				"oldAttrib03":       "18",
				"type":              "Bill Cycle Change",
			}
			return c.Status(200).JSON(data)
		} else if number == "21700015619323" {
			// Response for account number 21700015619323
			data := map[string]interface{}{
				"message":          "The item does not exist",
				"developerMessage": "Please input new Item",
			}
			return c.Status(500).JSON(data)
		} else if number == "11700015619323" {
			// Response for account number 21700015619323
			data := map[string]interface{}{
				"message": "The item does not exist",
			}
			return c.Status(404).JSON(data)
		} else if number == "01700015619323" {
			// Response for account number 21700015619323
			data := map[string]interface{}{
				"message": "The requested url was not found",
			}
			return c.Status(405).JSON(data)
		}
		return c.Status(400).JSON(map[string]string{
			"message": "Query queryProfileChange fail,required accountNumber",
		})
	})
	app.Post("/mydigitalid/myDID/v1/partner/verify-pin-by-personal-ref-id", func(c *fiber.Ctx) error {

		var body map[string]interface{}

		if err := c.BodyParser(&body); err != nil {
			log.Printf("Error parsing body: %v", err)

			return c.Status(400).JSON(map[string]interface{}{
				"resultDesc": "Invalid Request Body",
				"resultCode": "40001",
			})
		}

		pinCode := body["pin_code"]

		log.Printf("Received pinCode: %v", pinCode)

		// Response
		if pinCode == "475329" {
			data := map[string]interface{}{
				"resultDesc": "Success",
				"resultCode": "20000",
				"resultData": map[string]interface{}{
					"pin_verification_result": "Y",
				},
			}

			return c.Status(200).JSON(data)

		} else if pinCode == "375329" {
			data := map[string]interface{}{
				"resultCode":    "40300",
				"resultDesc":    "Forbidden",
				"resultMessage": "Missing or invalid parameter",
			}

			return c.Status(403).JSON(data)

		} else {
			return c.SendStatus(429)
		}

	})
	app.Use(func(c *fiber.Ctx) error {
		c.Context().Conn().Close()
		return nil
	})

}
