package spool_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/charlesfan/spool"
)

func TestPrint(t *testing.T) {
	pm := spool.NewPrintManager()

	// two printers
	pm.RegisterPrinter(spool.NewMockPrinter("Printer-A"))
	pm.RegisterPrinter(spool.NewMockPrinter("Printer-B"))

	// Metadata function
	pm.SetMetadataFunc(func(img []byte) map[string]string {
		return map[string]string{
			"user": "admin-testing",
			"time": time.Now().Format(time.RFC3339),
		}
	})

	pm.Callback = func(result spool.PrintJob) {
		fmt.Printf("✅ Job finish! Printer: %s, JobID: %s, Status: %s, Meta: %+v\n",
			result.Name, result.ID, result.Status, result.Metadata)
	}

	for i := 0; i < 5; i++ {
		data := []byte(fmt.Sprintf("photo-%d.jpg", i))
		err := pm.SubmitPrintJob(data)
		if err != nil {
			fmt.Println("print error: ", err)
		}
	}

	time.Sleep(5 * time.Second)
}
