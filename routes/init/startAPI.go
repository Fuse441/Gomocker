package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
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
						"periodLast":   "032019",
						"docType":      "RECEIPT",
						"createDtm":    "13/09/2023 13:12:11",
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

		fmt.Println("in come rounte")
		return c.JSON(data)
	})
}
