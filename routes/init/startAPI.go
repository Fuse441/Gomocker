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

	app.Use(func(c *fiber.Ctx) error {
		c.Context().Conn().Close()
		return nil
	})

}
