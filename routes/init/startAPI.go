package routes

import (
	"log"
	"time"

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
	app.Post("/sky-fbb-uat/fbb/account-profile/conductor/v1/order", func(c *fiber.Ctx) error {
		println("-------API FBB-------")
		println("-------Request from AC-------")

		data := map[string]interface{}{

			"orderNo":    "FS-sf5zp-241211170351817135375511",
			"orderRefId": "TF-77xjh-24121117035189134539",
		}

		// time.Sleep(22 * time.Second)

		return c.Status(200).JSON(data)
	})
	app.Post("/sky-fbb-uat/fbb/change-product/conductor/v1/order/qualify", func(c *fiber.Ctx) error {
		println("-------API FBB-------")
		println("-------Request from AC-------")

		data := map[string]interface{}{

			"orderNo":    "FS-sf5zp-241211170351817135375511",
			"orderRefId": "TF-77xjh-24121117035189134539",
		}

		time.Sleep(22 * time.Second)
		return c.Status(200).JSON(data)
	})
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
					"errorFlag":              "0",
					"resourceActivatedDate":  "20240304150200+0700",
					"transStatus":            "PENDING",
					"specialErrHandling": map[string]interface{}{
						"suppCode":         []map[string]interface{}{},
						"taskKeyCondition": []map[string]interface{}{},
						"taskDeveloperMessage": []string{
							"asd",
							"asd",
						},
					},
				},
			},
		}
		return c.JSON(data)
	})

	app.Get("/mobile-postpaid/conductor/v1/profileChange", func(c *fiber.Ctx) error {
		log.Printf("------- API mobile POSTPAID -------")

		accountNumber := c.Query("accountNumber")
		profileChangeNo := c.Query("profileChangeNo")
		state := c.Query("state")

		log.Printf("AccountNumber: %s, ProfileChangeNo: %s, State: %s", accountNumber, profileChangeNo, state)

		if accountNumber == "31700015619323" {
			data := []map[string]interface{}{
				{
					"rowId":             "8c4ed041-4516-428c-a2db-8f63f8b2d787",
					"type":              "Billing Profile Change",
					"accntId":           "8a7cc0198beb4169018bebe5403d010e",
					"statusCd":          "Completed",
					"statusDate":        "2023-11-29",
					"effectiveDate":     "2023-11-29",
					"statusUpdBy":       "SKY_USR",
					"prevReq":           "deab28b6-5132-40d6-84bb-f7acaf5730b8",
					"parRowId":          "8a7cc0198beb4169018bebe53ba40109",
					"created":           "2023-11-29",
					"createdBy":         "SKY_USR",
					"lastUpd":           "2023-11-29",
					"lastUpdBy":         "SKY_USR",
					"profileChgNo":      "C2311001353574",
					"newAttrib01":       "Y",
					"newAttrib02":       "N",
					"newAttrib05":       "Mr.Muscle Muscle",
					"newAttrib06":       "Corporate (Full)",
					"newAttrib08":       "THB",
					"newAttrib16":       "10400",
					"newAttrib17":       "0983982933",
					"newAttrib19":       "AWN",
					"newAttrib20":       "Next bill",
					"newAttrib23":       "11 rr",
					"newAttrib24":       "e",
					"newAttrib26":       "แขวงสามเสนใน เขตพญาไท",
					"newAttrib27":       "กรุงเทพ",
					"oldAttrib01":       "Y",
					"oldAttrib02":       "N",
					"oldAttrib05":       "Mr.Muscle Muscle",
					"oldAttrib06":       "Corporate (Full)",
					"oldAttrib08":       "THB",
					"oldAttrib16":       "10400",
					"oldAttrib17":       "0983982933",
					"oldAttrib19":       "AWN",
					"oldAttrib20":       "Next bill",
					"oldAttrib23":       "11 rr",
					"oldAttrib24":       "e",
					"oldAttrib26":       "แขวงสามเสนใน เขตพญาไท",
					"oldAttrib27":       "กรุงเทพ",
					"interfaceFlg":      "Y",
					"tmTransactionId":   "T2311000767641",
					"jobSequence":       "0",
					"passedPostProfChg": "N",
					"oldAttrib56":       "N",
					"oldAttrib57":       "0893640900",
					"oldAttrib58":       "Tester@gmail.com",
					"newAttrib56":       "N",
					"newAttrib57":       "0893640924",
					"newAttrib58":       "Tester@gmail.com",
					"oldAttrib60":       "SFF",
					"newAttrib60":       "SFF",
				},
			}

			return c.Status(200).JSON(data)

		} else if accountNumber == "21700015619323" {
			data := map[string]interface{}{
				"message": "The requested url was not found",
			}

			return c.Status(405).JSON(data)
		} else if accountNumber == "11700015619323" {
			data := map[string]interface{}{
				"message": "The item does not exist",
			}

			return c.Status(404).JSON(data)
		} else if accountNumber == "01700015619323" {
			data := map[string]interface{}{
				"message": "Query queryProfileChange fail,required accountNumber",
			}

			return c.Status(400).JSON(data)
		} else {
			return c.Status(500).JSON(map[string]interface{}{"error": "Invalid accountNumber"})

		}

	})

	app.Use(func(c *fiber.Ctx) error {
		c.Context().Conn().Close()
		return nil
	})

}
