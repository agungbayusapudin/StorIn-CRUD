package generateinvoicenumber

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateInvoiceNumber() string {
	id := uuid.New()
	return fmt.Sprintf("INV-%s", id.String())
}
