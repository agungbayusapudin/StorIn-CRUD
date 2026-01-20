package services

import "videocall/internal/app/billing/schema"

// flow billing
// 1. user pilih metode pembayaran
// 2. user melakukan pembayaran
// 3. sistem memverifikasi pembayaran
// 4. jika pembayaran berhasil, sistem mengupdate status langganan user
// 5. jika pembayaran gagal, sistem mengirim notifikasi kegagalan kepada user

type BillingServiceInterface interface {
	InitialPayment(userId int, paymentReq *schema.CreateInvoiceRequest) error
	CancelPayment(userId int) error
	VerifyPayment(userId int, paymentId string) error
	UpdatePaymentStatus(userId int, status string) error
}
